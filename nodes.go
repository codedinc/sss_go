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
	values []Value
}

func (self *Property) ToCSS(buf *bytes.Buffer) {
	buf.WriteString(self.name)
	buf.WriteString(": ")
	for _, value := range self.values {
		value.ToCSS(buf)
	}
	return
}

type Value interface {
	ToCSS(buf *bytes.Buffer)
}

type Literal struct {
	value string
}

type Variable struct {
	name string
	//What goes here Value or Literal?
}

func (self *Variable) ToCSS(buf *bytes.Buffer) {
	//Dummy function
	return
}

func (self *Literal) ToCSS(buf *bytes.Buffer) {
	buf.WriteString(self.value)
	return
}
