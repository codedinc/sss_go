parser.go: grammar.y
	go tool yacc -o parser.go grammar.y