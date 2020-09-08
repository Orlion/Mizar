package ast

type Argument struct {
	Expression Expression
}

type ArgumentList struct {
	list []*Argument
}
