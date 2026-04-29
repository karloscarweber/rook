package token

import (
	"bytes"
)

type Type string

// A token is a type that represents either a single caharacter or a group of
// characters that have atomic significance in Rook.
type Token struct {
	Type    Type
	Literal string
	Start   int
	Length  int
}

func (tk Token) String() string {
	var out bytes.Buffer

	out.WriteString(string(tk.Type))
	// out.WriteString(tk.Literal() + " ")
	out.WriteString("[")
	out.WriteString(string(tk.Start))
	out.WriteString(":")
	out.WriteString(string(tk.Length))
	out.WriteString("] - `")
	out.WriteString(tk.Literal)
	out.WriteString("`")

	out.WriteString("\n")

	return out.String()
}


const (
	ERROR = "ERROR"
	TRAP  = "TRAP"
	EOF   = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	// AST Representations of numbers don't correspond directly to WASM
	// types. That data is carred elsewhere in the AST.
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"
	BYTE   = "BYTE" // A hexadecimal number of the pattern: 0xFF

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	AND = "&&"
	OR  = "||"

	AMPERSAND = "&"
	PIPE      = "|"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	PUBLIC   = "PUBLIC"
	TYPE     = "TYPE"
	LET      = "LET"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	SWITCH   = "SWITCH"
	CASE     = "CASE"
	RETURN   = "RETURN"
)

var keywords = map[string]Type{
	"fun":    FUNCTION,
	"pub":    PUBLIC,
	"type":   TYPE,
	"let":    LET,
	"const":  CONST,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"switch": SWITCH,
	"case":   CASE,
	"return": RETURN,
}

// looks up a keyword and returns either the keyword token or the
// identifier token.
func Lookup(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// func New(tokenType Type, lit string, start int, length int) Token {
// 	return Token{tokenType, lit, start, length}
// }
