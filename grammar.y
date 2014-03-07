%{
package main

import (
  "fmt"
)
%}

%union{
  value string
}

%token NUMBER
%token SCOLON

%type <value> NUMBER

%%

// Rules

stylesheet:
  statements
;

statements:
  statement
| statements SCOLON statement
;

statement:
  NUMBER                        { fmt.Printf("%#v\n", $1) }
;