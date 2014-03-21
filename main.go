package main

import (
	"bufio"
	"os"
)

func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	// yyDebug = 1
	yyParse(NewLexer(bufio.NewReader(input)))

	input.Close()
}
