package parser

import (
	"errors"
	"fmt"
	"mizar/ast"
	"mizar/lexer"
	"mizar/utils"
)

type Parser struct {
	lexer         *lexer.Lexer
	stateStack    *utils.Stack
	symbolStack   *utils.Stack
	valueStack    *utils.Stack
	lrActionTable map[int]map[Symbol]*Action
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer, utils.NewStack(), utils.NewStack(), utils.NewStack(), nil}
}

var ExpectError = errors.New("expect error")
var EofError = errors.New("eof error") // 未parse完成但没有更多输入

// 自底向上分析
func (parser *Parser) Parse() (ast *ast.TranslationUnit, err error) {
	var (
		currentState int
		ok           bool
		action       *Action
		args         []interface{}
	)

	gsm := newGrammarStateManager()
	gsm.build()

	parser.lrActionTable = gsm.getLRStateTable()

	token, lexerErr := parser.lexer.NextToken()
	if lexerErr != nil {
		err = fmt.Errorf("lexer error: [%w]", lexerErr)
		return
	}

	valueToken := token

	currentSymbol := Symbol(token.T)
	parser.stateStack.Push(0)
	parser.valueStack.Push(token)

	for {
		currentState, ok = parser.stateStack.Top().(int)
		if !ok {
			panic("stateStack.(int) error")
		}

		action, err = parser.getAction(currentState, currentSymbol)
		if err != nil {
			break
		}

		if action.isReduce { // 做reduce操作
			args = []interface{}{}

			// 根据生成式右侧符号的数量从栈中pop出该数量个符号
			for i := len(action.reduceProduction.right); i > 0; i-- {
				parser.symbolStack.Pop()
				parser.stateStack.Pop()
				args = append(args, parser.valueStack.Pop())
			}

			// 反转参数数组
			for i := 0; i < len(args)/2; i++ {
				temp := args[i]
				args[i] = args[len(args)-1-i]
				args[len(args)-1-i] = temp
			}

			// 将生成式左侧的非终结符压入到符号栈
			parser.symbolStack.Push(action.reduceProduction.left)

			currentSymbol = action.reduceProduction.left

			// reduce操作意味着能够产生新的抽象语法树节点
			// 将新节点压入值堆栈
			parser.valueStack.Push(action.reduceProduction.buildFunc(args))
		} else { // 做shift操作
			// 转移之后的状态压入到状态栈
			parser.stateStack.Push(action.shiftStateNum)

			// 将符号压入到符号栈
			parser.symbolStack.Push(currentSymbol)

			if currentSymbol.isTerminals() {
				parser.valueStack.Push(valueToken)

				// 如果当前符号是终结符，则需要移进下一个符号
				token, lexerErr := parser.lexer.NextToken()
				if lexerErr != nil {
					err = fmt.Errorf("lexer error: [%w]", lexerErr)
					return
				}
				currentSymbol = Symbol(token.T)
				valueToken = token
			} else {
				token = parser.lexer.GetCurrentToken()
				if token == nil {
					panic("token=nil")
				}
				currentSymbol = Symbol(token.T)
				valueToken = token
			}
		}
	}

	return
}

func (parser *Parser) getAction(currentState int, symbol Symbol) (action *Action, err error) {
	jump, exists := parser.lrActionTable[currentState]
	if !exists {
		err = errors.New("no jump")
		return
	}

	action, exists = jump[symbol]
	if !exists {
		err = errors.New("no action")
		return
	}

	return
}
