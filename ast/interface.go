package ast

type Interface struct {
	Name       string
	MethodList []*InterfaceMethod
}

func (i *Interface) accept(vistor ASTVistor) {

}

type InterfaceMethodList struct {
	List []*InterfaceMethod
}

func (iml *InterfaceMethodList) accept(vistor ASTVistor) {

}

// 接口中的方法
type InterfaceMethod struct {
	Type          string
	Name          string
	ParameterList []*Parameter
}

func (im *InterfaceMethod) accept(vistor ASTVistor) {

}
