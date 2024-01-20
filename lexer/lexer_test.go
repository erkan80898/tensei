package lexer

import (
	"testing"

	"tensei/token"
)

func TestNextToken(t *testing.T) {
	input := `let num = 10.3;
num += 10;
num = num ** 2;

let secnum = num // 2;
secnum++;

let loop = fn(z) {
  while(z <= 1000){
    if (z == 2){
      break;
    }
    continue;
  }
};
let result = add(num, secnum);

let name = "Erk" + string(result);
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "num"},
		{token.ASSIGN, "="},
		{token.FLOAT, "10.3"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "num"},
		{token.ASSP, "+="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "num"},
		{token.ASSIGN, "="},
		{token.IDENT, "num"},
		{token.POWER, "**"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "secnum"},
		{token.ASSIGN, "="},
		{token.IDENT, "num"},
		{token.INTDIV, "//"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "secnum"},
		{token.INC, "++"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "loop"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "z"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.WHILE, "while"},

		{token.LPAREN, "("},
		{token.IDENT, "z"},
		{token.LEQ, "<="},
		{token.INT, "1000"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "z"},
		{token.EQ, "=="},
		{token.INT, "2"},

		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.BREAK, "break"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.CONTINUE, "continue"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "num"},
		{token.COMMA, ","},
		{token.IDENT, "secnum"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "name"},
		{token.ASSIGN, "="},
		{token.STRING, "Erk"},
		{token.PLUS, "+"},
		{token.IDENT, "string"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			println(tok.Type)
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
