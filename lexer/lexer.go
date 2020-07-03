package lexer

import (
	"errors"
)

type Lexer struct {
	input *Input
}

var TokenEofErr = errors.New("token eof")

func NewLexer(source string) *Lexer {
	input := newInput(source)
	return &Lexer{input}
}

func (lexer *Lexer) NextToken() (token *Token, err error) {
	r, err := lexer.input.nextRune()
	if err != nil {
		if err == inputEofErr {
			err = TokenEofErr
		}
		return
	}

	rStr := string([]rune{r})

	if ` ` == rStr || 9 == r {
		// 忽略空格
		return lexer.NextToken()
	}

	if `"` == rStr {
		return lexer.string()
	} else if r >= '1' && r <= '9' {
		lexer.input.back(1)
		return lexer.uint()
	} else if "=" == rStr {
		token = new(Token)
		token.T = TokenTypeAssign
		token.V = rStr
	} else if "+" == rStr {
		token = new(Token)
		token.T = TokenTypeAdd
		token.V = rStr
	} else if "-" == rStr {
		token = new(Token)
		token.T = TokenTypeSub
		token.V = rStr
	} else if "{" == rStr {
		token = new(Token)
		token.T = TokenTypeLc
		token.V = rStr
	} else if "}" == rStr {
		token = new(Token)
		token.T = TokenTypeRc
		token.V = rStr
	} else if "(" == rStr {
		token = new(Token)
		token.T = TokenTypeLp
		token.V = rStr
	} else if ")" == rStr {
		token = new(Token)
		token.T = TokenTypeRp
		token.V = rStr
	} else if ";" == rStr {
		token = new(Token)
		token.T = TokenTypeSemicolon
		token.V = rStr
	} else {
		// 回退一个字符
		lexer.input.back(1)
		token, err = lexer.keyword()
		if err != nil {
			// 如果不是关键字则尝试识别标识符
			token, err = lexer.identifier()
			if err != nil {
				err = errors.New("不识别的字符")
			}
		}
	}
	return
}

func (lexer *Lexer) string() (token *Token, err error) {
	var v []rune
	var r rune
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
		token.V = string(v)
		token.T = TokenTypeString
	} else {
		err = errors.New("不识别的字符")
	}

	return
}

func (lexer *Lexer) uint() (token *Token, err error) {
	var v []rune
	var r rune
	for {
		r, err = lexer.input.nextRune()
		if err != nil {
			if err == inputEofErr {
				err = nil
				break
			}
			return
		}

		if r >= '0' && r <= '9' {
			v = append(v, r)
		} else {
			break
		}
	}

	if len(v) >= 1 {
		token = new(Token)
		token.V = string(v)
		token.T = TokenTypeUint
	} else {
		err = errors.New("不识别的字符")
	}

	return
}

func (lexer *Lexer) keyword() (token *Token, err error) {
	var runes []rune
	for _, keyword := range Keywords {
		keywordRunes := []rune(keyword)
		runes, err = lexer.input.lookahead(len(keywordRunes))
		if err != nil {
			if err == inputEofErr {
				continue
			}
			return
		}
		if string(keyword) == string(runes) {
			token = new(Token)
			token.V = string(keyword)
			token.T = keyword
			return
		}
	}

	err = errors.New("不识别的字符")

	return
}

func (lexer *Lexer) identifier() (token *Token, err error) {
	var v []rune
	var r rune

	r, err = lexer.input.nextRune()
	if err != nil {
		if err == inputEofErr {
			err = TokenEofErr
		}
		return
	}

	// 标识符必须以字母下划线开始
	if r != '_' && r > 'Z' && r < 'A' && r < 'a' && r > 'z' {
		err = errors.New("不识别的字符")
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
			break
		}
	}

	if len(v) >= 1 {
		token = new(Token)
		token.V = string(v)
		token.T = TokenTypeUint
	} else {
		err = errors.New("不识别的字符")
	}

	return
}
