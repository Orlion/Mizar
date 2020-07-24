package interpreter

import (
	"errors"
	"fmt"
	"mizar/parser"
)

type Interpreter struct {
	globalEnvironment *globalEnvironment
	localEnvironment  *localEnvironment
	nativeFuncMap     map[string]NativeFunc
}

func New() *Interpreter {
	interpreter := &Interpreter{
		globalEnvironment: newGlobalEnvironment(),
	}

	interpreter.initNativeFunc()

	return interpreter
}

func (interpreter *Interpreter) initNativeFunc() {
	interpreter.nativeFuncMap = make(map[string]NativeFunc)
	interpreter.nativeFuncMap["print"] = func(args []*mizarValue) (mval *mizarValue, err error) {
		if len(args) != 1 {
			err = errors.New("print 函数只接收一个参数")
			return
		}

		if args[0].t == mizarValueTypeString {
			fmt.Print(args[0].v.str)
		} else {
			fmt.Print(args[0].v.number)
		}
		return
	}

	interpreter.nativeFuncMap["println"] = func(args []*mizarValue) (mval *mizarValue, err error) {
		if len(args) != 1 {
			err = errors.New("print 函数只接收一个参数")
			return
		}

		if args[0].t == mizarValueTypeString {
			fmt.Println(args[0].v.str)
		} else {
			fmt.Println(args[0].v.number)
		}
		return
	}

	return
}

func (interpreter *Interpreter) Exec(ast *parser.TranslationUnit) error {
	for _, definitionOrStatement := range ast.DefinitionOrStatementList {
		if err := interpreter.definitionOrStatement(definitionOrStatement); err != nil {
			return err
		}
	}

	return nil
}

func (interpreter *Interpreter) definitionOrStatement(definitionOrStatement *parser.DefinitionOrStatement) (err error) {
	if parser.DefinitionOrStatementTypeFunctionDefinition == definitionOrStatement.T {
		err = interpreter.functionDefinition(definitionOrStatement.FunctionDefinition)
	} else {
		_, err = interpreter.statement(definitionOrStatement.Statement)

	}

	return
}

func (interpreter *Interpreter) functionDefinition(functionDefinition *parser.FunctionDefinition) (err error) {
	if _, exists := interpreter.nativeFuncMap[functionDefinition.Name]; exists {
		err = errors.New(fmt.Sprintf("函数: %s 已由原生定义", functionDefinition.Name))
		return
	}

	interpreter.globalEnvironment.addFunc(functionDefinition.Name, functionDefinition)
	return nil
}

func (interpreter *Interpreter) statement(statement *parser.Statement) (result *StatementResult, err error) {
	switch statement.T {
	case parser.StatementTypeExpression:
		return interpreter.expressionStatement(statement.Expression)
	case parser.StatementTypeBreakStatement:
		return interpreter.breakStatement(statement.BreakStatement)
	case parser.StatementTypeContinueStatement:
		return interpreter.continueStatement(statement.ContinueStatement)
	case parser.StatementTypeReturnStatement:
		return interpreter.returnStatement(statement.ReturnStatement)
	case parser.StatementTypeIfStatement:
		return interpreter.ifStatement(statement.IfStatement)
	case parser.StatementTypeWhileStatement:
		return interpreter.whileStatement(statement.WhileStatement)
	}

	return
}

func (interpreter *Interpreter) ifStatement(ifStatement *parser.IfStatement) (result *StatementResult, err error) {
	conditionBool, err := interpreter.conditionBool(ifStatement.Expression)
	if err != nil {
		return
	}

	if conditionBool {
		// 走ifblock
		result, err = interpreter.block(ifStatement.Block)
	} else {
		if ifStatement.ElseIfList != nil {
			// 按顺序依次遍历elseIfList
			for _, elseIf := range ifStatement.ElseIfList.ElseIfList {
				conditionBool, err = interpreter.conditionBool(elseIf.Expression)
				if conditionBool {
					result, err = interpreter.block(elseIf.Block)
					break
				}
			}
		}

		// 如果elseIfList无一命中则走else
		if ifStatement.ElseBlock != nil {
			result, err = interpreter.block(ifStatement.ElseBlock)
		}
	}
	return
}

func (interpreter *Interpreter) whileStatement(whileStatement *parser.WhileStatement) (result *StatementResult, err error) {
	var conditionBool bool
	for {
		conditionBool, err = interpreter.conditionBool(whileStatement.Expression)
		if err != nil {
			return
		}

		if !conditionBool {
			break
		}

		result, err = interpreter.block(whileStatement.Block)
		if result.T == StatementResultTypeBreak {
			result.T = StatementResultTypeNormal
			break
		} else if result.T == StatementResultTypeReturn {
			break
		}
	}

	return
}

func (interpreter *Interpreter) breakStatement(breakStatement *parser.BreakStatement) (result *StatementResult, err error) {
	result = new(StatementResult)
	result.T = StatementResultTypeBreak
	return
}

func (interpreter *Interpreter) continueStatement(breakStatement *parser.ContinueStatement) (result *StatementResult, err error) {
	result = new(StatementResult)
	result.T = StatementResultTypeContinue
	return
}

func (interpreter *Interpreter) returnStatement(returnStatement *parser.ReturnStatement) (result *StatementResult, err error) {
	mval, err := interpreter.expression(returnStatement.Expression)
	if err != nil {
		return
	}

	result = new(StatementResult)
	result.T = StatementResultTypeReturn
	result.ReturnValue = mval

	return
}

// 判断if/while条件是true/false
func (interpreter *Interpreter) conditionBool(expression *parser.Expression) (b bool, err error) {
	conditionMval, err := interpreter.expression(expression)
	if err != nil {
		return
	}

	if conditionMval.t == mizarValueTypeNumber {
		b = conditionMval.v.number != 0
	} else {
		b = conditionMval.v.str != ""
	}

	return
}

func (interpreter *Interpreter) expressionStatement(expression *parser.Expression) (result *StatementResult, err error) {
	_, err = interpreter.expression(expression)
	if err != nil {
		return
	}

	result = new(StatementResult)
	result.T = StatementResultTypeNormal

	return
}

func (interpreter *Interpreter) expression(expression *parser.Expression) (mval *mizarValue, err error) {
	switch expression.T {
	case parser.ExpressionTypeAdditiveExpression:
		return interpreter.additiveExpression(expression.AdditiveExpression)
	case parser.ExpressionTypeAssignment:
		mval, err = interpreter.expression(expression.Expression)
		if err != nil {
			return
		}

		interpreter.setVariable(expression.Assignment.Identifier, mval)
		return
	default:
		err = errors.New("invalid expression.T")
		return
	}
}

func (interpreter *Interpreter) additiveExpression(additiveExpression *parser.AdditiveExpression) (mval *mizarValue, err error) {
	switch additiveExpression.T {
	case parser.AdditiveExpressionTypeNull:
		return interpreter.multiplicativeExpression(additiveExpression.MultiplicativeExpression)
	case parser.AdditiveExpressionTypeSub:
		// 减法
		var multiplicativeExpressionMVal, additiveExpressionMVal *mizarValue
		multiplicativeExpressionMVal, err = interpreter.multiplicativeExpression(additiveExpression.MultiplicativeExpression)
		if err != nil {
			return
		}

		additiveExpressionMVal, err = interpreter.additiveExpression(additiveExpression.AdditiveExpression)
		if err != nil {
			return
		}

		if multiplicativeExpressionMVal.t != mizarValueTypeNumber || additiveExpressionMVal.t != mizarValueTypeNumber {
			err = errors.New("非数字不能进行减法操作")
			return
		}

		mval = new(mizarValue)
		mval.t = mizarValueTypeNumber
		mval.v = &value{
			number: multiplicativeExpressionMVal.v.number - additiveExpressionMVal.v.number,
		}

		return
	case parser.AdditiveExpressionTypeAdd:
		// 加法
		var multiplicativeExpressionMVal, additiveExpressionMVal *mizarValue
		multiplicativeExpressionMVal, err = interpreter.multiplicativeExpression(additiveExpression.MultiplicativeExpression)
		if err != nil {
			return
		}

		additiveExpressionMVal, err = interpreter.additiveExpression(additiveExpression.AdditiveExpression)
		if err != nil {
			return
		}

		mval = new(mizarValue)

		// 如果有一方是字符串则为字符串拼接
		if multiplicativeExpressionMVal.t == mizarValueTypeString {
			if additiveExpressionMVal.t == mizarValueTypeString {
				// 两个字符串拼接
				mval.t = mizarValueTypeString
				mval.v = &value{
					str: multiplicativeExpressionMVal.v.str + additiveExpressionMVal.v.str,
				}
			} else {
				// 字符串与数字进行加法，字符串直接当0计算
				mval.t = mizarValueTypeNumber
				mval.v = &value{
					number: additiveExpressionMVal.v.number,
				}
			}
		} else {
			if additiveExpressionMVal.t == mizarValueTypeString {
				// 数字与字符串进行加法，字符串直接当0计算
				mval.t = mizarValueTypeNumber
				mval.v = &value{
					number: multiplicativeExpressionMVal.v.number,
				}
			} else {
				// 数字与数字进行加法
				mval.t = mizarValueTypeNumber
				mval.v = &value{
					number: multiplicativeExpressionMVal.v.number + additiveExpressionMVal.v.number,
				}
			}
		}

		return
	default:
		return
	}
}

func (interpreter *Interpreter) multiplicativeExpression(multiplicativeExpression *parser.MultiplicativeExpression) (mval *mizarValue, err error) {
	switch multiplicativeExpression.T {
	case parser.MultiplicativeExpressionTypeNull:
		return interpreter.primaryExpression(multiplicativeExpression.PrimaryExpression)
	case parser.MultiplicativeExpressionTypeMul:
		// 乘法
		var multiplicativeExpressionMVal, primaryExpressionMVal *mizarValue
		multiplicativeExpressionMVal, err = interpreter.multiplicativeExpression(multiplicativeExpression.MultiplicativeExpression)
		if err != nil {
			return
		}

		primaryExpressionMVal, err = interpreter.primaryExpression(multiplicativeExpression.PrimaryExpression)
		if err != nil {
			return
		}

		if multiplicativeExpressionMVal.t != mizarValueTypeNumber || primaryExpressionMVal.t != mizarValueTypeNumber {
			err = errors.New("非数字不能进行乘法操作")
			return
		}

		mval = new(mizarValue)
		mval.t = mizarValueTypeNumber
		mval.v = &value{
			number: primaryExpressionMVal.v.number * multiplicativeExpressionMVal.v.number,
		}

		return
	case parser.MultiplicativeExpressionTypeDiv:
		// 除法
		var multiplicativeExpressionMVal, primaryExpressionMVal *mizarValue
		multiplicativeExpressionMVal, err = interpreter.multiplicativeExpression(multiplicativeExpression.MultiplicativeExpression)
		if err != nil {
			return
		}

		primaryExpressionMVal, err = interpreter.primaryExpression(multiplicativeExpression.PrimaryExpression)
		if err != nil {
			return
		}

		if multiplicativeExpressionMVal.t != mizarValueTypeNumber || primaryExpressionMVal.t != mizarValueTypeNumber {
			err = errors.New("非数字不能进行除法操作")
			return
		}

		mval = new(mizarValue)
		mval.t = mizarValueTypeNumber
		mval.v = &value{
			number: primaryExpressionMVal.v.number / multiplicativeExpressionMVal.v.number,
		}

		return
	default:
		err = errors.New("invalid multiplicativeExpression.T")
		return
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
				number: primaryExpression.Number,
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
	if funcDefine == nil {
		return interpreter.callNativeFunc(funcCallExpression)
	}

	// 判断实参数与形参数是否一致
	var argsCount, paramsCount = 0, 0
	if funcCallExpression.ArgumentList != nil {
		argsCount = len(funcCallExpression.ArgumentList.ExpressionList)
	}
	if funcDefine.ParameterList != nil {
		paramsCount = len(funcDefine.ParameterList.ParameterList)
	}

	if argsCount != paramsCount {
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

		localEnvironment.setVariable(parameter.Name, argumentMVal)
	}

	// 修改当前解释器的局部作用域
	interpreter.localEnvironment = localEnvironment

	// 执行
	var result *StatementResult
	result, err = interpreter.block(funcDefine.Block)
	if err != nil {
		return
	}

	if result.T != StatementResultTypeReturn {
		err = errors.New(fmt.Sprintf("func: %s missing return", funcCallExpression.Identifier))
		return
	}

	mval = result.ReturnValue
	return
}

func (interpreter *Interpreter) callNativeFunc(funcCallExpression *parser.FuncCallExpression) (mval *mizarValue, err error) {
	if nativeFunc, exists := interpreter.nativeFuncMap[funcCallExpression.Identifier]; !exists {
		err = errors.New(fmt.Sprintf("函数: %s 未定义", funcCallExpression.Identifier))
	} else {
		var (
			argsMVal *mizarValue
			argsList []*mizarValue
		)
		for _, argumentExpression := range funcCallExpression.ArgumentList.ExpressionList {
			argsMVal, err = interpreter.expression(argumentExpression)
			if err != nil {
				return
			}
			argsList = append(argsList, argsMVal)
		}
		mval, err = nativeFunc(argsList)
	}

	return
}

func (interpreter *Interpreter) block(block *parser.Block) (result *StatementResult, err error) {
	if block != nil && block.StatementList != nil {
		for _, statement := range block.StatementList.StatementList {
			result, err = interpreter.statement(statement)
			if err != nil {
				break
			}

			if result.T != StatementResultTypeNormal {
				break
			}
		}
	}

	return
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

func (interpreter *Interpreter) setVariable(name string, mval *mizarValue) {
	// 判断局部作用域是否存在
	if interpreter.localEnvironment != nil {
		// 如果局部变量存在则修改
		if interpreter.localEnvironment.existsVariable(name) {
			interpreter.localEnvironment.setVariable(name, mval)
		} else {
			// 如果全局变量存在则修改全局变量
			if interpreter.globalEnvironment.existsVariable(name) {
				interpreter.globalEnvironment.setVariable(name, mval)
			} else {
				interpreter.localEnvironment.setVariable(name, mval)
			}
		}
	} else {
		// 直接set全局变量
		interpreter.globalEnvironment.setVariable(name, mval)
	}

	return
}
