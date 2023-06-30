parser grammar PuppetoParser;

options { tokenVocab=PuppetoLexer; }

literalValue
    : FLOAT
    | INTEGER
    | STRING
    | IDENTIFIER
    | NIL
    | BOOLEAN
    | NAN
    ;

object
    : L_CURLY ( pair ( COMMA pair )* )? COMMA? R_CURLY
    ;

pair
    : ( L_SQUARE expression R_SQUARE | IDENTIFIER | STRING ) COLON expression
    | IDENTIFIER
    ;

array
    : L_SQUARE (expression (COMMA expression)* COMMA?)? R_SQUARE
    ;

functionDefinition
    : FN name=IDENTIFIER? L_PARENTHES argumentList? R_PARENTHES scopeBody
    ;

argumentList
    : argument (COMMA argument)* COMMA? (SPREAD rest=IDENTIFIER)?
    ;

argument
    : IDENTIFIER ((TO expression)? | none)
    ;

none
    : // Empty rule
    ;

scopeBody
    : expression
    | scope
    ;

value
    : literalValue
    | object
    | array
    | functionDefinition
    | L_PARENTHES expression R_PARENTHES
    ;

expression
    : logicalExpression
    ;

logicalExpression
    : equalityExpression ( ( AND | OR ) equalityExpression )*
    ;

equalityExpression
    : comparisonExpression ( ( EQ | NEQ ) comparisonExpression )*
    ;

comparisonExpression
    : additiveExpression ( ( LT | GT | LTOE | GTOE ) additiveExpression )*
    ;

additiveExpression
    : multiplicativeExpression ( ( PLUS | MINUS ) multiplicativeExpression )*
    ;

multiplicativeExpression
    : powerExpression ( ( STAR | SLASH | PERCENT ) powerExpression )*
    ;

powerExpression
    : bitwiseExpression ( DOUBLE_STAR bitwiseExpression )*
    ;

bitwiseExpression
    : shiftExpression ( ( BAND | BOR | BXOR ) shiftExpression )*
    ;

shiftExpression
    : unaryExpression ( ( BLS | BRS ) unaryExpression )*
    ;

unaryExpression
    : ( NOT | MINUS ) unaryExpression
    | chainExpression
    ;

getProperty
    : value ((L_SQUARE expression R_SQUARE) | (DOT IDENTIFIER))
    ;

setProperty
    : getProperty (TWO_SIDES_OPERATOR | TO) expression
    ;

callExpression
    : value L_PARENTHES expression (COMMA expression)* COMMA? R_PARENTHES
    ;

chainExpression
    : getProperty
    | setProperty
    | callExpression
    | value
    ;

lhsVariableAssignation
    : ONE_SIDE_OPERATOR IDENTIFIER
    ;

rhsVariableAssignation
    : IDENTIFIER ONE_SIDE_OPERATOR
    ;

midVariableAssignation
    : IDENTIFIER (TWO_SIDES_OPERATOR | TO) expression
    ;

variableAssignation
    :   
        ( midVariableAssignation
        | lhsVariableAssignation
        | rhsVariableAssignation
        )+
    ;

variableDefinition
    : LET variableDefinitionBody (COMMA variableDefinitionBody)*
    ;

variableDefinitionBody
    : IDENTIFIER (TO expression)?
    ;

ifStatement
    : IF expression scopeBody ( ELSE IF expression scopeBody )* ( ELSE scopeBody )?
    ;

loopStatement
    : FOR expression programBodyWithBreakContinue
    ;

switchStatement
    : SWITCH expression L_CURLY switchCase* ( DEFAULT COLON programBodyWithBreakContinue* )? R_CURLY
    ;

switchCase
    : CASE expression COLON programBodyWithBreakContinue
    ;

breakStatement
    : BREAK
    ;

continueStatement
    : CONTINUE
    ;

tryStatement
    : TRY programBody (CATCH IDENTIFIER? programBody)?
    ;

statementList
    : variableDefinition
    | functionDefinition
    | ifStatement
    | loopStatement
    | switchStatement
    | tryStatement
    | echoStatement
    | variableAssignation
    ;

statement
    : (statementList | (PUP_END_TAG htmlList* PUP_START_TAG)+)
    ;

scope
    : L_CURLY programBodyList? R_CURLY
    ;

scopeWithBreakContinue
    : L_CURLY programBodyWithBreakContinue* R_CURLY
    ;

programBody
    : scope
    | statement
    | expression
    ;

programBodyList
    : programBody+
    ;

echoStatement
    : ECHO expression
    ;

programBodyWithBreakContinue
    : scopeWithBreakContinue
    | statement
    | breakStatement
    | continueStatement
    | expression
    ;

pupCode
    : PUP_START_TAG programBodyList? PUP_END_TAG
    ;

htmlList
    : WS | LT | HTML | AND | BAND | BLS | BOOLEAN | BOR | BREAK | BRS | BXOR | CASE | CATCH | COLON | COMMA | CONTINUE | DEFAULT | DOT | DOUBLE_STAR | ECHO | ELSE | EQ | FLOAT | FN | FOR | IDENTIFIER | IF | INTEGER | L_CURLY | L_PARENTHES | L_SQUARE | LET | LTOE | MINUS | GT | GTOE | NEQ |  NIL | NOT | ONE_SIDE_OPERATOR | OR | PERCENT | PLUS | R_CURLY | R_PARENTHES | R_SQUARE | SEMICOLON | SLASH | STAR | STRING | SWITCH | TO | TRY | TWO_SIDES_OPERATOR
    ;

html
    : (pupCode | htmlList)+ EOF
    ;

program
    : PUP_START_TAG programBodyList? PUP_END_TAG? EOF
    ;