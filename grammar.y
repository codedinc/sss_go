%{
package main
%}

%union{
  value      string
  values     []string
  node       Node
  rule       Rule
  rules      []Rule
  property   Property
  properties []Property
}

%token <value> DIMENSION
%token <value> NUMBER
%token <value> COLOR
%token <value> IDENTIFIER
%token <value> SELECTOR

// Declare return value (in %union) type of rules
%type <node> stylesheet
%type <rules> rules
%type <rule> rule
%type <properties> properties
%type <property> property
%type <value> selector value
%type <values> values

%%

// Rules

stylesheet:                         // HACK go yacc can't return a custom value. We store it in the lexer instead.
                                    // https://groups.google.com/forum/#!topic/golang-dev/nemcZF_KyYg
  rules                             { yylex.(*Lexer).output = StyleSheet{$1} }
;

rules:
  rule                              { $$ = []Rule{$1} }
| rules rule                        { $$ = append($1, $2) }
;

rule:
  selector '{' properties '}'       { $$ = Rule{$1, $3} }
;

selector:
  SELECTOR
| IDENTIFIER
;

properties:
  property                          { $$ = []Property{$1} }
| properties ';' property           { $$ = append($1, $3) }
| properties ';'                    { $$ = $1 }
;

property:
  IDENTIFIER ':' values             { $$ = Property{$1, $3} }
;

values:
  value                             { $$ = []string{$1} }
| values value                      { $$ = append($1, $2) }
;

value:
  IDENTIFIER
| DIMENSION
| COLOR
;
