grammar Vlang; 


// === Axioma principal ===
programa : declaraciones* | expresion EOF ;
// /home/sebas/Desktop/Compiladores 2/OLC2_EVJUNIO2025/Clase2/compiler/errors/error_strategy.go
declaraciones : varDcl
              // | funcDcl 
              // | structDcl 
              // |    
              | stmt    
              ; 

varDcl
    : 'mut' ID (ASSIGN expresion)?  #variableDeclaration
    ; 



stmt : expresion          #expresionStatement
     | 'print(' expresion ')' #printStatement
     | sentencias_control    #controlStatement
     ; 

sentencias_control
    : ifDcl             #if_context 
    | forDcl            #for_context 
    | whileDcl          #while_context
    ;

ifDcl
    : 'if' LPAREN expresion RPAREN LBRACK declaraciones* RBRACK ( 'else' LBRACK declaraciones* RBRACK )? 
    ;
forDcl
    : 'for' LPAREN ID ASSIGN expresion COMMA expresion RPAREN LBRACK declaraciones* RBRACK 
    ;
whileDcl
    : 'while' LPAREN expresion RPAREN LBRACK declaraciones* RBRACK 
    ;

// === Reglas de expresiones ===
expresion
    : valor                                                #valorexpr         
    | LPAREN expresion RPAREN                              #parentesisexpre
    | LBRACK expresion RBRACK                              #corchetesexpre
    | op=(NOT | MINUS) expresion                           #unario
    | expresion op=(MUL | DIV | MOD) expresion             #multdivmod
    | expresion op=(PLUS | MINUS) expresion                #sumres
    | expresion op=(LT | LE | GE | GT) expresion           #relacionales
    | expresion op=(EQ | NEQ) expresion                    #igualdad
    | expresion OR expresion                               #or
    | ID                                                   #id              
    | incredecre                                           #incredecr      
    | ID DOT ID                                            #expdotexp1             
    | ID DOT expresion                                     #expdotexp      
    | ID ASSIGN expresion                                  #asignacionfor
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



// === Incremento / Decremento ===
incredecre
    : ID INC    #incremento
    | ID DEC    #decremento
    ;

// === Tokens de palabras clave ===
LEN     : 'len' ;
CAP     : 'cap' ;
APPEND  : 'append' ;

// === Literales ===
BOOLEANO : 'true' | 'false' ;
ENTERO   : [0-9]+ ;
DECIMAL  : [0-9]+ '.' [0-9]+ ;
CADENA   : '"' (~["\\] | '\\' .)* '"' ;
CARACTER : '\'' . '\'' ;

// === Identificadores ===
ID : [a-zA-Z_][a-zA-Z0-9_]* ;

// === Operadores ===
PLUS    : '+' ;
MINUS   : '-' ;
MUL     : '*' ;
DIV     : '/' ;
MOD     : '%' ;
NOT     : '!' ;
OR      : '||' ;
EQ      : '==' ;
NEQ     : '!=' ;
LT      : '<' ;
LE      : '<=' ;
GT      : '>' ;
GE      : '>=' ;
ASSIGN  : ':=' ;
INC     : '++' ;
DEC     : '--' ;

// === Símbolos ===
LPAREN  : '(' ;
RPAREN  : ')' ;
LBRACK  : '[' ;
RBRACK  : ']' ;
DOT     : '.' ;
COMMA   : ',' ;

// === Espacios y comentarios ===
WS : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;