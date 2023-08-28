package main

import (
	"fmt"

	"github.com/fengde/gocommon/colorx"
)

func main() {
	s := colorx.WithColorPadding("hello world", colorx.FgRed)
	fmt.Println(s)
}
