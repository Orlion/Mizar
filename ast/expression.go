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
	StringLiteral       *StringLiteral
	IntLiteral          *IntLiteral
	DoubleLiteral       *DoubleLiteral
	NullLiteral         *NullLiteral
	BoolLiteral         *BoolLiteral
	NewObjectExpression *NewObjectExpression
	CallExpression      *CallExpression
	Type                ExpressionType
}

func (expr *Expression) accept(vistor ASTVistor) {

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

func (newObjExpr *NewObjectExpression) accept(ASTVistor) {

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

func (callExpr *CallExpression) accept(vistor ASTVistor) {

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

func (varCallExpr *VarCallExpression) accept(vistor ASTVistor) {

}

type MethodCallExpression struct {
	CallExpression *CallExpression
	Name           string
	ArgumentList   []*Expression
}

func (methodCallExpr *MethodCallExpression) accept(vistor ASTVistor) {

}

type MethodCall struct {
	Name         string
	ArgumentList []*Expression
}

func (mc *MethodCall) accept(vistor ASTVistor) {

}