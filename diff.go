package diff

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// LineDiff stores a specific difference of two compared files.
type LineDiff struct {
	line    int
	added   bool
	content string
}

func (d LineDiff) String() string {
	prefix := "-"
	if d.added {
		prefix = "+"
	}
	return fmt.Sprintf("%v \t %v \t %v", d.line, prefix, d.content)
}

// Compare returns the difference of the left and right file.
// Differences (added or removed line) are based on the left file.
func Compare(leftFile, rightFile string) []LineDiff {
	var diffs []LineDiff
	endLeft, endRight := false, false
	contNumRight := 0

	for numLeft := 0; ; numLeft++ {
		contentLeft, err := getLine(numLeft, leftFile)
		if err != nil {
			if err == errEndOfFile {
				endLeft = true
			} else {
				log.Fatal(err)
			}
			if endLeft == true && endRight == true {
				// Finished scanning both files
				break
			}
		}

		for numRight := contNumRight; ; numRight++ {
			contentRight, err := getLine(numRight, rightFile)
			if err != nil {
				if err != errEndOfFile {
					log.Fatal(err)
				}

				if endLeft == false {
					// scanned whole right side,
					// left line not existing anymore
					diffs = append(diffs, LineDiff{line: numLeft + 1, added: false, content: contentLeft})
					break
				}

				// finished left side
				// get remaining lines at end of right side
				endRight = true
				for addLineRight := contNumRight; addLineRight < numRight; addLineRight++ {
					newLine, err := getLine(addLineRight, rightFile)
					if err != nil && err != errEndOfFile {
						log.Fatal(err)
					}
					diffs = append(diffs, LineDiff{line: addLineRight + 1, added: true, content: newLine})
				}
				// break here, otherwise infinite loop
				break
			}

			// found the corresponding matching line on right side
			if contentLeft == contentRight {
				// get newly added lines on right side in between the matched lines
				for addLineRight := contNumRight; addLineRight < numRight; addLineRight++ {
					newLine, err := getLine(addLineRight, rightFile)
					if err != nil && err != errEndOfFile {
						log.Fatal(err)
					}
					diffs = append(diffs, LineDiff{line: addLineRight + 1, added: true, content: newLine})
				}
				contNumRight = numRight + 1
				break
			}
		}
	}
	return diffs
}

var errEndOfFile = errors.New("reached end of file")

func getLine(line int, fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	curLine := 0
	for scanner.Scan() {
		if curLine == line {
			return scanner.Text(), nil
		}
		curLine++
	}
	return "", errEndOfFile
}

func getContent(fileName string) ([]string, error) {
	var content []string
	file, err := os.Open(fileName)
	if err != nil {
		return []string{""}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		content = append(content, scanner.Text())

	}
	return content, nil
}
