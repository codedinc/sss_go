package main

import (
  "testing"
  "strings"
)

func TestLexEOF(t *testing.T) {
  lexer := NewLexerFromString("")
  assertLexToToken(t, lexer, 0, "\x00")
}

func TestLexNumber(t *testing.T) {
  lexer := NewLexerFromString("10")
  assertLexToToken(t, lexer, NUMBER, "10")
}

func TestLexDimension(t *testing.T) {
  lexer := NewLexerFromString("10px")
  assertLexToToken(t, lexer, DIMENSION, "10px")
}

func TestLexIgnoreSpaces(t *testing.T) {
  lexer := NewLexerFromString("10 20 30")
  assertLexToToken(t, lexer, NUMBER, "10")
  assertLexToToken(t, lexer, NUMBER, "20")
  assertLexToToken(t, lexer, NUMBER, "30")
}

func TestLexIgnoreLineBreaks(t *testing.T) {
  lexer := NewLexerFromString("10\n20")
  assertLexToToken(t, lexer, NUMBER, "10")
  assertLexToToken(t, lexer, NUMBER, "20")
}

func TestLexIdentifier(t *testing.T) {
  lexer := NewLexerFromString("bold")
  assertLexToToken(t, lexer, IDENTIFIER, "bold")
}

func TestLexSelector(t *testing.T) {
  lexer := NewLexerFromString("#id .class a:hover")
  assertLexToToken(t, lexer, SELECTOR, "#id")
  assertLexToToken(t, lexer, SELECTOR, ".class")
  assertLexToToken(t, lexer, SELECTOR, "a:hover")
}

func TestLexSingleChar(t *testing.T) {
  lexer := NewLexerFromString("{ }.")
  assertLexToToken(t, lexer, int('{'), "{")
  assertLexToToken(t, lexer, int('}'), "}")
  assertLexToToken(t, lexer, int('.'), ".")
}


// Helpers

func NewLexerFromString(input string) (*Lexer) {
  return NewLexer(strings.NewReader(input))
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