package ast

import (
	"fmt"
	"tensei/token"
	"tensei/util"
)

type Node interface {
	Display() string
	TokLiteral() string
}

type Statement interface {
	Node
	//sugar: used to double check we meant to use statement
	statementEnforce()
}

type Expression interface {
	Node
	//sugar: used to double check we meant to use expression
	expressionEnforce()
}

type Program struct {
	//special node - enforce display
	Node
	Statements []Statement
}

func (p *Program) display() string {
	str := ""
	for _, s := range p.Statements {
		str += s.Display() + util.NEW_LINE
	}
	return str
}

func (p *Program) TokLiteral() string {
	return "PROGRAM"
}

type LetStatement struct {
	Statement
	Token token.Token
	Lhs   *Identifier
	Rhs   Expression
}

func (s *LetStatement) Display() string {
	//expression not implemented yet
	// return fmt.Sprintf("%s:%slhs: %s, rhs %s", s.TokLiteral(), util.NEW_LINE+"\t", s.Lhs.Display(), s.Rhs.Display())
	return fmt.Sprintf("%s:%s\tlhs: %s%s\trhs: (todo)", s.TokLiteral(), util.NEW_LINE, s.Lhs.Display(), util.NEW_LINE)
}

func (s *LetStatement) TokLiteral() string {
	if s == nil {
		return ""
	}
	return s.Token.Literal
}

// "enforce" to me that this is a statemnt
func (s *LetStatement) statementEnforce() {}

// ident = expression since rhs can be an expression
type Identifier struct {
	Expression
	Token token.Token
	//TODO extend rhs for expression not just string
	Val string
}

func (e *Identifier) Display() string {
	return fmt.Sprintf("%s[%s]", token.IDENT, e.Val)
}

func (e *Identifier) TokLiteral() string {
	return e.Token.Literal
}

func (e *Identifier) expressionEnforce() {}
