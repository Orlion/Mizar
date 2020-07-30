package parser

import "mizar/lexer"

type ActionType int8

const (
	ActionTypeReduce ActionType = iota + 1
	ActionTypeShift
	ActionTypeAccept
)

// action表
type Action struct {
	T     ActionType
	State int
}

type ActionTable struct {
	m map[int]map[lexer.TokenType]*Action
}

func (at *ActionTable) getAction(e *Element, token *lexer.Token) *Action {
	if t2a, exists := at.m[state]; exists {
		if a, exists := t2a[tokenType]; exists {
			return a
		}
	}

	return nil
}
