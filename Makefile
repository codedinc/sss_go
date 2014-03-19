GO_FILES = main.go nodes.go parser.go lexer.go

sss: $(GO_FILES)
	go build -o $@ $^

parser.go: grammar.y
	go tool yacc -o $@ $^

test: sss *_test.go
	go test

clean:
	rm -f sss parser.go

.PHONY: clean test