package ast

type Class struct {
	Name                        string
	IsAbstract                  bool
	MethodDefinitionMap         map[string]map[string]*MethodDefinition
	AbstractMethodDefinitionMap map[string]map[string]*MethodDefinition // 抽象方法列表
	PropertyDefinitionMap       map[string]*PropertyDefinition
	Extends                     []string
	Implements                  []string
}

func (c *Class) Accept(visitor Visitor) {

}

type Extends struct {
	ClassNameList []string
}

func (ex *Extends) accept(visitor Visitor) {

}

type Implements struct {
	InterfaceNameList []string
}

func (id *Implements) accept(visitor Visitor) {

}

type ClassStatementList struct {
	MethodDefinitionMap         map[string]map[string]*MethodDefinition // map[MethodName]map[MethodParameterList]*Method
	AbstractMethodDefinitionMap map[string]map[string]*MethodDefinition // 抽象方法列表
	PropertyDefinitionMap       map[string]*PropertyDefinition          // 类属性声明
}

func (csl *ClassStatementList) accept(visitor Visitor) {

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

func (cs *ClassStatement) accept(visitor Visitor) {

}

// 类属性
type PropertyDefinition struct {
	ModifierType MemberModifierType // 修饰符
	Type         string
	Name         string
	Expr         *Expression
}

func (pd *PropertyDefinition) accept(visitor Visitor) {

}

// 方法定义
type MethodDefinition struct {
	ModifierType  MemberModifierType // 修饰符
	Type          string
	Name          string // 方法名
	ParameterList []*Parameter
	Block         *Block
}

func (md *MethodDefinition) Accept(visitor Visitor) {

}
