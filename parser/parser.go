package parser

import (
	"errors"
	"fmt"
	"mizar/lexer"
	"mizar/log"

	"github.com/sirupsen/logrus"
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
	log.Trace(nil, "definitionOrStatement()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "definitionOrStatement() end")
	}()
	functionDefinition, err := parser.functionDefinition()
	if err == nil {
		definitionOrStatement = &DefinitionOrStatement{
			FunctionDefinition: functionDefinition,
			T:                  "FunctionDefinition",
		}
	} else {
		var statement *Statement
		statement, err = parser.statement()
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
	log.Trace(nil, "functionDefinition()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "functionDefinition() end")
	}()
	var token *lexer.Token
	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if token.T != lexer.TokenTypeFunc {
		err = fmt.Errorf("[%w] functionDefinition expect token:TokenTypeFunc but %s", ExpectError, token.T)
		return
	}

	nameToken, err := parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] functionDefinition ecpect TokenTypeIdentifier but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if nameToken.T != lexer.TokenTypeIdentifier {
		err = fmt.Errorf("[%w] functionDefinition expect TokenTypeIdentifier, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] functionDefinition ecpect TokenTypeIdentifier but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] functionDefinition expect TokenTypeLp, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] functionDefinition ecpect TokenTypeRp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}

	var parameterList *ParameterList
	if token.T != lexer.TokenTypeRp {
		parameterList, err = parser.parameterList()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = fmt.Errorf("[%w] functionDefinition ecpect parameterList but eof, %s", EofError, err.Error())
			}
			return
		}

		token, err = parser.lexer.NextToken()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = fmt.Errorf("[%w] functionDefinition ecpect TokenTypeRp but eof, %s", EofError, err.Error())
			} else {
				err = fmt.Errorf("[lexer error] %w", err)
			}
			return
		}
		if token.T != lexer.TokenTypeRp {
			err = fmt.Errorf("[%w] functionDefinition expect TokenTypeRp, but %s", ExpectError, token.T)
		}
	}

	block, err := parser.block()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] functionDefinition ecpect block but eof, %s", EofError, err.Error())
		}
		return
	}

	functionDefinition = &FunctionDefinition{
		ParameterList: parameterList,
		Block:         block,
		Name:          nameToken.V,
	}

	return
}

func (parser *Parser) parameterList() (parameterList *ParameterList, err error) {
	log.Trace(nil, "parameterList()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "parameterList() end")
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if token.T != lexer.TokenTypeIdentifier {
		err = fmt.Errorf("[%w] parameterList expect TokenTypeIdentifier, but %s", ExpectError, token.T)
		return
	}

	parameterList = new(ParameterList)
	parameterList.ParameterList = append(parameterList.ParameterList, &Parameter{token.V})

	for {
		token, err = parser.lexer.NextToken()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = nil
			} else {
				err = fmt.Errorf("[lexer error] %w", err)
			}
			break
		}
		if token.T != lexer.TokenTypeComma {
			err = fmt.Errorf("[%w] parameterList expect TokenTypeComma, but %s", ExpectError, token.T)
			break
		}

		token, err = parser.lexer.NextToken()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = fmt.Errorf("[%w] parameterList ecpect TokenTypeIdentifier but eof, %s", EofError, err.Error())
			} else {
				err = fmt.Errorf("[lexer error] %w", err)
			}
			break
		}
		if token.T != lexer.TokenTypeIdentifier {
			err = fmt.Errorf("[%w] parameterList expect TokenTypeIdentifier, but %s", ExpectError, token.T)
			break
		}

		parameterList.ParameterList = append(parameterList.ParameterList, &Parameter{token.V})
	}

	return
}

func (parser *Parser) argumentList() (argumentList *ArgumentList, err error) {
	log.Trace(nil, "argumentList()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "argumentList() end")
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
			if errors.Is(err, lexer.TokenEofErr) {
				err = nil
			} else {
				err = fmt.Errorf("[lexer error] %w", err)
			}
			break
		}
		if token.T != lexer.TokenTypeComma {
			err = fmt.Errorf("[%w] argumentList expect TokenTypeComma, but %s", ExpectError, token.T)
			break
		}

		expression, err = parser.expression()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = fmt.Errorf("[%w] argumentList ecpect expression but eof, %s", EofError, err.Error())
			}
			break
		}

		argumentList.ExpressionList = append(argumentList.ExpressionList, expression)
	}

	return
}

func (parser *Parser) block() (block *Block, err error) {
	log.Trace(nil, "block()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "block() end")
	}()
	var token *lexer.Token
	token, err = parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if token.T != lexer.TokenTypeLc {
		err = fmt.Errorf("[%w] block expect TokenTypeLc, but %s", ExpectError, token.T)
		return
	}

	var statementList *StatementList
	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] block ecpect TokenTypeRc but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeRc {
		statementList, err = parser.statementList()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = fmt.Errorf("[%w] block ecpect statementList but eof, %s", EofError, err.Error())
			}
			return
		}

		token, err = parser.lexer.NextToken()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = fmt.Errorf("[%w] block ecpect TokenTypeRc but eof, %s", EofError, err.Error())
			} else {
				err = fmt.Errorf("[lexer error] %w", err)
			}
			return
		}
		if token.T != lexer.TokenTypeRc {
			err = fmt.Errorf("[%w] block expect TokenTypeRc, but %s", ExpectError, token.T)
			return
		}
	}

	block = new(Block)
	block.StatementList = statementList

	return
}

func (parser *Parser) statementList() (statementList *StatementList, err error) {
	log.Trace(nil, "statementList()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "statementList() end")
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
			if errors.Is(err, lexer.TokenEofErr) {
				err = nil
			}
			break
		}

		statementList.StatementList = append(statementList.StatementList, staement)
	}

	return
}

func (parser *Parser) statement() (statement *Statement, err error) {
	log.Trace(nil, "statement()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "statement() end")
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
					var expression *Expression
					expression, err = parser.expression()
					if err != nil {
						return
					}

					var token *lexer.Token
					token, err = parser.lexer.NextToken()
					if err != nil {
						if errors.Is(err, lexer.TokenEofErr) {
							err = fmt.Errorf("[%w] statement ecpect TokenTypeSemicolon but eof, %s", EofError, err.Error())
						} else {
							err = fmt.Errorf("[lexer error] %w", err)
						}
						return
					}
					if token.T != lexer.TokenTypeSemicolon {
						err = fmt.Errorf("[%w] statement expect TokenTypeSemicolon, but %s", ExpectError, token.T)
						return
					}

					statement = &Statement{
						T:          "Expression",
						Expression: expression,
					}
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
	log.Trace(nil, "breakStatement()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "breakStatement() end")
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if token.T != lexer.TokenTypeBreak {
		err = fmt.Errorf("[%w] breakStatement expect TokenTypeBreak, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] breakStatement ecpect TokenTypeSemicolon but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeSemicolon {
		err = fmt.Errorf("[%w] breakStatement expect TokenTypeSemicolon, but %s", ExpectError, token.T)
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
		log.Trace(logrus.Fields{
			"err": err,
		}, "continueStatement() end")
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if token.T != lexer.TokenTypeContinue {
		err = fmt.Errorf("[%w] continueStatement expect TokenTypeContinue, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] continueStatement ecpect TokenTypeSemicolon but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeSemicolon {
		err = fmt.Errorf("[%w] continueStatement expect TokenTypeSemicolon, but %s", ExpectError, token.T)
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
		log.Trace(logrus.Fields{
			"err": err,
		}, "returnStatement() end")
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if token.T != lexer.TokenTypeReturn {
		err = fmt.Errorf("[%w] returnStatement expect TokenTypeReturn, but %s", ExpectError, token.T)
		return
	}

	expression, err := parser.expression()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] returnStatement ecpect expression but eof, %s", EofError, err.Error())
		}
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] returnStatement ecpect TokenTypeSemicolon but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeSemicolon {
		err = fmt.Errorf("[%w] returnStatement expect TokenTypeSemicolon, but %s", ExpectError, token.T)
		return
	}

	returnStatement = &ReturnStatement{
		Expression: expression,
	}

	return
}

func (parser *Parser) whileStatement() (whileStatement *WhileStatement, err error) {
	log.Trace(nil, "whileStatement()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "whileStatement() end")
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if token.T != lexer.TokenTypeWhile {
		err = fmt.Errorf("[%w] whileStatement expect TokenTypeWhile, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] whileStatement ecpect TokenTypeLp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] whileStatement expect TokenTypeLp, but %s", ExpectError, token.T)
		return
	}

	expression, err := parser.expression()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] whileStatement ecpect expression but eof, %s", EofError, err.Error())
		}
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] whileStatement ecpect TokenTypeRp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeRp {
		err = fmt.Errorf("[%w] whileStatement expect TokenTypeRp, but %s", ExpectError, token.T)
		return
	}

	block, err := parser.block()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] whileStatement ecpect block but eof, %s", EofError, err.Error())
		}
		return
	}

	whileStatement = new(WhileStatement)
	whileStatement.Expression = expression
	whileStatement.Block = block

	return
}

func (parser *Parser) ifStatement() (ifStatement *IfStatement, err error) {
	log.Trace(nil, "ifStatement()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "ifStatement() end")
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if token.T != lexer.TokenTypeIf {
		err = fmt.Errorf("[%w] ifStatement expect TokenTypeIf, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] ifStatement ecpect TokenTypeLp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] ifStatement expect TokenTypeLp, but %s", ExpectError, token.T)
		return
	}

	expression, err := parser.expression()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] ifStatement ecpect expression but eof, %s", EofError, err.Error())
		}
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] ifStatement ecpect TokenTypeRp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeRp {
		err = fmt.Errorf("[%w] ifStatement expect TokenTypeRp, but %s", ExpectError, token.T)
		return
	}

	block, err := parser.block()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] ifStatement ecpect block but eof, %s", EofError, err.Error())
		}
		return
	}

	elseIfList, err := parser.elseIfList()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			ifStatement = &IfStatement{
				Expression: expression,
				Block:      block,
			}
			err = nil
		}
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			ifStatement = &IfStatement{
				Expression: expression,
				Block:      block,
			}
			err = nil
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeElse {
		ifStatement = &IfStatement{
			Expression: expression,
			Block:      block,
		}
		err = nil
		return
	}

	elseBlock, err := parser.block()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] ifStatement elseIf ecpect expression but eof, %s", EofError, err.Error())
		}
		return
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
	log.Trace(nil, "elseIfList()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "elseIfList() end")
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
			if errors.Is(err, lexer.TokenEofErr) {
				err = nil
			}
			break
		}

		elseIfList.ElseIfList = append(elseIfList.ElseIfList, elseIf)
	}

	return
}

func (parser *Parser) elseIf() (elseIf *ElseIf, err error) {
	log.Trace(nil, "elseIf()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "elseIf() end")
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if token.T != lexer.TokenTypeElseIf {
		err = fmt.Errorf("[%w] elseIf expect TokenTypeElseIf, but %s", ExpectError, token.T)
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] elseIf ecpect TokenTypeLp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] elseIf expect TokenTypeLp, but %s", ExpectError, token.T)
		return
	}

	expression, err := parser.expression()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] elseIf ecpect expression but eof, %s", EofError, err.Error())
		}
		return
	}

	token, err = parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] elseIf ecpect TokenTypeRp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if token.T != lexer.TokenTypeRp {
		err = fmt.Errorf("[%w] elseIf expect TokenTypeRp, but %s", ExpectError, token.T)
		return
	}

	block, err := parser.block()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] elseIf ecpect block but eof, %s", EofError, err.Error())
		}
		return
	}

	elseIf = &ElseIf{
		Expression: expression,
		Block:      block,
	}

	return
}

func (parser *Parser) expression() (expression *Expression, err error) {
	log.Trace(nil, "expression()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "expression() end")
	}()

	assignment, err := parser.assignment()
	if err != nil {
		var funcCallExpression *FuncCallExpression
		funcCallExpression, err = parser.funcCallExpression()
		if err != nil {
			var additiveExpression *AdditiveExpression
			additiveExpression, err = parser.additiveExpression()
			if err == nil {
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
	} else {
		expressionChild, err := parser.expression()
		if err != nil {
			if errors.Is(err, lexer.TokenEofErr) {
				err = fmt.Errorf("[%w] expression ecpect expression but eof, %s", EofError, err.Error())
			}
		} else {
			expression = &Expression{
				T:          "Assignment",
				Assignment: assignment,
				Expression: expressionChild,
			}
		}
	}

	return
}

func (parser *Parser) assignment() (assignment *Assignment, err error) {
	log.Trace(nil, "assignment()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "assignment() end")
	}()

	identifierToken, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if identifierToken.T != lexer.TokenTypeIdentifier {
		err = fmt.Errorf("[%w] expression expect TokenTypeIdentifier, but %s", ExpectError, identifierToken.T)
		return
	}

	assignToken, err := parser.lexer.NextToken()
	if err != nil {
		if errors.Is(err, lexer.TokenEofErr) {
			err = fmt.Errorf("[%w] expression ecpect TokenTypeAssign but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if assignToken.T != lexer.TokenTypeAssign {
		err = fmt.Errorf("[%w] expression expect TokenTypeAssign, but %s", ExpectError, assignToken.T)
		return
	}

	assignment = &Assignment{
		Identifier: identifierToken.V,
	}

	return
}

func (parser *Parser) additiveExpression() (additiveExpression *AdditiveExpression, err error) {
	log.Trace(nil, "additiveExpression()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "additiveExpression() end")
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
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}

	var t string
	if token.T != lexer.TokenTypeAdd {
		if token.T != lexer.TokenTypeSub {
			parser.lexer.Return(token)
			additiveExpression = &AdditiveExpression{
				T:                        "",
				MultiplicativeExpression: multiplicativeExpression,
			}
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
	log.Trace(nil, "multiplicativeExpression()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "multiplicativeExpression() end")
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
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}

	var t string
	if token.T != lexer.TokenTypeMul {
		if token.T != lexer.TokenTypeDiv {
			parser.lexer.Return(token)
			multiplicativeExpression = &MultiplicativeExpression{
				T:                 "",
				PrimaryExpression: primaryExpression,
			}
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
	log.Trace(nil, "primaryExpression()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "primaryExpression() end")
	}()
	token, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
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
						err = fmt.Errorf("[lexer error] %w", err)
						return
					}
					if lpToken.T != lexer.TokenTypeLp {
						err = fmt.Errorf("[%w] primaryExpression expect TokenTypeRp, but %s", ExpectError, lpToken.T)
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
							err = fmt.Errorf("[lexer error] %w", err)
						}
						return
					}
					if rpToken.T != lexer.TokenTypeRp {
						err = fmt.Errorf("[%w] primaryExpression expect TokenTypeRp, but %s", ExpectError, rpToken.T)
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
	log.Trace(nil, "funcCallExpression()")
	parser.lexer.Begin()
	defer func() {
		if err != nil {
			parser.lexer.Rollback()
		} else {
			parser.lexer.Commit()
		}
		log.Trace(logrus.Fields{
			"err": err,
		}, "funcCallExpression() end")
	}()
	identifierToken, err := parser.lexer.NextToken()
	if err != nil {
		err = fmt.Errorf("[lexer error] %w", err)
		return
	}
	if identifierToken.T != lexer.TokenTypeIdentifier {
		err = fmt.Errorf("[%w] funcCallExpression expect TokenTypeIdentifier, but %s", ExpectError, identifierToken.T)
		return
	}

	lpToken, err := parser.lexer.NextToken()
	if err != nil {
		if err == lexer.TokenEofErr {
			err = fmt.Errorf("[%w] funcCallExpression ecpect TokenTypeLp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
		}
		return
	}
	if lpToken.T != lexer.TokenTypeLp {
		err = fmt.Errorf("[%w] funcCallExpression expect TokenTypeLp, but %s", ExpectError, lpToken.T)
		return
	}

	rpToken, err := parser.lexer.NextToken()
	if err != nil {
		if err == lexer.TokenEofErr {
			err = fmt.Errorf("[%w] funcCallExpression ecpect TokenTypeRp but eof, %s", EofError, err.Error())
		} else {
			err = fmt.Errorf("[lexer error] %w", err)
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
				err = fmt.Errorf("[lexer error] %w", err)
			}
			return
		}
		if rpToken.T != lexer.TokenTypeRp {
			err = fmt.Errorf("[%w] funcCallExpression expect TokenTypeRp, but %s", ExpectError, lpToken.T)
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
