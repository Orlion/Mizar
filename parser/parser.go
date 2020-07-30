package parser

import (
	"errors"
	"mizar/lexer"
	"mizar/utils"
)

type Parser struct {
	lexer      *lexer.Lexer
	stateStack *utils.Stack
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer, utils.NewStack()}
}

var ExpectError = errors.New("expect error")
var EofError = errors.New("eof error") // 未parse完成但没有更多输入

// 自底向上分析
func (parser *Parser) Parse() (ast *TranslationUnit, err error) {
	var (
		token  *lexer.Token
		eInter interface{}
		action *Action
	)

	actionTable := new(ActionTable)

	// 将状态0压入堆栈
	parser.stack.Push(0)

	for {
		token, err = parser.lexer.NextToken()
		if err != nil {
			return
		}

		eInter = parser.stack.Top()
		e = eInter.(*Element)

		// 根据state和token从Action表中获取下一步要进行的操作
		action = actionTable.getAction(e, token)
		// 如果没有对应操作则识别失败
		if action == nil {
			err = nil
			break
		}

		if action.T == "stateX" {
			parser.stack.Push(&Element{
				State: X,
			})
		} else if action.T == "reduceX" {
			// count := 表达式X右边元素数量
			parser.stack.Pop(count) // 从栈中弹出count个元素
			// left = 表达式X左边的非终结符
			state := actionTable.getAction(parser.stack.Top(), left)
			parser.stack.Push(state)
		} else if action.T == "accept" {
			break
		}
	}

	return
}
