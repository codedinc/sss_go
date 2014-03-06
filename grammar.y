%{
package main

import (
  "fmt"
)

%}

%token NUMBER

%%

// Rules

stylesheet:
  /* empty */
  statements
;

statements:
  statement
| statements statement
;

statement:
  NUMBER
;