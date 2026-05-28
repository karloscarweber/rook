package parser

import (
	"rook/token"
	"rook/lexer"
	"rook/ast"
	"fmt"
	"bytes"
)

type (
	prefixFn func() ast.Node
	infixFn  func(ast.Node) ast.Node
)

// Parser
// makes an AST of tokens lexxed in the lexer.
// passes AST to Compiler.
type Parser struct {
	lexer      *lexer.Lexer

	// shortcuts to the actual token
	current    token.Token  // the current token
	peek 	   token.Token  // the next token

	position   int // The peek position.

	Errors     []string

	Nodes      []ast.Node // list of Nodes

	prefixFns  map[token.Type]prefixFn
	infixFns   map[token.Type]infixFn
}

// *Parser.New(input string)
//
// Creates a new Parser and then returns that.
func New(input string) *Parser {
	p := &Parser{lexer: lexer.New(input)}

	p.lexer.Tokenize()

	p.current = p.lexer.Tokens[0]
	p.position = 1
	p.peek = p.lexer.Tokens[p.position]

	p.registerFuncs()
	return p
}

// *Parser.registerFuncs()
//
// Registers the mapped functions for parsing.
func (p *Parser) registerFuncs() {

	p.prefixFns = make(map[token.Type]prefixFn)
	// MODULE
	// p.regsiterPrefixFunc(token.Type(token.MODULE), p.modulePrFN)
	// IDENT
	// TYPE
	// FUNCTION
	// IF
	// BANG
	// STRING
	// MINUS
	// TRUE
	// FALSE
	// LPAREN
	// LBRACKET
	// LBRACE

	p.infixFns = make(map[token.Type]infixFn)

}

func (p *Parser) regsiterPrefixFunc(tokenType token.Type, fn prefixFn) {
	p.prefixFns[tokenType] = fn
}

func (p *Parser) registerInfixFunc(tokenType token.Type, fn infixFn) {
	p.infixFns[tokenType] = fn
}

// *Parser.modulePrFN()
//
// Expects the pattern MODULE IDENTIFIER to declare
// a module statement.
// func (p *Parser) modulePrFN() ast.Node {

// 	// expect the next token to be an identifier
// 	if p.peek.Type == token.Type(token.IDENT) {
// 		// append(p.Nodes, p.current)
// 		// append(p.Nodes, p.peek)
// 		p.advance()
// 	} else {
// 		// error, unexpected token <>, Expecting an identifier.
// 		p.errors = append(p.errors, "Holy shit! ")
// 	}
// }

// *Parser.advance()
//
// advances the current and peek token forward.
func (p *Parser) advance() {
	if !p.isAtEnd() {
		p.position = p.position + 1
		p.current = p.peek
		p.peek = p.lexer.Tokens[p.position]
	}
}

// *Parser.Parse()
//
func (p *Parser) Parse() {
	tokens := p.lexer.Tokens

	// initialize tokens
	tok := tokens[0]
	for !p.isAtEnd() {

		// this switch checks the entry points to our grammer, which are the top level statements that we expect.
		switch tok.Type {
			// case token.ERROR:
				// This is an error token, what up?
			case token.MODULE:
				p.parseModule()
			case token.EOF:
				p.parseEOF()
			default:
				// this is an error
				p.tokenErrorMessage("Unrecognized token type: ", tok)
		}

		p.advance()
		tok = p.current
	}

	// for !l.IsAtEnd() {
	// 	l.Tokens = append(l.Tokens, l.NextToken())
	// }
	// begin parsing.x
}

func (p *Parser) parseModule() {
	moduleNode := ast.Module(p.peek.Literal)
	// expect the peeked token to be an ident
	if p.peek.Type != token.IDENT {
		p.tokenErrorMessage("Expecting module name, instead found ", p.peek)
	}

	p.Nodes = append(p.Nodes, moduleNode)
	p.advance()
}

func (p *Parser) parseEOF() {
	// We don't need to add the EOF token because we're making an ast.
}

// func (p *Parser) ParseProgram() *ast.Program {
// 	program := &ast.Program{}
// 	program.Statements = []ast.Statement{}
// 	for !p.curTokenIs(token.EOF) {
// 		stmt := p.parseStatement()
// 		if stmt != nil {
// 			program.Statements = append(program.Statements, stmt)
// 		}
// 		p.nextToken()
// 	}
// 	return program
// }

// parses a statement.
// func (p *Parser) parseStatement() ast.Statement {

// }

func (p *Parser) isAtEnd() bool {
	return p.position >= (len(p.lexer.Tokens) - 1)
}

func (p *Parser) error(message string) {
	p.Errors = append(p.Errors, message)
}

func (p *Parser) DebugErrors() {
	for _, element := range p.Errors {
		fmt.Println("ERROR: ", element)
	}
}

func (p *Parser) DebugTokens() {
	fmt.Println("__TOKENS__")
	for _, element := range p.lexer.Tokens {
		fmt.Print(element)
	}
}

func (p *Parser) tokenErrorMessage(message string, tok token.Token) {
	var out bytes.Buffer

	out.WriteString(message)
	out.WriteString(tok.String())

	// return out.String()
	p.Errors = append(p.Errors, out.String())
}
