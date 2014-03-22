package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	// yyDebug = 1
	lexer := NewLexer(bufio.NewReader(input))
	yyParse(lexer)
	fmt.Printf("%v\n", lexer.output)

	input.Close()
}
