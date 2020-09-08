package ast

type Statement interface {
}

type ExpressionStatement struct {
	Expression Expression
}

type VarDeclarationStatement struct {
	Type string
	Name string
}

type VarAssignStatement struct {
	Type       string
	Name       string // this.a.b().c
	Expression Expression
}

type WhileStatement struct {
	Expression *Block
}

type IfStatement struct {
	CondExpression Expression
	IfBlock        *Block
	ElseBlock      *Block
}

type ForStatement struct {
	InitExpression Expression
	CondExpression Expression
	PostExpression Expression
	Block          *Block
}

type BreakStatement struct {
	Expression Expression
}

type ContinueStatement struct {
	Expression Expression
}

type ReturnStatement struct {
	Expression Expression
}
