package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Compile(input io.Reader) string {
	styleSheet := Parse(bufio.NewReader(input))
	return styleSheet.ToCSS()
}

func CompileString(input string) string {
	styleSheet := Parse(strings.NewReader(input))
	return styleSheet.ToCSS()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: sss <file-name>\n")
		return
	}

	input, err := os.Open(os.Args[1])
	defer input.Close()
	if err != nil {
		fmt.Printf("Error! Path: %s does not exist.\n", os.Args[1])
		return
	}

	styleSheet := Parse(bufio.NewReader(input))
	fmt.Printf("%v\n", styleSheet.ToCSS())
}
