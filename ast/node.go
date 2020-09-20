package ast

type Visitor interface {
	Visit(Node)
}

type Node interface {
	Accept(Visitor)
}
