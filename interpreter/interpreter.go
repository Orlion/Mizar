package interpreter

import (
	"errors"
	"fmt"
	"mizar/parser"
)

type Interpreter struct {
	globalEnvironment *globalEnvironment
	localEnvironment  *localEnvironment
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		globalEnvironment: newGlobalEnvironment(),
	}
}

func (interpreter *Interpreter) Exec(ast *parser.TranslationUnit) error {
	for _, definitionOrStatement := range ast.DefinitionOrStatementList {
		if err := interpreter.definitionOrStatement(definitionOrStatement); err != nil {
			return err
		}
	}

	return nil
}

func (interpreter *Interpreter) definitionOrStatement(definitionOrStatement *parser.DefinitionOrStatement) error {
	if parser.DefinitionOrStatementTypeFunctionDefinition == definitionOrStatement.T {
		return interpreter.functionDefinition(definitionOrStatement.FunctionDefinition)
	} else {
		return interpreter.statement(definitionOrStatement.Statement)
	}
}

func (interpreter *Interpreter) functionDefinition(functionDefinition *parser.FunctionDefinition) error {
	interpreter.globalEnvironment.addFunc(functionDefinition.Name, functionDefinition)
	return nil
}

func (interpreter *Interpreter) statement(statement *parser.Statement) error {
	switch statement.T {
	case parser.StatementTypeExpression:

	}
	return nil
}

func (interpreter *Interpreter) expressionStatement(expression *parser.Expression) error {
	return interpreter.expression(expression)
}

func (interpreter *Interpreter) expression(expression *parser.Expression) (mval *mizarValue, err error) {
	switch expression.T {
	case parser.ExpressionTypeAdditiveExpression:
		return interpreter.additiveExpression(expression.AdditiveExpression)
	default:
		return nil
	}

}

func (interpreter *Interpreter) additiveExpression(additiveExpression *parser.AdditiveExpression) (mval *mizarValue, err error) {
	switch additiveExpression.T {
	case parser.AdditiveExpressionTypeNull:
		return interpreter.multiplicativeExpression(additiveExpression.MultiplicativeExpression)
	default:
		return nil
	}
}

func (interpreter *Interpreter) multiplicativeExpression(multiplicativeExpression *parser.MultiplicativeExpression) (mval *mizarValue, err error) {
	switch multiplicativeExpression.T {
	case parser.MultiplicativeExpressionTypeNull:

	}
}

func (interpreter *Interpreter) primaryExpression(primaryExpression *parser.PrimaryExpression) (mval *mizarValue, err error) {
	switch primaryExpression.T {
	case parser.PrimaryExpressionTypeString:
		mval = &mizarValue{
			v: &value{
				str: primaryExpression.String,
			},
			t: mizarValueTypeString,
		}
	case parser.PrimaryExpressionTypeNumber:
		mval = &mizarValue{
			v: &value{
				str: primaryExpression.Number,
			},
			t: mizarValueTypeNumber,
		}
	case parser.PrimaryExpressionTypeExpression:
		mval, err = interpreter.expression(primaryExpression.Expression)
	case parser.PrimaryExpressionTypeFuncCallExpression:
		mval, err = interpreter.funcCallExpression(primaryExpression.FuncCallExpression)
	case parser.PrimaryExpressionTypeIdentifier:
		// 从作用域中获取标识符对应的变量
		mval, err = interpreter.findVariable(primaryExpression.Identifier)
	}

	return
}

func (interpreter *Interpreter) funcCallExpression(funcCallExpression *parser.FuncCallExpression) (mval *mizarValue, err error) {
	// 从全局作用域中获取函数定义
	funcDefine := interpreter.globalEnvironment.getFunc(funcCallExpression.Identifier)
	if funcDefine != nil {
		err = errors.New(fmt.Sprintf("Undefined func: %s", funcCallExpression.Identifier))
		return
	}

	// 判断实参数与形参数是否一致
	if len(funcCallExpression.ArgumentList.ExpressionList) != len(funcDefine.ParameterList.ParameterList) {
		err = errors.New(fmt.Sprintf("call func: %s arguments number != parameters number", funcCallExpression.Identifier))
		return
	}

	// 重置局部作用域，开始进行执行
	interpreter.localEnvironment = nil

	// 创建新局部作用域
	localEnvironment := newLocalEnvironment()

	var argumentMVal *mizarValue
	// 将实参赋值给形参并注入到局部作用域中
	for k, parameter := range funcDefine.ParameterList.ParameterList {
		argumentMVal, err = interpreter.expression(funcCallExpression.ArgumentList.ExpressionList[k])
		if err != nil {
			return
		}

		localEnvironment.addVariable(parameter.Name, argumentMVal)
	}

	// 修改当前解释器的局部作用域
	interpreter.localEnvironment = localEnvironment

	// 执行
	return interpreter.block(funcDefine.Block)
}

func (interpreter *Interpreter) block(block *parser.Block) error {
	return nil
}

func (interpreter *Interpreter) findVariable(name string) (mval *mizarValue, err error) {
	if interpreter.localEnvironment != nil {
		// 先从局部作用域中尝试获取局部变量
		mval = interpreter.localEnvironment.getVariable(name)
		if nil == mval {
			// 再从全局作用域中尝试获取全局变量
			mval = interpreter.globalEnvironment.getVariable(name)
			if nil == mval {
				err = errors.New(fmt.Sprintf("Undefined variable: %s", name))
			}
		}
	} else {
		// 再从全局作用域中尝试获取全局变量
		mval = interpreter.globalEnvironment.getVariable(name)
		if nil == mval {
			err = errors.New(fmt.Sprintf("Undefined variable: %s", name))
		}
	}

	return
}
