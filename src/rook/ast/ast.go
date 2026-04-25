// rook/ast.go

package ast

import (
	"bytes"
	"rook/token"
	"rook/lexer"
)

// Node
//

type Node interface {
	// Token is on the structs
	String() string
	Children() []Node
}

type Statement interface {
	Node
	isStatement()
}

type Expression interface {
	Node
	isExpression()
}

type Program struct {
	FilePath   string
	Statements []Statement
	Nodes []Node
}


// Statements
// Types of statements
//
// const statements create a constant, an unchangabel value.
type ConstStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (cs *ConstStatement) isStatement()         { }
func (cs *ConstStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.String())
	out.WriteString(" = ")

	if cs.Value != nil {
		out.WriteString(cs.Value.String())
	}

	out.WriteString("\n")

	return out.String()
}

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
