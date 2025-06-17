# 🌟 Gramática del Lenguaje (BNF) 🌟

---

## 🎯 Axioma Principal

```bnf
<programa> ::= <declaraciones>* EOF

<declaraciones> ::= <varDcl>
                 | <stmt>
                 | <funcDcl>
                 | <funcMain>
                 | <structDcl>

<funcMain> ::= 'fn' 'main' '(' ')' <block>

<funcDcl> ::= 'fn' IDENTIFIER '(' <parametrosFormales>? ')' <TIPO>? <block>

<block> ::= '{' <declaraciones>* '}'

<varDcl> ::= 'mut' IDENTIFIER <TIPO>? ('=' <expresion>)?
           | 'mut' IDENTIFIER <sliceTipo>
           | IDENTIFIER '=' IDENTIFIER '{' <listaAsignaciones> '}'
           | IDENTIFIER '=' <sliceTipo> <sliceInit>
           | IDENTIFIER '=' IDENTIFIER
           | IDENTIFIER ('=' <expresion>)?
           | IDENTIFIER '=' <CASTEOS> '(' <expresion> ')'
           | IDENTIFIER '[' <expresion> ']' '=' <expresion>

<sliceTipo> ::= '[' ']' <TIPO>

<sliceInit> ::= '{' <listaExpresiones>? '}'

<TIPO> ::= 'int'
        | 'float64'
        | 'string'
        | 'bool'
        | 'rune'

<stmt> ::= 'print' '(' (<expresion> (',' <expresion>)*)? ')'
         | <expresion>
         | <sentencias_control>
         | <sentencias_transferencia>

<sentencias_control> ::= <ifDcl>
                      | <forDcl>
                      | <switchDcl>
                      | <whileDcl>

<sentencias_transferencia> ::= 'break'
                            | 'continue'
                            | 'return' <expresion>?

<CASTEOS> ::= 'Atoi'
           | 'parseFloat'

<ifDcl> ::= 'if' <expresion> '{' <declaraciones>* '}' <elseIfDcl>* <elseCondicional>?

<elseIfDcl> ::= 'else' 'if' <expresion> '{' <declaraciones>* '}'

<elseCondicional> ::= 'else' '{' <declaraciones>* '}'

<forDcl> ::= 'for' <asignacion> ';' <expresion> ';' <stmt>? <block>
           | 'for' <expresion> <block>
           | 'for' IDENTIFIER ',' IDENTIFIER '=' 'range' IDENTIFIER <block>

<asignacion> ::= IDENTIFIER '=' <expresion>

<switchDcl> ::= 'switch' <expresion> '{' <caseBlock>* <defaultBlock>? '}'

<caseBlock> ::= 'case' <expresion> ':' <declaraciones>*

<defaultBlock> ::= 'default' ':' <declaraciones>*

<llamadaFuncion> ::= 'indexOf' '(' (<expresion> (',' <expresion>)*)? ')'
                  | 'join' '(' (<expresion> (',' <expresion>)*)? ')'
                  | IDENTIFIER '(' (<expresion> (',' <expresion>)*)? ')'
                  | 'len' '(' (<expresion> (',' <expresion>)*)? ')'
                  | 'append' '(' (<expresion> (',' <expresion>)*)? ')'

<funcCall> ::= IDENTIFIER '(' <parametrosReales>? ')'

<parametrosFormales> ::= <parametro> (',' <parametro>)*

<parametro> ::= IDENTIFIER <TIPO>

<parametrosReales> ::= <expresion> (',' <expresion>)*

<structDcl> ::= 'struct' IDENTIFIER '{' <atributosStruct> '}'

<atributosStruct> ::= <atributoStruct>+

<atributoStruct> ::= IDENTIFIER <TIPO>
                  | IDENTIFIER IDENTIFIER

<listaAsignaciones> ::= <asignacionStruct> (',' <asignacionStruct>)*

<asignacionStruct> ::= IDENTIFIER ':' <expresion>

<whileDcl> ::= 'while' '(' <expresion> ')' '[' <declaraciones>* ']'

<expresion> ::= <expresion> ('*' | '/' | '%') <expresion>
             | <expresion> ('+' | '-') <expresion>
             | <expresion> ('==' | '!=') <expresion>
             | ('!' | '-') <expresion>
             | <expresion> ('<' | '<=' | '>=' | '>') <expresion>
             | <expresion> ('&&' | '||') <expresion>
             | <valor>
             | '(' <expresion> ')'
             | '[' <expresion> ']'
             | IDENTIFIER '[' <expresion> ']'
             | <llamadaFuncion>
             | IDENTIFIER '.' IDENTIFIER '=' <expresion>
             | IDENTIFIER '{' <listaAsignaciones> '}'
             | IDENTIFIER
             | <incredecre>
             | IDENTIFIER '.' IDENTIFIER
             | IDENTIFIER '.' <expresion>
             | IDENTIFIER <TIPO> '=' <expresion>
             | IDENTIFIER ('+=' | '-=') <expresion>
             | 'typeOf' '(' '[' <expresion> (',' <expresion>)* ']' ')'
             | 'typeOf' '(' <expresion> ')'

<parametros> ::= <expresion> (',' <expresion>)*

<valores> ::= <valor>

<valor> ::= ENTERO
         | DECIMAL
         | CADENA
         | BOOLEANO
         | CARACTER

<listaExpresiones> ::= <expresion> (',' <expresion>)*

<incredecre> ::= IDENTIFIER '++'
              | IDENTIFIER '--'