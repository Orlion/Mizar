package ast

type Class struct {
	Name                         string
	IsAbstract                   bool
	MethodDefinitionList         []*MethodDefinition
	AbstractMethodDefinitionList []*MethodDefinition // 抽象方法列表
	PropertyDefinitionList       []*PropertyDefinition
	Extends                      []string
	Implements                   []string
}

func (c *Class) accept(visitor ASTVistor) {

}

type Extends struct {
	ClassNameList []string
}

func (ex *Extends) accept(visitor ASTVistor) {

}

type Implements struct {
	InterfaceNameList []string
}

func (id *Implements) accept(visitor ASTVistor) {

}

type ClassStatementList struct {
	MethodDefinitionList         []*MethodDefinition
	AbstractMethodDefinitionList []*MethodDefinition // 抽象方法列表
	PropertyDefinitionList       []*PropertyDefinition
}

func (csl *ClassStatementList) accept(visitor ASTVistor) {

}

type ClassStatementType int8

const (
	ClassStatementTypeProperty ClassStatementType = iota + 1
	ClassStatementTypeMethod
	ClassStatementTypeAbstractMethod
)

type ClassStatement struct {
	PropertyDefinition *PropertyDefinition
	MethodDefinition   *MethodDefinition
	Type               ClassStatementType
}

func (cs *ClassStatement) accept(visitor ASTVistor) {

}

// 类属性
type PropertyDefinition struct {
	ModifierType MemberModifierType // 修饰符
	Type         string
	Name         string
	Expr         *Expression
}

func (pd *PropertyDefinition) accept(visitor ASTVistor) {

}

// 方法定义
type MethodDefinition struct {
	ModifierType  MemberModifierType // 修饰符
	Type          string
	Name          string // 方法名
	ParameterList []*Parameter
	Block         *Block
}

func (md *MethodDefinition) accept(visitor ASTVistor) {

}
