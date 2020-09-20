package ast

type ArgumentList struct {
	List []*Expression
}

func (argsList *ArgumentList) Accept(visitor Visitor) {

}
