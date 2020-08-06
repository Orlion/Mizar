package lexer

import (
	"mizar/log"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNumber(t *testing.T) {
	log.Init(logrus.ErrorLevel)
	lexer := NewLexer("123")
	if token, err := lexer.number(); err != nil {
		t.Error(err)
	} else {
		if token.T != TokenIntLiteral {
			t.Error(token)
		}
	}

	lexer = NewLexer("123.0")
	if token, err := lexer.number(); err != nil {
		t.Error(err)
	} else {
		if token.T != TokenDoubleLiteral {
			t.Error(token)
		}
	}

	lexer = NewLexer("123.00000000001")
	if token, err := lexer.number(); err != nil {
		t.Error(err)
	} else {
		if token.T != TokenDoubleLiteral {
			t.Error(token)
		}
	}

	lexer = NewLexer("0")
	if token, err := lexer.number(); err != nil {
		t.Error(err)
	} else {
		if token.T != TokenIntLiteral {
			t.Error(token)
		}
	}

	lexer = NewLexer("0.1")
	if token, err := lexer.number(); err != nil {
		t.Error(err)
	} else {
		if token.T != TokenDoubleLiteral {
			t.Error(token)
		}
	}

	lexer = NewLexer("0.10")
	if token, err := lexer.number(); err != nil {
		t.Error(err)
	} else {
		if token.T != TokenDoubleLiteral {
			t.Error(token)
		}
	}

	lexer = NewLexer("01")
	if _, err := lexer.number(); err != nil {
		t.Error(err)
	}
}

func TestNextToken(t *testing.T) {
	log.Init(logrus.TraceLevel)
	lexer := NewLexer("123=123.12(00.0100-001")
	if token, err := lexer.NextToken(); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if token.T != TokenIntLiteral {
			t.Error(token)
			t.FailNow()
		}
	}

	if token, err := lexer.NextToken(); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if token.T != TokenAssign {
			t.Error(token)
			t.FailNow()
		}
	}

	if token, err := lexer.NextToken(); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if token.T != TokenDoubleLiteral {
			t.Error(token)
			t.FailNow()
		}
	}

	if token, err := lexer.NextToken(); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if token.T != TokenLp {
			t.Error(token)
			t.FailNow()
		}
	}
}
