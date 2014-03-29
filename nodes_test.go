package main

import "testing"

func TestBuildNodes(t *testing.T) {
	value := Literal{"#fff"}
	property := Property{"color", []Value{&value}}
	rule := Rule{"h1", []*Property{&property}}
	styleSheet := StyleSheet{[]*Rule{&rule}}

	_ = styleSheet // Avoid declared and not used error
}
