grammar V4LangGrammar;

program: dcl*;

dcl
    : varDcl SEMICOLON?
    | matrixDcl SEMICOLON?
    | sliceDcl SEMICOLON? 
    | statement SEMICOLON?
    | funcDcl SEMICOLON?
    | structDcl SEMICOLON?
    ;

varDcl
	: MUT ID type ASSIGN expr 
	| MUT ID type
	| MUT ID L_CURLY R_CURLY type // var numbers []int
	| ID DECLARE_ASSIGN expr
	| ID type
	| ID ID
	;

funcDcl
	: FUNC ID L_PAREN params? R_PAREN type? L_CURLY dcl* R_CURLY // func nombreFuncion (params) type { dcl* }
	| FUNC receiver ID L_PAREN params? R_PAREN type? L_CURLY dcl* R_CURLY // func (p Persona) nombreFuncion (params) type { dcl* }  
	;

params
	: ID (L_BRACKET R_BRACKET)? type (',' ID (L_BRACKET R_BRACKET)? type)*
	;

structDcl
	:  'type' ID STRUCT L_CURLY structBody* R_CURLY
	;

structBody
	: varDcl
	| funcDcl
	;

// methodDecl
// 	: FUNC receiver ID L_PAREN params? R_PAREN type? L_CURLY dcl* R_CURLY // func (p Persona) nombreFuncion (params) type { dcl* }  
// 	;

receiver
	: L_PAREN this = ID structName = ID R_PAREN
	;

sliceDcl 
	: MUT ID DECLARE_ASSIGN L_BRACKET R_BRACKET type L_CURLY sliceValues R_CURLY // numbers := []int {1,2,3}
	| ID type // slice[]int
	;

sliceValues
	: expr (',' expr)* // 1 | 1,2,3
	;

matrixDcl
	: ID DECLARE_ASSIGN L_BRACKET R_BRACKET L_BRACKET R_BRACKET type L_CURLY matrixValues R_CURLY // matrix := [][]int {{1,2,3},{4,5,6}}
	;

matrixValues
	: L_CURLY sliceValues R_CURLY (',' L_CURLY sliceValues R_CURLY)* ','? // {1,2,3}, {4,5,6}
	;

statement
	: expr  											# ExprStmt 
	| PRINT L_PAREN args? R_PAREN  						# PrintStmt
	| L_CURLY dcl* R_CURLY							    # BlockStmt
	| IF expr statement (elseIfStmt)* (elseStmt)? 		# IfStmt
	| 'switch' expr L_CURLY switchCase* defaultCase? R_CURLY 	# SwitchStmt
	| FOR varDcl ';' expr ';' expr statement			# ForIncrementStmt
	| FOR expr statement					    		# ForStmt
	| FOR ID ',' ID DECLARE_ASSIGN RANGE ID statement  # ForRangeStmt
	| 'break' 											# BreakStmt
	| 'continue' 										# ContinueStmt
	| 'return' expr?									# ReturnStmt
	; 

elseIfStmt
	: ELSE IF expr statement
	;

elseStmt
	: ELSE statement
	;

switchCase
	: 'case' expr ':' dcl*
	;

defaultCase
	: 'default' ':' dcl*
	;

args
	: expr (',' expr)*
	;

expr
	: '-' expr						# Negate
	| '!' expr						# Not
	| expr call+ 					# CallStmt
	| expr op = ('*' | '/') expr	# MulDiv
	| expr op = ('+' | '-') expr	# AddSub
	| expr op = ('<' | '>' | '<=' | '>=') expr # Relational
	| expr '%' expr					# Mod
	| expr op = ('==' | '!=') expr	# Equality
	| ID op = ('+=' | '-=' | '*=' | '/=') expr # AssignArithmetic
	| expr op = ('&&' | '||') expr	# Logical
	| 'new' ID L_PAREN args? R_PAREN # NewInstance
	| ID L_CURLY instanceValues R_CURLY # StructInstance // persona := Persona{Nombre: "Juan", Edad: 20}
	| ID L_BRACKET expr R_BRACKET L_BRACKET expr R_BRACKET ASSIGN expr # SetValueMatrix // matrix[1][2] = 3
	| ID L_BRACKET expr R_BRACKET L_BRACKET expr R_BRACKET? # GetValueMatrix // matrix[1][2] 
	| ID L_BRACKET expr R_BRACKET ASSIGN valor = expr # SetValueIndex // numbers[2] = 3
	| ID L_BRACKET expr R_BRACKET 		 # GetValueIndex // numbers[2]
	| LEN L_PAREN expr R_PAREN 		 # Len
	| JOIN L_PAREN ID ',' expr R_PAREN # StringJoin
	| ATOI L_PAREN expr R_PAREN 		 # Atoi
	| PARSEFLOAT L_PAREN expr R_PAREN  # ParseFloat
	| TYPEOF L_PAREN ID R_PAREN 	 # TypeOf
	| SLICE_INDEX L_PAREN ID ',' expr R_PAREN # SliceIndex
	| expr ASSIGN expr				# Assign
	| ID op = ('++' | '--')			# IncrementDecrement
	| ID							# Identifier
	| RUNE 							# Rune
	| BOOL							# Boolean
	| FLOAT							# Float
	| STRING						# String
	| INT							# Int
	| L_CURLY sliceValues* R_CURLY	# Slice
	| L_BRACKET R_BRACKET type L_CURLY sliceValues R_CURLY # SliceWithValues // []int{1,2,3}
	| APPEND L_PAREN ID ',' expr R_PAREN # Append 
	| NIL 							# Nil
	// | L_CURLY matrixValues R_CURLY    # Matrix
	| L_PAREN expr R_PAREN			# Parens
	;

call
	: L_PAREN args? R_PAREN # FuncCall 
	| '.' ID # Get
	;

instanceValues
	: ID ':' expr (',' ID ':' expr)* ','? // Nombre: "Juan", Edad: 20
	;



//Inicio Analizador léxico
//Tokens del analizador léxico
MUT: 'mut';	
PRINT: 'fmt.Println';
IF: 'if';
ELSE: 'else';
APPEND: 'append';
LEN: 'len';
SLICE_INDEX: 'slices.Index';
JOIN: 'strings.Join';
ATOI: 'strconv.Atoi';
PARSEFLOAT: 'strconv.ParseFloat';
TYPEOF: 'reflect.TypeOf';
FOR: 'for';
RANGE: 'range';
BREAK: 'break';
NIL: 'nil';
FUNC: 'func';
STRUCT: 'struct';


L_CURLY: '{';
R_CURLY: '}';
L_PAREN: '(';
R_PAREN: ')';
L_BRACKET: '[';
R_BRACKET: ']';
DECLARE_ASSIGN: ':=';
ASSIGN: '=';
SEMICOLON: ';';

INT: [0-9]+;
BOOL: 'true' | 'false';
FLOAT: [0-9]+'.'[0-9]+;
STRING: '"' (ESC|.)*? '"';
RUNE: '\'' ~["]* '\'';
WS: [ \t\r\n]+ -> skip;
ID: [a-zA-Z_0-9]+;
COMMENT: '//' .*? '\r'? '\n' -> skip;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;
type: 'int' | 'bool' | 'float64' | 'string' | 'rune';
fragment
ESC: '\\"' | '\\\\' | '\\/' | '\\n' | '\\r' | '\\t';