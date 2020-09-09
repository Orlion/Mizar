package ast

type TypeVar struct {
	Type string
	Name string
}

func (typeVar *TypeVar) accept(ASTVistor) {

}
