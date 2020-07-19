package parser

import (
	"errors"
	"fmt"
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
func (parser *Parser) Parse() (*TranslationUnit, error) {
	return parser.translationUnit()
}

func (parser *Parser) translationUnit() (translationUnit *TranslationUnit, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	definitionOrStatement, err := parser.definitionOrStatement()
	if err != nil {
		return
	}

	translationUnit = new(TranslationUnit)
	translationUnit.DefinitionOrStatementList = append(translationUnit.DefinitionOrStatementList, definitionOrStatement)
	for {
		definitionOrStatement, err = parser.definitionOrStatement()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = nil
			}
			break
		}

		translationUnit.DefinitionOrStatementList = append(translationUnit.DefinitionOrStatementList, definitionOrStatement)
	}

	return
}

func (parser *Parser) definitionOrStatement() (definitionOrStatement *DefinitionOrStatement, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	functionDefinition, err := parser.functionDefinition()
	if err == nil {
		definitionOrStatement = &DefinitionOrStatement{
			FunctionDefinition: functionDefinition,
			T:                  "FunctionDefinition",
		}
	} else {
		statement, err := parser.statement()
		if err == nil {
			definitionOrStatement = &DefinitionOrStatement{
				Statement: statement,
				T:         "Statement",
			}
		}
	}

	return
}

func (parser *Parser) functionDefinition() (functionDefinition *FunctionDefinition, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	var token *lexer.Token
	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeFunc {
		err = fmt.Errorf("[%w] expect token:TokenTypeFunc", ExpectError)
		return
	}

	NameToken, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if NameToken.T != lexer.TokenTypeIdentifier {
		err = fmt.Errorf("[%w] expect TokenTypeIdentifier, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] expect TokenTypeLp, but %s", ExpectError, token.T)
		return
	}

	parameterList, err := parser.parameterList()
	if err != nil {
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeRp {
		err = fmt.Errorf("[%w] expect TokenTypeLp, but %s", ExpectError, token.T)
		return
	}

	block, err := parser.block()
	if err != nil {
		return
	}

	functionDefinition = &FunctionDefinition{
		ParameterList: parameterList,
		Block:         block,
		Name:          NameToken.V,
	}

	return
}

func (parser *Parser) parameterList() (parameterList *ParameterList, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeIdentifier {
		err = fmt.Errorf("[%w] expect TokenTypeIdentifier, but %s", ExpectError, token.T)
		return
	}

	parameterList = new(ParameterList)
	parameterList.ParameterList = append(parameterList.ParameterList, &Parameter{token.V})

	for {
		token, err = parser.lexer.NextToken()
		if err != nil {
			err = fmt.Errorf("[lexer error] %s", err)
			break
		}
		if token.T != lexer.TokenTypeComma {
			err = fmt.Errorf("[%w] expect TokenTypeComma, but %s", ExpectError, token.T)
			break
		}

		token, err = parser.lexer.NextToken()
		if err != nil {
			err = fmt.Errorf("[lexer error] %s", err)
			break
		}
		if token.T != lexer.TokenTypeIdentifier {
			err = fmt.Errorf("[%w] expect TokenTypeIdentifier, but %s", ExpectError, token.T)
			break
		}

		parameterList.ParameterList = append(parameterList.ParameterList, &Parameter{token.V})
	}

	return
}

func (parser *Parser) argumentList() (argumentList *ArgumentList, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	expression, err := parser.expression()
	if err != nil {
		return
	}

	argumentList = new(ArgumentList)
	argumentList.ExpressionList = append(argumentList.ExpressionList, expression)

	var token *lexer.Token
	for {
		token, err = parser.lexer.NextToken()
		if err != nil {
			err = fmt.Errorf("[lexer error] %s", err)
			break
		}
		if token.T != lexer.TokenTypeComma {
			err = fmt.Errorf("[%w] expect TokenTypeComma, but %s", ExpectError, token.T)
			break
		}

		expression, err = parser.expression()
		if err != nil {
			break
		}

		argumentList.ExpressionList = append(argumentList.ExpressionList, expression)
	}

	return
}

func (parser *Parser) block() (block *Block, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	var token *lexer.Token
	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeLc {
		err = fmt.Errorf("[%w] expect TokenTypeLc, but %s", ExpectError, token.T)
		return
	}

	var statementList *StatementList
	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeRc {
		statementList, err = parser.statementList()
		if err != nil {
			return
		}

		token, err = parser.lexer.NextToken()
		if err != nil {
			err = fmt.Errorf("[lexer error] %s", err)
			return
		}
		if token.T != lexer.TokenTypeRc {
			err = fmt.Errorf("[%w] expect TokenTypeRc, but %s", ExpectError, token.T)
			return
		}
	}

	block = new(Block)
	block.StatementList = statementList

	return
}

func (parser *Parser) statementList() (statementList *StatementList, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	staement, err := parser.statement()
	if err != nil {
		return
	}

	statementList = new(StatementList)
	statementList.StatementList = append(statementList.StatementList, staement)

	for {
		staement, err = parser.statement()
		if err != nil {
			break
		}

		statementList.StatementList = append(statementList.StatementList, staement)
	}

	return
}

func (parser *Parser) statement() (statement *Statement, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	whileStatement, err := parser.whileStatement()
	if err != nil {
		var breakStatement *BreakStatement
		breakStatement, err = parser.breakStatement()
		if err != nil {
			var continueStatement *ContinueStatement
			continueStatement, err = parser.continueStatement()
			if err != nil {
				var returnStatement *ReturnStatement
				returnStatement, err = parser.returnStatement()
				if err != nil {

				} else {
					statement = &Statement{
						T:               "ReturnStatement",
						ReturnStatement: returnStatement,
					}
				}
			} else {
				statement = &Statement{
					T:                 "ContinueStatement",
					ContinueStatement: continueStatement,
				}
			}
		} else {
			statement = &Statement{
				T:              "BreakStatement",
				BreakStatement: breakStatement,
			}
		}
	} else {
		statement = &Statement{
			T:              "WhileStatement",
			WhileStatement: whileStatement,
		}
	}

	return
}

func (parser *Parser) breakStatement() (breakStatement *BreakStatement, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeBreak {
		err = fmt.Errorf("[%w] expect TokenTypeBreak, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeSemicolon {
		err = fmt.Errorf("[%w] expect TokenTypeSemicolon, but %s", ExpectError, token.T)
		return
	}

	breakStatement = new(BreakStatement)

	return
}

func (parser *Parser) continueStatement() (continueStatement *ContinueStatement, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeContinue {
		err = fmt.Errorf("[%w] expect TokenTypeContinue, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeSemicolon {
		err = fmt.Errorf("[%w] expect TokenTypeSemicolon, but %s", ExpectError, token.T)
		return
	}

	continueStatement = new(ContinueStatement)

	return
}

func (parser *Parser) returnStatement() (returnStatement *ReturnStatement, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeReturn {
		err = fmt.Errorf("[%w] expect TokenTypeReturn, but %s", ExpectError, token.T)
		return
	}

	expression, err := parser.expression()
	if err != nil {
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeSemicolon {
		err = fmt.Errorf("[%w] expect TokenTypeSemicolon, but %s", ExpectError, token.T)
		return
	}

	returnStatement = &ReturnStatement{
		Expression: expression,
	}

	return
}

func (parser *Parser) whileStatement() (whileStatement *WhileStatement, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeWhile {
		err = fmt.Errorf("[%w] expect TokenTypeWhile, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] expect TokenTypeLp, but %s", ExpectError, token.T)
		return
	}

	expression, err := parser.expression()
	if err != nil {
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeRp {
		err = fmt.Errorf("[%w] expect TokenTypeRp, but %s", ExpectError, token.T)
		return
	}

	block, err := parser.block()
	if err != nil {
		return
	}

	whileStatement = new(WhileStatement)
	whileStatement.Expression = expression
	whileStatement.Block = block

	return
}

func (parser *Parser) ifStatement() (ifStatement *IfStatement, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeIf {
		err = fmt.Errorf("[%w] expect TokenTypeIf, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] expect TokenTypeLp, but %s", ExpectError, token.T)
		return
	}

	expression, err := parser.expression()
	if err != nil {
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeRp {
		err = fmt.Errorf("[%w] expect TokenTypeRp, but %s", ExpectError, token.T)
		return
	}

	block, err := parser.block()
	if err != nil {
		return
	}

	elseIfList, err := parser.elseIfList()
	if err != nil {
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeElse {
		err = fmt.Errorf("[%w] expect TokenTypeElse, but %s", ExpectError, token.T)
		return
	}

	elseBlock, err := parser.block()
	if err != nil {

	}

	ifStatement = &IfStatement{
		Expression: expression,
		Block:      block,
		ElseBlock:  elseBlock,
		ElseIfList: elseIfList,
	}

	return
}

func (parser *Parser) elseIfList() (elseIfList *ElseIfList, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	elseIf, err := parser.elseIf()
	if err != nil {
		return
	}

	elseIfList = new(ElseIfList)
	elseIfList.ElseIfList = append(elseIfList.ElseIfList, elseIf)

	for {
		elseIf, err = parser.elseIf()
		if err != nil {
			break
		}

		elseIfList.ElseIfList = append(elseIfList.ElseIfList, elseIf)
	}

	return
}

func (parser *Parser) elseIf() (elseIf *ElseIf, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeElseIf {
		err = fmt.Errorf("[%w] expect TokenTypeElseIf, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] expect TokenTypeLp, but %s", ExpectError, token.T)
		return
	}

	expression, err := parser.expression()
	if err != nil {
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeRp {
		err = fmt.Errorf("[%w] expect TokenTypeRp, but %s", ExpectError, token.T)
		return
	}

	block, err := parser.block()
	if err != nil {
		return
	}

	elseIf = &ElseIf{
		Expression: expression,
		Block:      block,
	}

	return
}

func (parser *Parser) expression() (expression *Expression, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	funcCallExpression, err := parser.funcCallExpression()
	if err != nil {
		var additiveExpression *AdditiveExpression
		additiveExpression, err = parser.additiveExpression()
		if err != nil {
			var identifierToken, token *lexer.Token
			identifierToken, err = parser.lexer.NextToken()
			if err != nil {
				err = fmt.Errorf("[lexer error] %s", err)
				return
			}
			if identifierToken.T != lexer.TokenTypeIdentifier {
				err = fmt.Errorf("[%w] expect TokenTypeIdentifier, but %s", ExpectError, token.T)
				return
			}
			token, err = parser.lexer.NextToken()
			if err != nil {
				err = fmt.Errorf("[lexer error] %s", err)
				return
			}
			if token.T != lexer.TokenTypeAssign {
				err = fmt.Errorf("[%w] expect TokenTypeAssign, but %s", ExpectError, token.T)
				return
			}

			var additiveExpression *AdditiveExpression
			additiveExpression, err = parser.additiveExpression()
			if err != nil {
				return
			}

			expression = &Expression{
				T:                  "Identifier",
				AdditiveExpression: additiveExpression,
				Identifier:         identifierToken.V,
			}
		} else {
			expression = &Expression{
				T:                  "AdditiveExpression",
				AdditiveExpression: additiveExpression,
			}
		}
	} else {
		expression = &Expression{
			T:                  "FuncCallExpression",
			FuncCallExpression: funcCallExpression,
		}
	}

	return
}

func (parser *Parser) additiveExpression() (additiveExpression *AdditiveExpression, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	multiplicativeExpression, err := parser.multiplicativeExpression()
	if err != nil {
		return
	}

	token, err := parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			additiveExpression = &AdditiveExpression{
				T:                        "",
				MultiplicativeExpression: multiplicativeExpression,
			}
			err = nil
		} else {
			err = fmt.Errorf("[lexer error] %s", err)
		}
		return
	}

	var t string
	if token.T != lexer.TokenTypeAdd {
		if token.T != lexer.TokenTypeSub {
			err = fmt.Errorf("[%w] expect TokenTypeAdd or TokenTypeSub, but %s", ExpectError, token.T)
			return
		} else {
			t = "Sub"
		}
	} else {
		t = "Add"
	}

	additiveExpressionChild, err := parser.additiveExpression()
	if err != nil {
		return
	}

	additiveExpression = &AdditiveExpression{
		T:                        t,
		MultiplicativeExpression: multiplicativeExpression,
		AdditiveExpression:       additiveExpressionChild,
	}

	return
}

func (parser *Parser) multiplicativeExpression() (multiplicativeExpression *MultiplicativeExpression, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	primaryExpression, err := parser.primaryExpression()
	if err != nil {
		return
	}

	token, err := parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			multiplicativeExpression = &MultiplicativeExpression{
				T:                 "",
				PrimaryExpression: primaryExpression,
			}
			err = nil
		} else {
			err = fmt.Errorf("[lexer error] %s", err)
		}
		return
	}

	var t string
	if token.T != lexer.TokenTypeMul {
		if token.T != lexer.TokenTypeDiv {
			err = fmt.Errorf("[%w] expect TokenTypeMul or TokenTypeDiv, but %s", ExpectError, token.T)
			return
		} else {
			t = "Div"
		}
	} else {
		t = "Mul"
	}

	multiplicativeExpressionChild, err := parser.multiplicativeExpression()
	if err != nil {
		return
	}

	multiplicativeExpression = &MultiplicativeExpression{
		T:                        t,
		PrimaryExpression:        primaryExpression,
		MultiplicativeExpression: multiplicativeExpressionChild,
	}

	return
}

func (parser *Parser) primaryExpression() (primaryExpression *PrimaryExpression, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if token.T != lexer.TokenTypeString {
		if token.T != lexer.TokenTypeNumber {
			if token.T != lexer.TokenTypeIdentifier {
				var funcCallExpression *FuncCallExpression
				funcCallExpression, err = parser.funcCallExpression()
				if err != nil {
					var lpToken *lexer.Token
					lpToken, err = parser.lexer.NextToken()
					if err != nil {
						err = fmt.Errorf("[lexer error] %s", err)
						return
					}
					if lpToken.T != lexer.TokenTypeLp {
						err = fmt.Errorf("[%w] expect TokenTypeRp, but %s", ExpectError, lpToken.T)
						return
					}

					var expression *Expression
					expression, err = parser.expression()
					if err != nil {
						if errors.Is(err, lexer.TokenEofErr) {
							err = fmt.Errorf("[%w] primaryExpression ecpect expression but eof, %s", EofError, err.Error())
						}
						return
					}

					var rpToken *lexer.Token
					rpToken, err = parser.lexer.NextToken()
					if err != nil {
						if errors.Is(err, lexer.TokenEofErr) {
							err = fmt.Errorf("[%w] primaryExpression ecpect TokenTypeRp but eof, %s", EofError, err.Error())
						} else {
							err = fmt.Errorf("[lexer error] %s", err)
						}
						return
					}
					if rpToken.T != lexer.TokenTypeRp {
						err = fmt.Errorf("[%w] expect TokenTypeRp, but %s", ExpectError, rpToken.T)
						return
					}

					primaryExpression = &PrimaryExpression{
						T:          "Expression",
						Expression: expression,
					}
				} else {
					primaryExpression = &PrimaryExpression{
						T:                  "FuncCallExpression",
						FuncCallExpression: funcCallExpression,
					}
				}
			} else {
				primaryExpression = &PrimaryExpression{
					T:          "Identifier",
					Identifier: token.V,
				}
			}
		} else {
			primaryExpression = &PrimaryExpression{
				T:      "Number",
				Number: token.V,
			}
		}
	} else {
		primaryExpression = &PrimaryExpression{
			T:      "String",
			String: token.V,
		}
	}
	return
}

func (parser *Parser) funcCallExpression() (funcCallExpression *FuncCallExpression, err error) {
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
	}()
	identifierToken, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %s", err)
		return
	}
	if identifierToken.T != lexer.TokenTypeIdentifier {
		err = fmt.Errorf("[%w] expect TokenTypeIdentifier, but %s", ExpectError, identifierToken.T)
		return
	}

	lpToken, err := parser.lexer.NextToken()
	if err != nil {
		if err == lexer.TokenEofErr {
			err = fmt.Errorf("[%w] funcCallExpression ecpect TokenTypeLp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %s", err)
		}
		return
	}
	if lpToken.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] expect TokenTypeLp, but %s", ExpectError, lpToken.T)
		return
	}

	rpToken, err := parser.lexer.NextToken()
	if err != nil {
		if err == lexer.TokenEofErr {
			err = fmt.Errorf("[%w] funcCallExpression ecpect TokenTypeRp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %s", err)
		}
		return
	}
	if rpToken.T != lexer.TokenTypeRp {
		var argumentList *ArgumentList
		argumentList, err = parser.argumentList()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = fmt.Errorf("[%w] funcCallExpression ecpect argumentList but eof, %s", EofError, err.Error())
			}
			return nil, err
		}

		rpToken, err = parser.lexer.NextToken()
		if err != nil {
			if err == lexer.TokenEofErr {
				err = fmt.Errorf("[%w] funcCallExpression ecpect TokenTypeRp but eof, %s", EofError, err.Error())
			} else {
				err = fmt.Errorf("[lexer error] %s", err)
			}
			return
		}
		if rpToken.T != lexer.TokenTypeRp {
			err = fmt.Errorf("[%w] expect TokenTypeRp, but %s", ExpectError, lpToken.T)
			return
		}

		funcCallExpression = &FuncCallExpression{
			Identifier:   identifierToken.V,
			ArgumentList: argumentList,
		}
	} else {
		funcCallExpression = &FuncCallExpression{
			Identifier: identifierToken.V,
		}
	}

	return
}
