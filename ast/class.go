package ast

type ImplementsDeclaration struct {
	InterfaceNameList []string
}

func (id *ImplementsDeclaration) accept(visitor ASTVistor) {

}

type ClassStatementList struct {
	List []*ClassStatement
}

func (csl *ClassStatementList) accept(visitor ASTVistor) {

}

type ClassStatementType int8

const (
	ClassStatementTypeProperty ClassStatementType = iota + 1
	ClassStatementTypeMethod
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
