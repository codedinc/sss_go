%{
package main

import (
  "fmt"
)
%}

%union{
  value string
}

%token <value> DIMENSION
%token <value> NUMBER
%token <value> COLOR
%token <value> IDENTIFIER
%token <value> SELECTOR

%%

// Rules

stylesheet:
  statements
;

statements:
  statement
| statements ';' statement
;

statement:
  NUMBER                        { fmt.Printf("NUMBER: %q\n", $1) }
| IDENTIFIER                    { fmt.Printf("IDENTIFIER: %q\n", $1) }
;
