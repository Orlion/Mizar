package ast

type ExpressionType int8

const (
	ExpressionTypeString ExpressionType = iota + 1
	ExpressionTypeInt
	ExpressionTypeDouble
	ExpressionTypeNull
	ExpressionTypeBool
	ExpressionTypeNewObject
	ExpressionTypeCall
)

type Expression struct {
	StringLiteral       string
	IntLiteral          int64
	DoubleLiteral       float64
	NullLiteral         *struct{}
	BoolLiteral         bool
	NewObjectExpression *NewObjectExpression
	CallExpression      *CallExpression
	Type                ExpressionType
}

func (expr *Expression) Accept(visitor Visitor) {

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
	ArgumentList []*Expression
}

func (newObjExpr *NewObjectExpression) accept(visitor Visitor) {

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

func (callExpr *CallExpression) Accept(visitor Visitor) {

}

type VarCallExpressionType int8

const (
	VarCallExpressionTypeThis VarCallExpressionType = iota + 1 // this
	VarCallExpressionTypeVar                                   // abc
	VarCallExpressionTypeCall                                  // this.a().b
)

type VarCallExpression struct {
	CallExpression *CallExpression
	This           string
	Var            string
	Type           VarCallExpressionType
}

func (varCallExpr *VarCallExpression) Accept(visitor Visitor) {

}

type MethodCallExpression struct {
	CallExpression *CallExpression
	Name           string
	ArgumentList   []*Expression
}

func (methodCallExpr *MethodCallExpression) Accept(visitor Visitor) {

}

type MethodCall struct {
	Name         string
	ArgumentList []*Expression
}

func (mc *MethodCall) Accept(visitor Visitor) {

}
