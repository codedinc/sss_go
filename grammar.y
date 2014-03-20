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
  selector_statement
| selector_statement statements
;

selector_statement:
  SELECTOR    '{' properties '}'   { fmt.Printf("SELECTOR: %q\n", $1) }
| IDENTIFIER  '{' properties '}'   { fmt.Printf("IDENTIFIER: %q\n", $1) }
;

properties:
  property
| properties ';' property
| properties ';'
;

property:
  IDENTIFIER ':' IDENTIFIER     { fmt.Printf("IDENTIFIER: %q\n", $1) }
| IDENTIFIER ':' COLOR          { fmt.Printf("IDENTIFIER: %q\n", $1) }
| IDENTIFIER ':' dimensions     { fmt.Printf("IDENTIFIER: %q\n", $1) }
;

dimensions:
  DIMENSION
| dimensions DIMENSION
;