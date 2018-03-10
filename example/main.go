package main

import (
	"fmt"

	"github.com/sj14/diff"
)

func main() {
	diffs := diff.Compare("left.xml", "right.xml")

	for _, d := range diffs {
		fmt.Println(d)
	}
}
