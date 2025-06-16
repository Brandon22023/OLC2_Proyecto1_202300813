grammar Vlang; 


// === Axioma principal ===
programa : (declaraciones)* EOF ;

funcMain
    : 'fn' 'main' LPAREN RPAREN block
    ;

//Declaracion funciones
funcDcl
    : 'fn' ID LPAREN parametrosFormales? RPAREN (TIPO)? block                    
    ;

block
    : LBRACE declaraciones* RBRACE
    ;
// /home/sebas/Desktop/Compiladores 2/OLC2_EVJUNIO2025/Clase2/compiler/errors/error_strategy.go
declaraciones : varDcl   
              | stmt
              | funcDcl
              | funcMain
              | structDcl    
              ; 

varDcl
    : 'mut' ID (TIPO)? (ASSIGN expresion)?       #variableDeclaration
    | 'mut' ID sliceTipo                         #sliceEmptyDeclaration
    | ID ASSIGN ID LBRACE listaAsignaciones RBRACE  #structDirectInitDeclaration
    | ID ASSIGN sliceTipo sliceInit        #sliceInitDeclaration
    | ID ASSIGN ID                               #sliceAssignment
    | ID (ASSIGN expresion)?                     #variableDeclarationImmutable
    | ID ASSIGN CASTEOS LPAREN expresion RPAREN  #variableCastDeclaration
    | ID LBRACK expresion RBRACK ASSIGN expresion  #sliceAssignmentIndex
    ; 

sliceTipo
    : LBRACK RBRACK TIPO  // ejemplo: []int
    ;

sliceInit
    : LBRACE listaExpresiones? RBRACE  // ejemplo: []int{1, 2, 3}
    ;

TIPO
    : TIPO_ENTERO  
    | TIPO_DECIMAL 
    | TIPO_CADENA 
    | TIPO_BOOLEANO
    | TIPO_CHAR
    ;


stmt : PRINT LPAREN (expresion (COMMA expresion)*)? RPAREN #printStatement
     | expresion                    #expresionStatement
     | sentencias_control           #controlStatement
     | sentencias_transferencia     #transfersentence
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

CASTEOS
    : ATOI                       
    | PARSEFLOAT
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
    : FOR asignacion SEMICOLON expresion SEMICOLON (stmt)? block  #forClasico
    | FOR expresion block                                      #forCondicionUnica
    | FOR ID COMMA ID ASSIGN 'range' ID block                        #forRangeSlice
    ;
asignacion
    : ID ASSIGN expresion
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
    | LEN LPAREN (expresion (COMMA expresion)*)? RPAREN
    | APPEND LPAREN (expresion (COMMA expresion)*)? RPAREN
    ;



// Llamada funcion y parametros normales
funcCall
    : ID LPAREN parametrosReales? RPAREN
    ;

parametrosFormales
    : parametro (COMMA parametro)*
    ;

parametro 
    : ID TIPO
    ;

parametrosReales
    : expresion (COMMA expresion)*
    ;

// === Declaraciones de Structs ===
structDcl
    : 'struct' ID LBRACE atributosStruct RBRACE
    ;

atributosStruct
    : atributoStruct+
    ;

atributoStruct
    : ID TIPO                   # atributoPrimitivo
    | ID ID                     # atributoStructAnidado
    ;

listaAsignaciones
    : asignacionStruct (COMMA asignacionStruct)*
    ;

asignacionStruct
    : ID COLON expresion
    ;

//esto se tiene que eliminar porque while y do-while no existen en Vlang, solo se usa for
whileDcl
    : 'while' LPAREN expresion RPAREN LBRACK declaraciones* RBRACK 
    ;


// === Reglas de expresiones ===
expresion
    : expresion op=(MUL | DIV | MOD) expresion             #multdivmod
    | expresion op=(PLUS | MINUS) expresion                #sumres
    | expresion op=(EQ | NEQ) expresion                    #igualdad
    | op=(NOT | MINUS) expresion                           #unario
    |expresion op=(LT | LE | GE | GT) expresion           #relacionales
    | expresion op=(AND | OR) expresion                    #OPERADORESLOGICOS
    | valor                                                #valorexpr         
    | LPAREN expresion RPAREN                              #parentesisexpre
    | LBRACK expresion RBRACK                              #corchetesexpre
    | ID LBRACK expresion RBRACK                           #PARAPRINTSLICE
    | llamadaFuncion                                       #llamadaFuncionExpr
    | ID DOT ID ASSIGN expresion                           #structAttrAssign
    | ID LBRACE listaAsignaciones RBRACE                   #structInstanceCreation
    | ID                                                   #id              
    | incredecre                                           #incredecr      
    | ID DOT ID                                            #expdotexp1             
    | ID DOT expresion                                     #expdotexp      
    | ID TIPO ASSIGN expresion                             #asignacionLUEGO
    | ID op=(SUMAIMPLICITA | RESTOIMPLICITO) expresion     #IMCPLICIT
    | TYPEOF LPAREN LBRACK expresion (COMMA expresion)* RBRACK RPAREN         #casteo_paratipo_slice
    | TYPEOF LPAREN expresion RPAREN                      #casteo_paratipo   // <-- AGREGA ESTA LÍNEA
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
    // CASTEOS

// === Listas de expresiones ===
listaExpresiones : expresion (COMMA expresion)* ;

//FUNCIONES que parecen casteos
ATOI  : 'Atoi' ;
PARSEFLOAT : 'parseFloat' ;
TYPEOF : 'typeOf' ;


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