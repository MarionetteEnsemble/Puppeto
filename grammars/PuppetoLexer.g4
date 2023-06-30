lexer grammar PuppetoLexer;

PUP_START_TAG
    : '<?pup'
    ;
PUP_END_TAG
    : '?>'
    ;

TWO_SIDES_OPERATOR
    : '+='
    | '-='
    | '*='
    | '/='
    | '&='
    | '|='
    | '^='
    | '**='
    | '%='
    ;
ONE_SIDE_OPERATOR
    : '++'
    | '--'
    ;
EQ
    : '=='
    ;
NEQ
    : '!='
    ;
LTOE
    : '<='
    ;
MTOE
    : '>='
    ;

FOR
    : 'for'
    ;
TO
    : '='
    ;
ECHO
    : 'echo'
    ;
L_CURLY
    : '{'
    ;
R_CURLY
    : '}'
    ;
COMMA
    : ','
    ;
L_SQUARE
    : '['
    ;
R_SQUARE
    : ']'
    ;
FN
    : 'fn'
    ;
L_PARENTHES
    : '('
    ;
R_PARENTHES
    : ')'
    ;

AND
    : '&&'
    ;
OR
    : '||'
    ;
COLON
    : ':'
    ;
DOT
    : '.'
    ;
PLUS
    : '+'
    ;
MINUS
    : '-'
    ;
DOUBLE_STAR
    : '**'
    ;
PERCENT
    : '%'
    ;
BAND
    : '&'
    ;
BOR
    : '|'
    ;
BXOR
    : '^'
    ;
BLS
    : '<<'
    ;
BRS 
    : '>>'
    ;
NOT
    : '!'
    ;
IF
    : 'if'
    ;
ELSE
    : 'else'
    ;
SWITCH
    : 'switch'
    ;
CASE
    : 'case'
    ;
BREAK
    : 'break'
    ;
TRY
    : 'try'
    ;
CONTINUE
    : 'continue'
    ;
CATCH
    : 'catch'
    ;
DEFAULT
    : 'default'
    ;
LET
    : 'let'
    ;
STAR
    : '*'
    ;
SLASH
    : '/'
    ;

fragment DIGIT
    : [0-9]
    ;

FLOAT
    : '-'? DIGIT+ '.' DIGIT* 
    | '-'? '.' DIGIT+
    ;
INTEGER
    : DIGIT+
    ;

NIL
    : 'nil'
    ;

BOOLEAN
    : 'true'
    | 'false'
    ;

fragment IDENTIFIER_START
    : [A-Z]
    | [a-z] 
    | '_' 
    | '$'
    ;
fragment IDENTIFIER_END
    : IDENTIFIER_START 
    | DIGIT
    ;

IDENTIFIER
    : IDENTIFIER_START IDENTIFIER_END*
    ;

STRING
    : '\'' ( ~'\'' | '\\' . )* '\'' 
    | '"' ( ~'"' | '\\' . )* '"'
    ;

LT
    : '<'
    ;
MT
    : '>'
    ;


WS
    : [ \t\r\n]+ -> channel(HIDDEN)
    ;

SEMICOLON
    : ';'+ -> channel(HIDDEN)
    ;

HTML
    : '<' ~('?') ~('>')* '>' HTML*
    | ~('<')+?
    ;