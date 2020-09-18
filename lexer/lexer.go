package lexer

import (
	"errors"
	"fmt"
	"mizar/log"

	"github.com/Orlion/merak/lexer"
	"github.com/sirupsen/logrus"
)

type Lexer struct {
	input *Input
}

var TokenEofErr = errors.New("token eof")
var TokenUnknownErr = errors.New("不识别的字符")

func NewLexer(source string) *Lexer {
	input := newInput(source)
	return &Lexer{input: input}
}

func (lexer *Lexer) Next() (token lexer.Token, err error) {
	if lexer.input.isEof() {
		t := new(Token)
		t.T = EoiToken
		t.StartColumn = lexer.input.ColumnNum
		t.EndColumn = lexer.input.ColumnNum
		t.StartLine = lexer.input.LineNum
		t.EndLine = lexer.input.LineNum

		token = t
		return
	}

	// 首先匹配保留字
	token, err = lexer.reservedWords()
	if err == nil {
		log.Trace(logrus.Fields{
			"token": token,
			"err":   err,
		}, "lexer.NextToken output token")
		return
	}

	r, err := lexer.input.nextRune()
	if err != nil {
		if err == inputEofErr {
			err = TokenEofErr
		}
		return
	}

	if 9 == r || r == 10 || r == 13 || r == ' ' {
		// 忽略空格换行
		return lexer.Next()
	}

	if '"' == r {
		token, err = lexer.string()
	} else if r >= '0' && r <= '9' {
		lexer.input.back(1)
		token, err = lexer.number()
	} else {
		// 回退一个字符
		lexer.input.back(1)
		// 如果不是关键字则尝试识别标识符
		token, err = lexer.identifier()
	}

	log.Trace(logrus.Fields{
		"r":     r,
		"token": token,
		"err":   err,
	}, "lexer.NextToken output token")

	return
}

func (lexer *Lexer) string() (token *Token, err error) {
	var v []rune
	var r rune

	startColumn := lexer.input.ColumnNum
	startLine := lexer.input.LineNum

	for {
		r, err = lexer.input.nextRune()
		rStr := string([]rune{r})
		if err != nil {
			if err == inputEofErr {
				err = nil
				break
			}
			return
		}

		if `"` == rStr {
			break
		}

		v = append(v, r)
	}

	if len(v) >= 1 {
		token = new(Token)
		token.Lexeme = string(v)
		token.T = TokenStringLiteral
		token.StartColumn = startColumn
		token.StartLine = startLine
		token.EndColumn = lexer.input.ColumnNum
		token.EndLine = lexer.input.LineNum
	} else {
		err = TokenUnknownErr
	}

	return
}

func (lexer *Lexer) number() (token *Token, err error) {
	var v []rune
	var r rune

	startColumn := lexer.input.ColumnNum
	startLine := lexer.input.LineNum

	// 识别状态机
	state := 0
State:
	for {
		r, err = lexer.input.nextRune()
		if err != nil {
			if err == inputEofErr {
				err = nil
				break State
			}
			return
		}

		switch state {
		case 0:
			if '0' <= r && '9' >= r {
				state = 1
			} else {
				lexer.input.back(1)
				break State
			}
		case 1:
			if '.' == r {
				state = 2
			} else if '0' <= r && '9' >= r {
				state = 1
			} else {
				lexer.input.back(1)
				break State
			}
		case 2:
			if '0' <= r && '9' >= r {
				state = 2
			} else {
				lexer.input.back(1)
				break State
			}
		}

		v = append(v, r)
	}

	switch state {
	case 1:
		token = new(Token)
		token.T = TokenIntLiteral
		token.Lexeme = string(v)
		token.StartColumn = startColumn
		token.StartLine = startLine
		token.EndColumn = lexer.input.ColumnNum
		token.EndLine = lexer.input.LineNum
	case 2:
		token = new(Token)
		token.T = TokenDoubleLiteral
		token.Lexeme = string(v)
		token.StartColumn = startColumn
		token.StartLine = startLine
		token.EndColumn = lexer.input.ColumnNum
		token.EndLine = lexer.input.LineNum
	default:
		err = TokenUnknownErr
	}

	return
}

func (lexer *Lexer) reservedWords() (token *Token, err error) {
	var runes []rune

	startColumn := lexer.input.ColumnNum
	startLine := lexer.input.LineNum

	for _, reservedWord := range reservedWords {
		reservedWordRunes := []rune(reservedWord)
		reservedWordRunesLen := len(reservedWordRunes)
		runes, err = lexer.input.lookahead(reservedWordRunesLen)
		if err != nil {
			if err == inputEofErr {
				continue
			}
			return
		}
		if reservedWord == string(runes) {
			lexer.input.advance(reservedWordRunesLen)
			tokenT, exists := reservedWords2TokenTypeMap[reservedWord]
			if !exists {
				panic(fmt.Sprintf("reservedWords2TokenTypeMap 中不存在 %s", reservedWord))
			}
			token = new(Token)
			token.T = tokenT
			token.Lexeme = reservedWord
			token.StartColumn = startColumn
			token.StartLine = startLine
			token.EndColumn = lexer.input.ColumnNum
			token.EndLine = lexer.input.LineNum
			return
		}
	}

	err = TokenUnknownErr

	return
}

func (lexer *Lexer) identifier() (token *Token, err error) {
	var v []rune
	var r rune

	startColumn := lexer.input.ColumnNum
	startLine := lexer.input.LineNum

	r, err = lexer.input.nextRune()
	if err != nil {
		if err == inputEofErr {
			err = TokenEofErr
		}
		return
	}

	// 标识符必须以字母下划线开始
	if !(r == '_' || (r <= 'Z' && r >= 'A') || (r <= 'z' && r >= 'a') || (r >= '0' && r <= '9')) {
		err = errors.New("不识别的字符")
		lexer.input.back(1)
		return
	}

	v = append(v, r)

	for {
		r, err = lexer.input.nextRune()
		if err != nil {
			if err == inputEofErr {
				err = nil
				break
			}
			return
		}

		if r == '_' || (r <= 'Z' && r >= 'A') || (r <= 'z' && r >= 'a') || (r >= '0' && r <= '9') {
			v = append(v, r)
		} else {
			lexer.input.back(1)
			break
		}
	}

	if len(v) >= 1 {
		token = new(Token)
		token.Lexeme = string(v)
		token.T = TokenIdentifier
		token.StartColumn = startColumn
		token.StartLine = startLine
		token.EndLine = lexer.input.LineNum
		token.EndColumn = lexer.input.ColumnNum
	} else {
		err = TokenUnknownErr
	}

	return
}
