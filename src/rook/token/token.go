package token

type TokenType string

// A token is a type that represents either a single caharacter or a group of
// characters that have atomic significance in Rook.
type Token struct {
	Type    TokenType
	Literal string
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

var keywords = map[string]TokenType{
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
func Lookup(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
