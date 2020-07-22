package lexer

import (
	"errors"
	"fmt"
	"mizar/log"
	"mizar/utils"

	"github.com/sirupsen/logrus"
)

type transaction struct {
	id     int
	queue  *utils.Queue // 该条事务持有的token队列
	parent *transaction
}

func newTransaction(id int, parent *transaction) *transaction {
	return &transaction{id: id, parent: parent, queue: utils.NewQueue()}
}

type Lexer struct {
	input              *Input
	currentTransaction *transaction
	queue              *utils.Queue
}

var TokenEofErr = errors.New("token eof")

func NewLexer(source string) *Lexer {
	input := newInput(source)
	return &Lexer{input: input, queue: utils.NewQueue()}
}

func (lexer *Lexer) NextToken() (token *Token, err error) {
	// 如果栈中有则优先从栈中pop出
	if lexer.queue.Len() > 0 {
		var tokenElement interface{}
		tokenElement, err = lexer.queue.Pop()
		if err != nil {
			err = fmt.Errorf("lexer.currentTransaction.tempQueue.Pop() error: [%w]", err)
			return
		}

		var b bool
		token, b = tokenElement.(*Token)
		if !b {
			err = errors.New("lexer.currentTransaction.tempQueue.Pop() error: interface convert to *Token error")
		} else {
			lexer.currentTransaction.queue.Push(token)
			log.Trace(logrus.Fields{
				"token": token,
			}, "lexer.NextToken pop token")
		}

		return
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
	if err == nil && token != nil && lexer.currentTransaction != nil {
		lexer.currentTransaction.queue.Push(token)
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
		keywordRunesLen := len(keywordRunes)
		runes, err = lexer.input.lookahead(keywordRunesLen)
		if err != nil {
			if err == inputEofErr {
				continue
			}
			return
		}
		if string(keyword) == string(runes) {
			lexer.input.advance(keywordRunesLen)
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
			lexer.input.back(1)
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
	if nil == lexer.currentTransaction {
		lexer.currentTransaction = newTransaction(0, nil)
	} else {
		lexer.currentTransaction = newTransaction(lexer.currentTransaction.id+1, lexer.currentTransaction)
	}
	log.Trace(logrus.Fields{
		"transaction": lexer.currentTransaction,
	}, "lexer.Begin")
}

func (lexer *Lexer) Rollback() {
	log.Trace(logrus.Fields{
		"transaction": lexer.currentTransaction,
	}, "lexer.Rollback")

	lexer.queue.PushQueue(lexer.currentTransaction.queue)

	lexer.currentTransaction = lexer.currentTransaction.parent
}

func (lexer *Lexer) Commit() {
	log.Trace(logrus.Fields{
		"transaction": lexer.currentTransaction,
	}, "lexer.Commit")
	if lexer.currentTransaction.parent != nil {
		lexer.currentTransaction.parent.queue.PushQueue(lexer.currentTransaction.queue)
	}
	lexer.currentTransaction = lexer.currentTransaction.parent
}

// 将token还回去
func (lexer *Lexer) Return(token *Token) {
	log.Trace(logrus.Fields{
		"transaction": lexer.currentTransaction,
		"token":       token,
	}, "lexer.Return")
	lexer.queue.Push(token)
	lexer.currentTransaction.queue.Pop()
}
