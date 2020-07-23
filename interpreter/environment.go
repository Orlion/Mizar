package interpreter

import "mizar/parser"

type localEnvironment struct {
	variableMap map[string]*mizarValue
}

func newLocalEnvironment() *localEnvironment {
	return &localEnvironment{
		variableMap: make(map[string]*mizarValue),
	}
}

func (localEnvironment *localEnvironment) addVariable(name string, mval *mizarValue) {
	localEnvironment.variableMap[name] = mval
}

func (localEnvironment *localEnvironment) getVariable(name string) *mizarValue {
	if mval, exists := localEnvironment.variableMap[name]; exists {
		return mval
	} else {
		return nil
	}
}

type globalEnvironment struct {
	funcMap     map[string]*parser.FunctionDefinition
	variableMap map[string]*mizarValue
}

func newGlobalEnvironment() *globalEnvironment {
	return &globalEnvironment{make(map[string]*parser.FunctionDefinition), make(map[string]*mizarValue)}
}

func (globalEnvironment *globalEnvironment) addVariable(name string, mval *mizarValue) {
	globalEnvironment.variableMap[name] = mval
}

func (globalEnvironment *globalEnvironment) getVariable(name string) *mizarValue {
	if mval, exists := globalEnvironment.variableMap[name]; exists {
		return mval
	} else {
		return nil
	}
}

func (globalEnvironment *globalEnvironment) addFunc(name string, functionDefinition *parser.FunctionDefinition) {
	globalEnvironment.funcMap[name] = functionDefinition
}

func (globalEnvironment *globalEnvironment) getFunc(name string) *parser.FunctionDefinition {
	if functionDefinition, exists := globalEnvironment.funcMap[name]; exists {
		return functionDefinition
	} else {
		return nil
	}
}
