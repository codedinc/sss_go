%{
// Originally based on https://github.com/cznic/golex/blob/master/calc/tokenizer.l

package main

import (
  "bufio"
  "fmt"
)

type yylexer struct{
  src     *bufio.Reader
  buf     []byte
  empty   bool
  current byte
}

func newLexer(src *bufio.Reader) (y *yylexer) {
  y = &yylexer{src: src}
  if b, err := src.ReadByte(); err == nil {
    y.current = b
  }
  return
}

func (y *yylexer) getc() byte {
  if y.current != 0 {
    y.buf = append(y.buf, y.current)
  }
  y.current = 0
  if b, err := y.src.ReadByte(); err == nil {
    y.current = b
  }
  return y.current
}

func (y yylexer) Error(e string) {
  fmt.Printf("%v: %q\n", e, y.current)
}

func (y *yylexer) Lex(lval *yySymType) int {
  c := y.current
  if y.empty {
    c, y.empty = y.getc(), false
  }
%}

%yyc c
%yyn c = y.getc()

DIGIT             [0-9]
NUMBER            {DIGIT}+

%%
  // Code executed before every scan cycle
  y.buf = y.buf[:0]

[ \t\n\r]+        // Ignore whitespace

{NUMBER}
  lval.value = string(y.buf)
  return NUMBER

[;]
  lval.value = string(y.buf)
  return SCOLON

%%
  // Code executed at the end of lexing process.
  y.empty = true
  return int(c)
}