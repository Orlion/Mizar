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
	MethodList         []*ClassMethod
	AbstractMethodList []*ClassMethod // 抽象方法列表
	PropertyList       []*Property
}

// 接口中的方法
type InterfaceMethod struct {
	Type          string
	ParameterList []*Parameter
}

type MemberModifier string

const (
	ModiferPublic    MemberModifier = "public"
	ModiferProtected MemberModifier = "protected"
	ModiferPrivate   MemberModifier = "private"
	ModiferAbstract  MemberModifier = "abstract"
)

// 类属性
type Property struct {
	Modifier MemberModifier // 修饰符
	Type     string
	Name     string
	Expr     Expression
}

// 类中的方法
type ClassMethod struct {
	Modifier      MemberModifier // 修饰符
	Type          string
	Name          string // 方法名
	ParameterList *ParameterList
	Block         *Block
}

type Block struct {
	StatementList []Statement
}
