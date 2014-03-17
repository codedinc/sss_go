%{
package main

import (
  "fmt"
)
%}

%union{
  value string
}

%token <value> NUMBER
%token <value> IDENTIFIER

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
