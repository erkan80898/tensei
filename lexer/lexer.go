package lexer

import (
	"tensei/token"
)

type Lexer struct {
	source    string
	cursor    int //points to ch
	lookahead int //looks ahead to determine action
	ch        byte
}

func New(src string) *Lexer {
	l := &Lexer{source: src, cursor: -1, lookahead: -1}
	l.step()
	return l
}

// let x = 5 + 5;
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.prepare()

	switch l.ch {
	case ',':
		tok = newtoken(token.COMMA, ",")
	case ';':
		tok = newtoken(token.SEMICOLON, ";")
	case '(':
		tok = newtoken(token.LPAREN, "(")
	case ')':
		tok = newtoken(token.RPAREN, ")")
	case '{':
		tok = newtoken(token.LBRACE, "{")
	case '}':
		tok = newtoken(token.RBRACE, "}")
	case '=':
		switch l.peek() {
		case '=':
			tok = newtoken(token.EQ, "==")
			l.step()
		default:
			tok = newtoken(token.ASSIGN, "=")
		}
	case '-':
		switch l.peek() {
		case '-':
			tok = newtoken(token.DEC, "--")
			l.step()
		case '=':
			tok = newtoken(token.ASSM, "-=")
			l.step()
		default:
			tok = newtoken(token.MINUS, "-")
		}
	case '+':
		switch l.peek() {
		case '+':
			tok = newtoken(token.INC, "++")
			l.step()
		case '=':
			tok = newtoken(token.ASSP, "+=")
			l.step()
		default:
			tok = newtoken(token.PLUS, "+")
		}
	case '!':
		switch l.peek() {
		case '=':
			tok = newtoken(token.NOT_EQ, "!=")
			l.step()
		default:
			tok = newtoken(token.BANG, "!")
		}
	case '*':
		switch l.peek() {
		case '*':
			tok = newtoken(token.POWER, "**")
			l.step()
		case '=':
			tok = newtoken(token.ASSMUL, "*=")
			l.step()
		default:
			tok = newtoken(token.ASTERISK, "*")
		}
	case '/':
		switch l.peek() {
		case '=':
			tok = newtoken(token.ASSDIV, "/=")
			l.step()
		case '/':
			tok = newtoken(token.INTDIV, "//")
			l.step()
		default:
			tok = newtoken(token.SLASH, "/")
		}
	case '>':
		switch l.peek() {
		case '=':
			tok = newtoken(token.GEQ, ">=")
			l.step()
		default:
			tok = newtoken(token.GT, ">")
		}
	case '<':
		switch l.peek() {
		case '=':
			tok = newtoken(token.LEQ, "<=")
			l.step()
		default:
			tok = newtoken(token.LT, "<")
		}
	case '|':
		switch l.peek() {
		case '|':
			tok = newtoken(token.OR, "||")
			l.step()
		default:
			tok = newtoken(token.ILLEGAL, wrapIllegalLiteral("Illegal Token- |"+string(l.peek())))
		}
	case '&':
		switch l.peek() {
		case '&':
			tok = newtoken(token.AND, "&&")
			l.step()
		default:
			tok = newtoken(token.ILLEGAL, wrapIllegalLiteral("Illegal Token- &"+string(l.peek())))
		}
	case '"':
		l.step()
		ttype, lit := l.readstring()
		return newtoken(ttype, lit)
	case 0:
		tok = newtoken(token.EOF, "")
	default:
		if isDigit(l.ch) {
			ttype, lit := l.readnumber()
			return newtoken(ttype, lit)
		} else if isLetter(l.ch) {
			ttype, lit := l.readident()
			return newtoken(ttype, lit)
		} else {
			tok = newtoken(token.ILLEGAL, wrapIllegalLiteral("Illegal Token- "+string(l.ch)))
		}
	}

	l.step()
	return tok
}

func (l *Lexer) step() {
	//eol?
	if l.cursor++; l.cursor >= len(l.source) {
		l.ch = 0
	} else {
		l.ch = l.source[l.cursor]
		l.lookahead = l.cursor + 1
	}
}

func (l *Lexer) peek() byte {
	return l.source[l.lookahead]
}

// after a token - jump whitespaces to prepare for next token
func (l *Lexer) prepare() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.step()
	}
}

func (l *Lexer) readnumber() (tok token.TokenType, literal string) {
	start := l.cursor
	tok = token.INT
	malFlag := false
	for isDigit(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			if tok == token.FLOAT {
				malFlag = true
			}
			tok = token.FLOAT
		}
		l.step()
	}

	if malFlag == true {
		return token.ILLEGAL, wrapIllegalLiteral("Malformed float: " + l.source[start:l.cursor])
	}

	literal = l.source[start:l.cursor]
	return tok, literal
}

func (l *Lexer) readident() (tok token.TokenType, literal string) {
	start := l.cursor
	for isLetter(l.ch) {
		l.step()
	}
	literal = l.source[start:l.cursor]
	tok = token.LookupIdent(literal)
	return tok, literal
}

func (l *Lexer) readstring() (tok token.TokenType, literal string) {
	start := l.cursor
	for l.ch != '"' {
		if l.ch == 0 {
			return token.ILLEGAL, wrapIllegalLiteral("No closing \": " + l.source[start:l.cursor])
		}
		l.step()
	}
	literal = l.source[start:l.cursor]
	//step past last ending "
	l.step()
	return token.STRING, literal
}

func newtoken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func wrapIllegalLiteral(msg string) string {
	return "\"" + msg + "\""
}
