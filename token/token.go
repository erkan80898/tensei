package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"
	FLOAT = "FLOAT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="
	LEQ    = "<="
	GEQ    = ">="

	ASSP   = "+="
	ASSM   = "-="
	ASSMUL = "*="
	ASSDIV = "/="
	INTDIV = "//"
	POWER  = "**"
	INC    = "++"
	DEC    = "--"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELIF     = "ELIF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	WHILE    = "WHILE"
	FOR      = "FOR"
	DO       = "DO"
	CONTINUE = "CONTINUE"
	BREAK    = "BREAK"
	SWITCH   = "SWITCH"
	TYPE     = "TYPE"
	STRUCT   = "STRUCT"
	TYPEOF   = "TYPEOF"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"elif":     ELIF,
	"else":     ELSE,
	"return":   RETURN,
	"while":    WHILE,
	"for":      FOR,
	"do":       DO,
	"continue": CONTINUE,
	"break":    BREAK,
	"switch":   SWITCH,
	"type":     TYPE,
	"struct":   STRUCT,
	"typeof":   TYPEOF,
}

// return either ident or keyword
func LookupIdent(x string) TokenType {
	if tok, ok := keywords[x]; ok {
		return tok
	}
	return IDENT
}
