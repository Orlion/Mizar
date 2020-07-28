package parser

import (
	"errors"
	"mizar/lexer"
)

type Parser struct {
	lexer *lexer.Lexer
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer}
}

var ExpectError = errors.New("expect error")
var EofError = errors.New("eof error") // 未parse完成但没有更多输入

// 自底向上分析
func (parser *Parser) Parse() (ast *TranslationUnit, err error) {
	return
}
