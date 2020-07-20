package lexer

import (
	"container/list"
	"errors"
	"mizar/log"

	"github.com/sirupsen/logrus"
)

type Lexer struct {
	input           *Input
	openTransaction int // 是否开启事务，事务开启后会收集生成的token，写入到stack中
	tempQueue       *list.List
	queue           *list.List
}

var TokenEofErr = errors.New("token eof")

func NewLexer(source string) *Lexer {
	input := newInput(source)
	queue := list.New()
	tempQueue := list.New()
	return &Lexer{input: input, tempQueue: tempQueue, queue: queue}
}

func (lexer *Lexer) NextToken() (token *Token, err error) {
	// 如果栈中有则优先从栈中pop出
	if lexer.queue.Len() > 0 {
		e := lexer.queue.Front()
		lexer.queue.Remove(e)
		var b bool
		token, b = e.Value.(*Token)
		if b {
			if lexer.openTransaction > 0 {
				// 如果开启了事务则将token push到临时队列中
				lexer.tempQueue.PushBack(token)
			}
			log.Trace(logrus.Fields{
				"token": token,
			}, "lexer.NextToken pop token")
			return
		} else {
			log.Error(logrus.Fields{
				"e": e,
			}, "")
			err = errors.New("lexer queue pop error")
		}
	}

	r, err := lexer.input.nextRune()
	if err != nil {
		if err == inputEofErr {
			err = TokenEofErr
		}
		log.Error(logrus.Fields{
			"err": err,
		}, "lexer.NextToken err")
		return
	}

	rStr := string([]rune{r})

	if ` ` == rStr || 9 == r || "\n" == rStr {
		// 忽略空格
		return lexer.NextToken()
	}

	if `"` == rStr {
		token, err = lexer.string()
	} else if r >= '1' && r <= '9' {
		lexer.input.back(1)
		token, err = lexer.number()
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
	} else if "*" == rStr {
		token = new(Token)
		token.T = TokenTypeMul
		token.V = rStr
	} else if "/" == rStr {
		token = new(Token)
		token.T = TokenTypeDiv
		token.V = rStr
	} else if "," == rStr {
		token = new(Token)
		token.T = TokenTypeComma
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

	// 如果开启了事务则将token临时记录下来
	if err == nil && token != nil && lexer.openTransaction > 0 {
		lexer.tempQueue.PushBack(token)
	}

	log.Trace(logrus.Fields{
		"token": token,
	}, "lexer.NextToken output token")

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

func (lexer *Lexer) number() (token *Token, err error) {
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
			// 还回去
			lexer.input.back(1)
			break
		}
	}

	if len(v) >= 1 {
		token = new(Token)
		token.V = string(v)
		token.T = TokenTypeNumber
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
		token.T = TokenTypeIdentifier
	} else {
		err = errors.New("不识别的字符")
	}

	return
}

func (lexer *Lexer) Begin() {
	log.Trace(logrus.Fields{
		"openTransaction": lexer.openTransaction,
	}, "lexer.Begin")
	lexer.openTransaction++
}

func (lexer *Lexer) Rollback() {
	log.Trace(logrus.Fields{
		"openTransaction": lexer.openTransaction,
	}, "lexer.Rollback")
	lexer.openTransaction--
	// 将tempQueue中数据写入到queue
	lexer.queue.PushBackList(lexer.tempQueue)
	lexer.tempQueue.Init()
}

func (lexer *Lexer) Commit() {
	log.Trace(logrus.Fields{
		"openTransaction": lexer.openTransaction,
	}, "lexer.Commit")
	lexer.openTransaction--
	lexer.tempQueue.Init()
	if lexer.openTransaction <= 0 {
		lexer.queue.Init()
	}
}

// 将token还回去
func (lexer *Lexer) Return(token *Token) {
	log.Trace(logrus.Fields{
		"token": token,
	}, "lexer.Return")
	lexer.queue.PushBack(token)
}