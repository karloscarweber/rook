// grammar.js
import * as ohm from 'ohm-js';

const strict = String.raw`
  Rook {
	Module = (FunctionDecl|ExternFunctionDecl|TypeDecl)*

	Statement = LetStatement
			  | IfStatement
			  | WhileStatement
			  | ExprStatement

	//+ "let x = 3 + 4;", "let distance = 100 + 2;"
	//- "let y;"
	LetStatement = let identifier "=" Expr ";"

	//+ "if x < 10 {}", "if z { 42; }", "if x {} else if y {} else { 42; }"
	//- "if x < 10 { 3 } else {}"
	IfStatement = if Expr BlockStatements (else (BlockStatements|IfStatement))?

	//+ "while 0 {}", "while x < 10 { x := x + 1; }"
	//- "while 1 { 42 }", "while x < 10 { x := x + 1 }"
	WhileStatement = while Expr BlockStatements

	//+ "func zero(): i32 { 0 }", "func add(x, y): i64 { x + y }"
	//- "func x", "func x();"
	FunctionDecl = func identifier "(" Params? ")" ":" identifier BlockExpr

	//+ "extern func print(x);"
	ExternFunctionDecl = extern func identifier "(" Params? ")" ";"

	//+ "String: (u32);"
	//+ "pointer: (u32);"
	//+ "tuple: (i32, i32);"
	//- "age (i32);"
	//- "slash ( )"
	//+ "morning: (i32, u32, f64);"
	//- "night: (i32, u32, );"
	TypeDecl = identifier ":" "(" Params? ")" ";"

	Params = identifier ("," identifier)*

	//+ "{ 42 }", "{ 66 + 99 }", "{ 1 + 2 - 3 }"
	//+ "{ let x = 3; 42 }"
	//- "{ 3abc }"
	BlockExpr = "{" Statement* Expr? "}"

	//+ "{}", "{ let x = 3; }", "{ 42; 99; }"
	//- "{ 42 }", "{ x := 1 }"
	BlockStatements = "{" Statement* "}"

	ExprStatement = Expr ";"

	Expr = AssignmentExpr  -- assignment
		  | PrimaryExpr (binaryOp PrimaryExpr)*  -- binary

	//+ "x := 3", "y := 2 + 1", "arr[x + 1] := 3"
	AssignmentExpr = identifier ":=" Expr  -- var
				   | identifier "[" Expr "]" ":=" Expr  -- array

	PrimaryExpr = "(" Expr ")"  -- paren
				| number
				| stringLiteral
				| CallExpr
				| identifier "[" Expr "]"  -- index
				| identifier  -- var
				| IfExpr

	CallExpr = identifier "(" Args? ")"

	Args = Expr ("," Expr)*

	//+ "if x { 42 } else { 99 }", "if x { 42 } else if y { 99 } else { 0 }"
	//- "if x { 42 }"
	IfExpr = if Expr BlockExpr else (BlockExpr|IfExpr)

	binaryOp = "+" | "-" | "*" | "/" | compareOp | logicalOp
	compareOp = "==" | "!=" | "<=" | "<" | ">=" | ">"
	logicalOp = and | or
	number = digit+

	keyword = if | else | func | let | while | and | or | extern | type
	if = "if" ~identPart
	else = "else" ~identPart
	func = "func" ~identPart
	let = "let" ~identPart
	while = "while" ~identPart
	and = "and" ~identPart
	or = "or" ~identPart
	extern = "extern" ~identPart
	type = "type" ~identPart

	//+ "x", "élan", "_", "_99"
	//- "1", "$nope"
	identifier = ~keyword identStart identPart*
	identStart = letter | "_"
	identPart = identStart | digit

	stringLiteral = quote (~quote any)* quote
	quote = "\""

	space += singleLineComment | multiLineComment
	singleLineComment = "//" (~"\n" any)*
	multiLineComment = "/*" (~"*/" any)* "*/"

	// Examples:
	//+ "func addOne(x): i32 { x + one }", "func one():i64 { 1 } func two():u32 { 2 }"
	//- "42", "let x", "func x {}"
  }
`;

const rook = ohm.grammar(strict);

const loose = String.raw`
	Crook <: Rook {
		Module := (FunctionDecl|ExternFunctionDecl|AllTypeDecls)*

		AllTypeDecls = (TypeDecl
			|TypeDeclNoLeftParen
			|TypeDeclNoRightParen
			|TypeDeclNoColo
			|TypeDeclNoSemi)

		//+ "String : (u32);"
		//+ "pointer: (u32);"
		//+ "tuple: (i32, i32);"
		//- "age (i32);"
		//- "slash ( )"
		//+ "morning: (i32, u32, f64);"
		//- "night: (i32, u32, );"
		//- "day: i32, u32);"
		//- "night: (i32, u32;"
		TypeDecl := TDStart "(" Params? ")" ";"
			TDStart = identifier ":"

		//+ "day: i32, u32);"
		//- "day: (i32, u32);"
		TypeDeclNoLeftParen = TDStart Params? ")" ";"

		//+ "night: (i32, u32;"
		//- "night: (i32, i32);"
		TypeDeclNoRightParen = TDStart "(" Params? ";"

		//+ "evening (i32, u32);"
		//- "evening: (i32, u32);"
		TypeDeclNoColo = identifier ~":" "(" Params? ")" ";"

		//+ "evening: (i32, u32)"
		//- "evening: (i32, u32);"
		TypeDeclNoSemi = TDStart "(" Params? ")" (~";" any)*

	}
`;

const crook = ohm.grammar(loose, {Rook: rook})

export {
	strict,
	loose,
	rook,
	crook,
}
