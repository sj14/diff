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

	contentLeft, err := getContent(leftFile)
	if err != nil {
		log.Fatal(err)
	}

	contentRight, err := getContent(rightFile)
	if err != nil {
		log.Fatal(err)
	}

	continueNumRight := 0

	for il := 0; il < len(contentLeft); il++ {
		found := false
		for ir := continueNumRight; ir < len(contentRight); ir++ {
			if contentLeft[il] == contentRight[ir] {
				found = true
				// output lines on right side which were in between last matching and currently matching line
				for k := continueNumRight; k < ir; k++ {
					diffs = append(diffs, LineDiff{line: k + 1, added: true, content: contentRight[k]})
				}
				continueNumRight = ir + 1
				break
			}
		}
		if !found {
			// same content not found on right side
			diffs = append(diffs, LineDiff{line: il + 1, added: false, content: contentLeft[il]})
		}
	}
	// add remaining new lines at the end of right side
	for k := continueNumRight; k < len(contentRight); k++ {
		diffs = append(diffs, LineDiff{line: k + 1, added: true, content: contentRight[k]})
	}
	return diffs
}

var ErrEndOfFile = errors.New("reached end of file")

func getContent(fileName string) ([]string, error) {
	var content []string
	file, err := os.Open(fileName)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content, nil
}
