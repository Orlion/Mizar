package ast

type ArgumentList struct {
	List []*Expression
}

func (argsList *ArgumentList) accept(vistor ASTVistor) {

}
