package ast

type ASTVistor interface {
	visit()
}

type Node interface {
	accept(ASTVistor)
}
