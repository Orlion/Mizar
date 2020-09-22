package ast

type Interface struct {
	Name      string
	MethodMap map[string]map[string]*InterfaceMethod
}

func (i *Interface) Accept(visitor Visitor) {

}

type InterfaceMethodList struct {
	List []*InterfaceMethod
}

func (iml *InterfaceMethodList) Accept(visitor Visitor) {

}

// 接口中的方法
type InterfaceMethod struct {
	Type          string
	Name          string
	ParameterList []*Parameter
}

func (im *InterfaceMethod) Accept(visitor Visitor) {

}
