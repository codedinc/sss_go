package main

import (
	"bytes"
)

type StyleSheet struct {
	rules []*Rule
}

func (self *StyleSheet) ToCSS() string {
	buf := bytes.NewBuffer([]byte{})
	for _, rule := range self.rules {
		rule.ToCSS(buf)
	}
	return buf.String()
}

type Rule struct {
	selector   string
	properties []*Property
}

func (self *Rule) ToCSS(buf *bytes.Buffer) {
	buf.WriteString(self.selector)
	buf.WriteString(" {\n")
	for _, property := range self.properties {
		buf.WriteString("  ")
		property.ToCSS(buf)
		buf.WriteString(";\n")
	}
	buf.WriteString("}\n")
	return
}

type Property struct {
	name   string
	values []string
}

func (self *Property) ToCSS(buf *bytes.Buffer) {
	buf.WriteString(self.name)
	buf.WriteString(": ")
	for _, value := range self.values {
		buf.WriteString(value)
	}
	return
}
