package ast

type Ast struct {
	InterfaceDefineList []*InterfaceDefine
	ClassDefineList     []*ClassDefine
}

type InterfaceDefine struct {
	Name       string
	MethodList []*InterfaceMethod
}

type ClassDefine struct {
	MethodDefinitionList         []*MethodDefinition
	AbstractMethodDefinitionList []*MethodDefinition // 抽象方法列表
	PropertyDefinitionList       []*PropertyDefinition
}

// 接口中的方法
type InterfaceMethod struct {
	Type          string
	ParameterList []*Parameter
}
