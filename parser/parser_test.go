package parser

import (
	"tensei/ast"
	"tensei/lexer"
	"tensei/util"
	"testing"
)

func TestLets(t *testing.T) {
	input := `let x = 1;
  let y = 2 + x;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram return nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("Program statements length=%d, expected=%d", len(program.Statements), 2)
	}

	//ERRORS WITHIN PARSER
	err := p.RetrieveErrors()
	if len(err) != 0 {
		t.Errorf("Parsering complete with %d ERRORS", len(err))
		for i, msg := range err {
			t.Errorf("parse ERROR %d,%s%s", i, util.NEW_LINE+"\t", msg)
		}
		t.FailNow()
	}

	//NOTE-ERRORS COULD RESIDE IN LEXER
	//TODO change test to take into account non-string iden (such as idents that are expressions)
	//Validating parser results
	tests := []struct {
		expectedIdName string
	}{
		{"x"},
		{"y"},
	}

	for i, tt := range tests {
		curState := program.Statements[i]
		if !LetValid(t, curState, tt.expectedIdName) {
			//just return t will contain fail state from subcalls
			return
		}
	}
}

func LetValid(t *testing.T, state ast.Statement, expectedName string) bool {
	if state.TokLiteral() != "let" {
		t.Errorf("statement had=%s, expected=let", state.TokLiteral())
		return false
	}

	letState, ok := state.(*ast.LetStatement)
	if !ok {
		t.Errorf("letState type had=%T, expected=*ast.LetStatement", letState)
		return false
	}
	if letState.Lhs.Val != expectedName {
		t.Errorf("letState.Lhs.Val had=%s. expected=%s", letState.Lhs.Val, expectedName)
		return false
	}
	if letState.Lhs.TokLiteral() != expectedName {
		t.Errorf("letState.Lhs.TokLiteral() had=%s, expected=%s", letState.Lhs.TokLiteral(), expectedName)
		return false
	}
	return true
}
