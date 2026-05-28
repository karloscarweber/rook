// rook/ast.go

package ast

import (
	// "bytes"
	// "rook/token"
	// "rook/lexer"
)

type NodeType int

const (
	StatementNode NodeType = iota
	ExpressionNode
	ModuleNode
	TypeNode
	// ReturnNode
	// TypeNode
	// FunctionNode
	// IdentifierNode
	// NumberNode
)

var nodeTypes = map[NodeType]string {
	StatementNode: "statement",
	ExpressionNode: "expression",
	ModuleNode: "module",
	TypeNode: "type",
	// ReturnNode: "return",
	// TypeNode: "type",
	// FunctionNode: "function",
	// IdentifierNode: "identifier",
	// NumberNode: "number",
}

func (nt NodeType) String() string {
	return nodeTypes[nt]
}

type Node struct {
	Type NodeType
	Value string
	Children []Node
}

// returns a new Module Node
func Module(val string) Node {
	n := Node{Type: ModuleNode, Value: val}
	return n
}

// func (s *Statement) Type() token.Type {
// 	return s.Token.Type
// }

// expressionTypes
// const (
// 	EX_STRING     = "STRING"
// 	EX_NUMBER     = "NUMBER"
// 	EX_IDENTIFIER = "IDENTIFIER"
// )

// type Expression struct {
// 	Token token.Token
// 	// Type string
// 	// Type() string

// 	// str string
// 	// numInt int
// 	// numFloat float
// }

// func (e *Expression) Type() token.Type {
// 	return e.Token.Type
// }

// // returns a new Module
// func Identifier(tok token.Token) Expression {
// 	s := Expression{Token: tok}
// 	return s
// }

// type Module struct {
// 	FilePath   string
// 	Statements []Statement
// 	Nodes []Node
// }


// type ModuleStatement struct {
// 	Token token.Token
// }

// Statements
// Types of statements
//
// const statements create a constant, an unchangabel value.
// type ConstStatement struct {
// 	Token token.Token
// 	Name *Identifier
// 	Value Expression
// }

// func (cs *ConstStatement) isStatement()         { }
// func (cs *ConstStatement) TokenLiteral() string { return cs.Token.Literal }
// func (cs *ConstStatement) String() string {
// 	var out bytes.Buffer

// 	out.WriteString(cs.TokenLiteral() + " ")
// 	out.WriteString(cs.Name.String())
// 	out.WriteString(" = ")

// 	if cs.Value != nil {
// 		out.WriteString(cs.Value.String())
// 	}

// 	out.WriteString("\n")

// 	return out.String()
// }

// type Identifier struct {
// 	Literal string
// }
// func (id *Identifier) String() string {
// 	var out bytes.Buffer

// 	out.WriteString(id.String())

// 	return out.String()
// }

// LetStatement
//
// When we declare variables or constants.
// type LetStatement struct {
// 	Token token.Token
// 	Name  *Identifier
// 	Value Expression
// }

// func (ls *LetStatement) statementNode()       {}
// func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
// func (ls *LetStatement) String() string {
// 	var out bytes.Buffer

// 	out.WriteString(ls.TokenLiteral() + " ")
// 	out.WriteString(ls.Name.String())
// 	out.WriteString(" = ")

// 	if ls.Value != nil {
// 		out.WriteString(ls.Value.String())
// 	}

// 	out.WriteString(";")

// 	return out.String()
// }


// type Node interface {
// 	TokenLiteral() string
// 	String() string
// }

// type Statement interface {
// 	Node
// 	statementNode()
// }

// type Expression interface {
// 	Node
// 	expressionNode()
// }

// type Program struct {
// 	FilePath   string
// 	Statements []Statement
// }

// // Program.String() string
// //
// // Returns a string with all the statements just plopped out there.
// func (p *Program) String() string {
// 	var out bytes.Buffer

// 	for _, s := range p.Statements {
// 		out.WriteString(s.String())
// 	}

// 	return out.String()
// }

// // Statements

// // LetStatement
// //
// // When we declare variables or constants.
// type LetStatement struct {
// 	Token token.Token
// 	Name  *Identifier
// 	Value Expression
// }

// func (ls *LetStatement) statementNode()       {}
// func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
// func (ls *LetStatement) String() string {
// 	var out bytes.Buffer

// 	out.WriteString(ls.TokenLiteral() + " ")
// 	out.WriteString(ls.Name.String())
// 	out.WriteString(" = ")

// 	if ls.Value != nil {
// 		out.WriteString(ls.Value.String())
// 	}

// 	out.WriteString(";")

// 	return out.String()
// }

// type ReturnStatement struct {
// 	Token       token.Token
// 	ReturnValue Expression
// }

// func (rs *ReturnStatement) statementNode()       {}
// func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
// func (rs *ReturnStatement) String() string {
// 	var out bytes.Buffer

// 	out.WriteString(rs.TokenLiteral() + "  ")

// 	if rs.ReturnValue != nil {
// 		out.WriteString(rs.ReturnValue.String())
// 	}

// 	out.WriteString(";")

// 	return out.String()
// }

// // BlockStatement
// type BlockStatement struct {
// 	Token      token.Token
// 	Statements []Statement
// }

// func (bs *BlockStatement) expressionNode()      {}
// func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
// func (bs *BlockStatement) String() string {
// 	var out bytes.Buffer

// 	for _, s := range bs.Statements {
// 		out.WriteString(s.String())
// 	}

// 	return out.String()
// }

// // ExpressionStatement
// type ExpressionStatement struct {
// 	Token      token.Token
// 	Expression Expression
// }

// func (es *ExpressionStatement) statementNode()       {}
// func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
// func (es *ExpressionStatement) String() string {
// 	if es.Expressio != nil {
// 		return es.Expression.String()
// 	}
// 	return ""
// }

// type Identifier struct {
// 	Token token.Token
// 	Value string
// }

// func (i *Identifier) expressionNode()      {}
// func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
// func (i *Identifier) String() string       { return i.Value }

// type IntegerLiteral struct {
// 	Token token.Token
// 	Value int64
// }

// func (il *IntegerLiteral) expressionNode()      {}
// func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
// func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// type FloatLiteral struct {
// 	Token token.Token
// 	Value int64
// }

// func (il *FloatLiteral) expressionNode()      {}
// func (il *FloatLiteral) TokenLiteral() string { return il.Token.Literal }
// func (il *FloatLiteral) String() string       { return il.Token.Literal }

// type PrefixExpression struct {
// 	Token    token.Token
// 	Operator string
// 	Right    Expression
// }

// func (pe *PrefixExpression) expressionNode()      {}
// func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
// func (pe *PrefixExpression) String() string {
// 	var out bytes.Buffer

// 	out.WriteString("(")
// 	out.WriteString(pe.Operator)
// 	out.WriteString(pe.Right.String())
// 	out.WriteString(")")

// 	return out.String()
// }

// // PrefixExpression
// // InfixExpression
// // Boolean
// // IfExpression
// // FunctionLiteral
// // CallExpression
// // StringLiteral
// // ArrayLiteral
// // IndexExpression
// // HashLiteral


// ast structs
//
// Node
// 		type (the type of the node)
// 		value (left node of a node type) (value only referst to a literal, like a string node has a literal value or a number node has a literal value. All other nodes will have children or arguments.)
// 		children	(potential arguments, or members)
// 		right expression (an expression or the right node, could be a value. refers to another node)
// right expression is the first child?
//
//
// type
// value
// children


// Focus on S Expressions.


// Module Statement
// module <modulename
//
// ```
// module rook
// ```
//
// resulting AST S Expression Shape:
//
// (module "rook")


// type LetStatement struct {
// 	Token token.Token // the token.LET token
// 	Name  *Identifier
// 	Value Expression
// }

// type ReturnStatement struct {
// 	Token       token.Token // the 'return' token
// 	ReturnValue Expression
// }

// type ExpressionStatement struct {
// 	Token      token.Token // the first token of the expression
// 	Expression Expression
// }

// type BlockStatement struct {
// 	Token      token.Token // the { token
// 	Statements []Statement
// }

// type Identifier struct {
// 	Token token.Token // the token.IDENT token
// 	Value string
// }

// type Boolean struct {
// 	Token token.Token
// 	Value bool
// }

// type IntegerLiteral struct {
// 	Token token.Token
// 	Value int64
// }

// type PrefixExpression struct {
// 	Token    token.Token // The prefix token, e.g. !
// 	Operator string
// 	Right    Expression
// }

// type InfixExpression struct {
// 	Token    token.Token // The operator token, e.g. +
// 	Left     Expression
// 	Operator string
// 	Right    Expression
// }

// type IfExpression struct {
// 	Token       token.Token // The 'if' token
// 	Condition   Expression
// 	Consequence *BlockStatement
// 	Alternative *BlockStatement
// }

// type FunctionLiteral struct {
// 	Token      token.Token // The 'fn' token
// 	Parameters []*Identifier
// 	Body       *BlockStatement
// }

// type CallExpression struct {
// 	Token     token.Token // The '(' token
// 	Function  Expression  // Identifier or FunctionLiteral
// 	Arguments []Expression
// }

// type StringLiteral struct {
// 	Token token.Token
// 	Value string
// }

// type ArrayLiteral struct {
// 	Token    token.Token // the '[' token
// 	Elements []Expression
// }

// type IndexExpression struct {
// 	Token token.Token // The [ token
// 	Left  Expression
// 	Index Expression
// }

// type HashLiteral struct {
// 	Token token.Token // the '{' token
// 	Pairs map[Expression]Expression
// }
