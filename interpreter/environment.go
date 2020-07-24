package interpreter

import (
	"errors"
	"fmt"
	"mizar/parser"
)

type localEnvironment struct {
	variableMap map[string]*mizarValue
}

func newLocalEnvironment() *localEnvironment {
	return &localEnvironment{
		variableMap: make(map[string]*mizarValue),
	}
}

func (localEnvironment *localEnvironment) setVariable(name string, mval *mizarValue) {
	localEnvironment.variableMap[name] = mval
}

func (localEnvironment *localEnvironment) getVariable(name string) *mizarValue {
	if mval, exists := localEnvironment.variableMap[name]; exists {
		return mval
	} else {
		return nil
	}
}

func (localEnvironment *localEnvironment) existsVariable(name string) (exists bool) {
	_, exists = localEnvironment.variableMap[name]
	return
}

type globalEnvironment struct {
	funcMap     map[string]*parser.FunctionDefinition
	variableMap map[string]*mizarValue
}

func newGlobalEnvironment() *globalEnvironment {
	return &globalEnvironment{make(map[string]*parser.FunctionDefinition), make(map[string]*mizarValue)}
}

func (globalEnvironment *globalEnvironment) setVariable(name string, mval *mizarValue) {
	globalEnvironment.variableMap[name] = mval
}

func (globalEnvironment *globalEnvironment) getVariable(name string) *mizarValue {
	if mval, exists := globalEnvironment.variableMap[name]; exists {
		return mval
	} else {
		return nil
	}
}

func (globalEnvironment *globalEnvironment) existsVariable(name string) (exists bool) {
	_, exists = globalEnvironment.variableMap[name]
	return
}

func (globalEnvironment *globalEnvironment) addFunc(name string, functionDefinition *parser.FunctionDefinition) error {
	// 需要判断函数是否已定义
	if _, exists := globalEnvironment.funcMap[name]; exists {
		return errors.New(fmt.Sprintf("func: %s 已经定义过了，请勿重复定义", name))
	}

	globalEnvironment.funcMap[name] = functionDefinition

	return nil
}

func (globalEnvironment *globalEnvironment) getFunc(name string) *parser.FunctionDefinition {
	if functionDefinition, exists := globalEnvironment.funcMap[name]; exists {
		return functionDefinition
	} else {
		return nil
	}
}
