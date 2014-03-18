package main

import (
  "testing"
  "strings"
)

func TestLexEOF(t *testing.T) {
  lexer := newLexerFromString("")
  assertLexToToken(t, lexer, 0, "\x00")
}

func TestLexNumber(t *testing.T) {
  lexer := newLexerFromString("10")
  assertLexToToken(t, lexer, NUMBER, "10")
}

func TestLexIgnoreSpaces(t *testing.T) {
  lexer := newLexerFromString("10 20 30")
  assertLexToToken(t, lexer, NUMBER, "10")
  assertLexToToken(t, lexer, NUMBER, "20")
  assertLexToToken(t, lexer, NUMBER, "30")
}

func TestLexIgnoreLineBreaks(t *testing.T) {
  lexer := newLexerFromString("10\n20")
  assertLexToToken(t, lexer, NUMBER, "10")
  assertLexToToken(t, lexer, NUMBER, "20")
}

func TestLexIdentifier(t *testing.T) {
  lexer := newLexerFromString("bold")
  assertLexToToken(t, lexer, IDENTIFIER, "bold")
}

func TestLexSingleChar(t *testing.T) {
  lexer := newLexerFromString("{ }.")
  assertLexToToken(t, lexer, int('{'), "{")
  assertLexToToken(t, lexer, int('}'), "}")
  assertLexToToken(t, lexer, int('.'), ".")
}


// Helpers

func newLexerFromString(input string) (*Lexer) {
  return newLexer(strings.NewReader(input))
}

func assertLexToToken(t *testing.T, lexer *Lexer, expectedType int, expectedValue string) {
  var lval yySymType
  tokType := lexer.Lex(&lval)
  
  if tokType != expectedType {
    t.Errorf("Bad token type. Got %v, expected %v", tokType, expectedType)
  }

  if lval.value != expectedValue {
    t.Errorf("Bad token value. Got %q, expected %q", lval.value, expectedValue)
  }
}