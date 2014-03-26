package main

import (
	"testing"
)

func TestCompileRule(t *testing.T) {
	assertCompilesToSameCSS(t, "h1 {\n  height: 10px;\n}\n")
}

func TestCompileRules(t *testing.T) {
	assertCompilesToSameCSS(t, "h1 {\n  height: 10px;\n}\n"+
		"p {\n  width: 10px;\n  display: block;\n}\n")
}

func assertCompilesToSameCSS(t *testing.T, expected string) {
	actual := CompileString(expected)
	if expected != actual {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}
