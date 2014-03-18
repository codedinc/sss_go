package main

import (
  "fmt"
  "regexp"
  "io"
)

type Token struct {
  kind      int
  value     string
}

type Lexer struct {
  reader    io.Reader
  token    *Token      // Last token returned
  buf     []byte
  start     int
  end       int
  eof       bool
}

func newLexer(reader io.Reader) (lexer *Lexer) {
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
  if (l.end == 0) {
    l.read()
  }

  // Loop until we have a token.
  for {
    advance, token := l.findToken(l.buf[l.start:l.end])
    
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

var spaceRegExp = regexp.MustCompile(`^\s+`)
var numberRegExp = regexp.MustCompile(`^[0-9]+`)
var identifierRegExp = regexp.MustCompile(`^[a-zA-Z][\w\-]*`)

func (l *Lexer) findToken(data []byte) (advance int, tok *Token) {
  var match []byte

  if match = spaceRegExp.Find(data); match != nil {
    // Skip spaces
    return len(match), nil
  } else if match = numberRegExp.Find(data); match != nil {
    return token(NUMBER, match)
  } else if match = identifierRegExp.Find(data); match != nil {
    return token(IDENTIFIER, match)
  }

  // Catch all remaining single char as token
  char := data[:1]
  return 1, &Token{int(char[0]), string(char)}
}

func token(kind int, value []byte) (advance int, tok *Token) {
  return len(value), &Token{kind, string(value)}
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
