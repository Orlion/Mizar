package ast

type Expression interface {
	Node
}

type StringLiteral struct {
	Value string
}

type IntLiteral struct {
	Value int64
}

type DoubleLiteral struct {
	Value float64
}

type NullLiteral struct {
}

type BoolLiteral struct {
	Value bool
}

type NewObjectExpression struct {
	Name         string
	ArgumentList []*Argument
}

type CallExpressionType int8

const (
	CallExpressionTypeValCall CallExpressionType = iota + 1
	CallExpressionTypeMethodCall
)

type CallExpression struct {
	VarCallExpression    *VarCallExpression
	MethodCallExpression *MethodCallExpression
	Type                 CallExpressionType
}

type VarCallExpressionType int8

const (
	VarCallExpressionTypeThis VarCallExpressionType = iota + 1 // this
	VarCallExpressionTypeVar                                   // abc
	VarCallExpressionTypeCall                                  // this.a().b
)

type VarCallExpression struct {
	CallExpression        *CallExpression
	This                  string
	Var                   string
	VarCallExpressionType VarCallExpressionType
}

type MethodCallExpression struct {
	CallExpression *CallExpression
	Name           string
	ArgumentList   []*Argument
}

type Argument struct {
	Expression Expression
}
