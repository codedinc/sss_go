%{
package main

import (
  "io"
)
%}

%union{
  string     string
  strings    []string
  value      Value
  values     []Value
  rule       *Rule
  rules      []*Rule
  declaration   Value
  declarations []Value
}

%token <string> DIMENSION
%token <string> NUMBER
%token <string> COLOR
%token <string> IDENTIFIER
%token <string> SELECTOR
%token <string> VARIABLE

// Declare return value (in %union) type of rules
%type <rules> rules
%type <rule> rule
%type <declaration> declaration
%type <declarations> declarations
%type <property> property
%type <string> selector
%type <value> value
%type <values> values

%%

// Rules

stylesheet:                         // HACK go yacc can't return a custom value. We store it in the lexer instead.
                                    // https://groups.google.com/forum/#!topic/golang-dev/nemcZF_KyYg
  rules                             { yylex.(*Lexer).output = &StyleSheet{$1} }
;

rules:
  rule                              { $$ = []*Rule{$1} }
| rules rule                        { $$ = append($1, $2) }
;

rule:
  selector '{' declarations '}'       { $$ = &Rule{$1, $3} }
;

selector:
  SELECTOR
| IDENTIFIER
;

declarations:
  declaration                         { $$ = []Value{$1} }
| declarations ';' property           { $$ = append($1, $3) }
| declarations ';'                    { $$ = $1 }
;

declaration:
  property
|  variableDeclaration
;

property:
  IDENTIFIER ':' values             { $$ = &Property{$1, $3} }
;

values:
  value                             { $$ = []Value{$1} }
| values value                      { $$ = append($1, $2) }
;

value:
  IDENTIFIER                        { $$ = &Literal{$1}  }
| DIMENSION                         { $$ = &Literal{$1}  }
| COLOR                             { $$ = &Literal{$1}  }
| VARIABLE                          { $$ = &Variable{$1} }
;

variableDeclaration:
  VARIABLE ':' values               { $$ = &Assign($1, $3) }
;


%%

func Parse(reader io.Reader) *StyleSheet {
  // yyDebug = 1
  lexer := NewLexer(reader)
  yyParse(lexer)
  return lexer.output
}