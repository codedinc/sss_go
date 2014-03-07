sss: main.go parser.go lexer.go
	go build -o $@ $^

parser.go: grammar.y
	go tool yacc -o $@ $^

lexer.go: tokens.l
	golex -o $@ $^

clean:
	rm -f sss parser.go lexer.go

.PHONY: clean