package main

type StyleSheet struct {
  rules []Rule
}

type Rule struct {
  selector string
  properties []Property
}

type Property struct {
  name string
  value string
}