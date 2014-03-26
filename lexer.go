package main

import (
	"fmt"
	"io"
	"regexp"
)

type Token struct {
	kind  int
	value string
}

type Lexer struct {
	reader io.Reader
	token  *Token // Last token returned
	buf    []byte
	start  int
	end    int
	eof    bool
	output *StyleSheet
}

type LexingRule struct {
	tokenType int
	regexp    *regexp.Regexp
}

// Reusable regexp chunks (macros)
var number = `[0-9]+(\.[0-9]+)?`        // matches: 10 and 3.14
var name = `[a-zA-Z][\w\-]*`            // matches: body, background-color and myClassName
var selector = `(\.|\#|\:\:|\:)` + name // matches: #id, .class, :hover and ::before

// Define the lexing rules
var rules = []LexingRule{
	// Token type   Regexp to match that token
	// (-1 ignore)
	LexingRule{-1, regexp.MustCompile(`^\s+`)},

	LexingRule{DIMENSION, regexp.MustCompile(`^` + number + `(px|em|\%)`)},
	LexingRule{NUMBER, regexp.MustCompile(`^` + number)},
	LexingRule{COLOR, regexp.MustCompile(`^\#[0-9A-Fa-f]{3,6}`)},

	LexingRule{SELECTOR, regexp.MustCompile(`^` + selector)},
	LexingRule{SELECTOR, regexp.MustCompile(`^` + name + selector)},

	LexingRule{IDENTIFIER, regexp.MustCompile(`^` + name)},
}

func NewLexer(reader io.Reader) (lexer *Lexer) {
	return &Lexer{
		reader: reader,
		buf:    make([]byte, 4096),
	}
}

// Read and fill the buffer (buf).
func (l *Lexer) read() int {
	// Based on http://golang.org/src/pkg/bufio/scan.go

	// Shift data to beginning of buffer if there's lots of empty space
	if l.start > 0 && (l.end == len(l.buf) || l.start > len(l.buf)/2) {
		copy(l.buf, l.buf[l.start:l.end])
		l.end -= l.start
		l.start = 0
	}

	// Is the buffer full? If so, resize.
	if l.end == len(l.buf) {
		newSize := len(l.buf) * 2
		newBuf := make([]byte, newSize)
		copy(newBuf, l.buf[l.start:l.end])
		l.buf = newBuf
		l.end -= l.start
		l.start = 0
	}

	// Read some input
	n, _ := l.reader.Read(l.buf[l.end:])
	l.end += n
	l.eof = n == 0

	return n
}

func (l *Lexer) Scan() bool {
	if l.end == 0 {
		l.read()
	}

	// Loop until we have a token.
	for {
		advance, token := l.applyRules(l.buf[l.start:l.end])

		l.start += advance
		l.token = token

		// Woohoo! A token.
		if token != nil {
			return true
		}

		// No more data, we're done
		if l.start == l.end && l.eof {
			return false
		}

		// Read more data
		l.read()
	}
}

// Apply the lexing rules and return matched token.
func (l *Lexer) applyRules(data []byte) (advance int, token *Token) {
	var match []byte

	// Apply each lexing rule until one matches
	for _, rule := range rules {
		match = rule.regexp.Find(data)

		if match != nil {
			token = nil
			// tokenType of -1 means we ignore and do not emit a token.
			if rule.tokenType != -1 {
				token = &Token{rule.tokenType, string(match)}
			}
			return len(match), token
		}
	}

	// Catch all remaining single char as token
	char := data[:1]
	return 1, &Token{int(char[0]), string(char)}
}

//// API used by the parser.

func (l *Lexer) Error(e string) {
	fmt.Printf("%v: %q\n", e, l.token.value)
}

func (l *Lexer) Lex(lval *yySymType) int {
	if l.Scan() {
		lval.value = l.token.value

		return l.token.kind
	}

	// EOF
	return 0
}
