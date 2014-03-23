package main

import "testing"

func TestBuildNodes(t *testing.T) {
	property := Property{"color", []string{"#fff"}}
	rule := Rule{"h1", []*Property{&property}}
	styleSheet := StyleSheet{[]*Rule{&rule}}

	_ = styleSheet // Avoid declared and not used error
}
