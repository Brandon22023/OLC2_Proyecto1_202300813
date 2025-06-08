grammar Vlang; 


// === Axioma principal ===
programa : funcMain* EOF ;

funcMain
    : 'fn' 'main' LPAREN RPAREN block
    ;

block
    : LBRACE declaraciones* RBRACE
    ;
// /home/sebas/Desktop/Compiladores 2/OLC2_EVJUNIO2025/Clase2/compiler/errors/error_strategy.go
declaraciones : varDcl   
              | stmt    
              ; 

varDcl
    : 'mut' ID (TIPO)? (ASSIGN expresion)?  #variableDeclaration
    ; 


TIPO
    : TIPO_ENTERO  
    | TIPO_DECIMAL 
    | TIPO_CADENA 
    | TIPO_BOOLEANO
    | TIPO_CHAR
    ;

tipoSlice
    : LBRACK RBRACK TIPO
    ;

stmt : PRINT LPAREN (expresion (COMMA expresion)*)? RPAREN #printStatement
     | expresion          #expresionStatement
     | sentencias_control    #controlStatement
     | sentencias_transferencia  #transfersentence
     ; 

sentencias_control
    : ifDcl             #if_context 
    | forDcl            #for_context
    | switchDcl         #switch_context 
    | whileDcl          #while_context
    ;

sentencias_transferencia
     : BREAK                     #breakStatement
     | CONTINUE                  #continueStatement
     | RETURN (expresion)?       #returnStatement
     ;

ifDcl
    : IF expresion LBRACE declaraciones* RBRACE (elseIfDcl)* (elseCondicional)?
    ;


elseIfDcl
    : ELSE IF expresion LBRACE declaraciones* RBRACE
    ;

elseCondicional
    :ELSE LBRACE declaraciones* RBRACE
    ;

forDcl
    : FOR (LPAREN)? (stmt)? SEMICOLON expresion SEMICOLON (stmt)? block   #forClasico
    | FOR (LPAREN)? expresion (RPAREN)? block                             #forCondicionUnica
    ;

switchDcl
    : SWITCH expresion LBRACE caseBlock*  defaultBlock? RBRACE
    ;

caseBlock
    : 'case' expresion COLON declaraciones*
    ;

defaultBlock
    : 'default' COLON declaraciones*
    ;

llamadaFuncion
    : INDEXOF LPAREN (expresion (COMMA expresion)*)? RPAREN
    | JOIN LPAREN (expresion (COMMA expresion)*)? RPAREN
    | ID LPAREN (expresion (COMMA expresion)*)? RPAREN
    ;

//esto se tiene que eliminar porque while y do-while no existen en Vlang, solo se usa for
whileDcl
    : 'while' LPAREN expresion RPAREN LBRACK declaraciones* RBRACK 
    ;


// === Reglas de expresiones ===
expresion
    : expresion op=(LT | LE | GE | GT) expresion           #relacionales
    | expresion op=(EQ | NEQ) expresion                    #igualdad
    | expresion op=(AND | OR) expresion                    #OPERADORESLOGICOS
    | expresion op=(MUL | DIV | MOD) expresion             #multdivmod
    | expresion op=(PLUS | MINUS) expresion                #sumres
    | op=(NOT | MINUS) expresion                           #unario
    | valor                                                #valorexpr         
    | LPAREN expresion RPAREN                              #parentesisexpre
    | LBRACK expresion RBRACK                              #corchetesexpre
    | tipoSlice LBRACE listaExpresiones? RBRACE            #sliceCreacionv
    | llamadaFuncion                                       #llamadaFuncionExpr
    | ID                                                   #id              
    | incredecre                                           #incredecr      
    | ID DOT ID                                            #expdotexp1             
    | ID DOT expresion                                     #expdotexp      
    | ID TIPO ASSIGN expresion                             #asignacionLUEGO
    | ID op=(SUMAIMPLICITA | RESTOIMPLICITO) expresion     #IMCPLICIT
    ;

// === Parámetros en llamadas ===
parametros : expresion (COMMA expresion)* ;

// === Tipos de valores simples ===
valores : valor ;

// === Subcontextos para valores ===
valor
    : ENTERO    #valorEntero
    | DECIMAL   #valorDecimal
    | CADENA    #valorCadena
    | BOOLEANO  #valorBooleano
    | CARACTER  #valorCaracter
    ;
// === Listas de expresiones ===
listaExpresiones : expresion (COMMA expresion)* ;



// === Incremento / Decremento ===
incredecre
    : ID INC    #incremento
    | ID DEC    #decremento
    ;

// === Tokens de palabras clave ===
LEN     : 'len' ;
CAP     : 'cap' ;
APPEND  : 'append' ;
IF      : 'if' ;
ELSE    : 'else' ;
FOR     : 'for' ;
SWITCH  : 'switch' ;
INDEXOF : 'indexOf' ;
JOIN    : 'join' ;
BREAK   : 'break' ;
CONTINUE: 'continue' ;
RETURN  : 'return' ;
// === Literales ===
BOOLEANO : 'true' | 'false' ;
ENTERO   : [0-9]+ ;
DECIMAL  : [0-9]+ '.' [0-9]+ ;
CADENA   : '"' (~["\\] | '\\' .)* '"' ;
CARACTER : '\'' . '\'' ;
// === Palabras TIPO ===

TIPO_ENTERO : 'int' ;
TIPO_DECIMAL : 'float64' ;
TIPO_CADENA : 'string' ;
TIPO_BOOLEANO : 'bool' ;
TIPO_CHAR : 'rune' ;
//TIPO_SLICE : 'slice' ;
// === Identificadores ===
PRINT : 'print' ;
ID : [a-zA-Z_][a-zA-Z0-9_]* ;

// === Operadores ===
INC     : '++' ;
DEC     : '--' ;
SUMAIMPLICITA : '+=';
RESTOIMPLICITO : '-=' ;
PLUS    : '+' ;
MINUS   : '-' ;
MUL     : '*' ;
DIV     : '/' ;
MOD     : '%' ;
NOT     : '!' ;
OR      : '||' ;
AND     : '&&' ;
// === Operadores de comparación ===
EQ      : '==' ;
NEQ     : '!=' ;
LE      : '<=' ;
GE      : '>=' ;
LT      : '<' ;
GT      : '>' ;
ASSIGN  : '=' ;


// === Símbolos ===
LPAREN  : '(' ;
RPAREN  : ')' ;
LBRACK  : '[' ;
RBRACK  : ']' ;
LBRACE  : '{' ;
RBRACE  : '}' ;
SEMICOLON : ';' ;
COLON  : ':' ;
DOT     : '.' ;
COMMA   : ',' ;

// === Espacios y comentarios ===
WS : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;