package parser

import (
	"fmt"
	"tensei/ast"
	"tensei/lexer"
	"tensei/token"
)

type Parser struct {
	l      *lexer.Lexer
	cur    token.Token
	looka  token.Token
	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) RetrieveErrors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("next token=%s, expected=%s", p.looka.Type, t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.cur = p.looka
	p.looka = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curIs(token.EOF) {
		s := p.parseStatement()
		if s != nil {
			program.Statements = append(program.Statements, s)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.cur.Type {
	case token.LET:
		return p.parseLet()
	default:
		return nil
	}
}

func (p *Parser) parseLet() *ast.LetStatement {
	s := &ast.LetStatement{Token: p.cur}

	if !p.peekAndStepIf(token.IDENT) {
		return nil
	}

	s.Lhs = &ast.Identifier{Token: p.cur, Val: p.cur.Literal}

	if !p.peekAndStepIf(token.ASSIGN) {
		return nil
	}

	//Skipping expression till ;
	//ie let x = 2 + 3;
	//is let x = 2
	//todo come back to handle them
	for !p.curIs(token.SEMICOLON) {
		p.nextToken()
	}

	return s
}

func (p *Parser) curIs(t token.TokenType) bool {
	return p.cur.Type == t
}

func (p *Parser) lookaIs(t token.TokenType) bool {
	return p.looka.Type == t
}

func (p *Parser) peekAndStepIf(t token.TokenType) bool {
	if p.lookaIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
