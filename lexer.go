package main

import (
  "bufio"
  "fmt"
  "regexp"
  "io"
  "encoding/binary"
)

type Lexer struct{
  scanner  *bufio.Scanner
  current   string
}

func newLexer(reader io.Reader) (lexer *Lexer) {
  scanner := bufio.NewScanner(reader)
  scanner.Split(scan)
  return &Lexer{scanner: scanner}
}

var spaceRegExp = regexp.MustCompile(`^\s+`)
var numberRegExp = regexp.MustCompile(`^[0-9]+`)
var identifierRegExp = regexp.MustCompile(`^[a-zA-Z][\w\-]*`)

// Split function for the Scanner.
func scan(data []byte, atEOF bool) (advance int, token []byte, err error) {
  var match []byte

  if atEOF && len(data) == 0 {
    return 0, nil, nil
  }

  if match = spaceRegExp.Find(data); match != nil {
    // Skip spaces
    return len(match), nil, nil
  } else if match = numberRegExp.Find(data); match != nil {
    return encodeToken(match, NUMBER)
  } else if match = identifierRegExp.Find(data); match != nil {
    return encodeToken(match, IDENTIFIER)
  }

  // Catch all remaining single char as token
  tok := data[:1]
  return encodeToken(tok, int(tok[0]))
}


// Go's built-in Scanner only supports returning a token value. We need to
// return a token type too. To get around this limitation, we encode the
// token type in the first three bytes of the token's value returned by the
// scanner.
// The following functions take care of encoding and decoding the tokens.
//
// Token bytes format:
// - First 3 byte: token type (int)
// - Remaining bytes: token value (string)

var tokTypeBytesLen = 3 // # of bytes to store the type when encoded in a []byte.

func encodeToken(value []byte, tokenType int) (advance int, token []byte, err error) {
  token = make([]byte, len(value) + tokTypeBytesLen)
  writenLen := binary.PutUvarint(token, uint64(tokenType))

  if (writenLen > 3) {
    panic(fmt.Sprintf("Encoding failed. Expected token type to be written to %v bytes, used %v bytes.",
                      tokTypeBytesLen, writenLen))
  }
  
  copy(token[tokTypeBytesLen:], value)

  return len(value), token, nil
}

func decodeToken(buf []byte) (tokenType int, value string) {
  val, _ := binary.Uvarint(buf)
  return int(val), string(buf[tokTypeBytesLen:])
}


//// API used by the parser.

func (l Lexer) Error(e string) {
  fmt.Printf("%v: %q\n", e, l.current)
}

func (l Lexer) Lex(lval *yySymType) int {
  if l.scanner.Scan() {
    tokType, value := decodeToken(l.scanner.Bytes())

    l.current = value
    lval.value = value

    return int(tokType)
  }

  // EOF
  return 0
}
