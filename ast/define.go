package ast

type TranslationUnit struct {
	InterfaceList []*Interface
	ClassList     []*Class
}

func (tu *TranslationUnit) Accept(visitor Visitor) {
	visitor.Visit(tu)
}

type ClassInterfaceType int8

const (
	ClassInterfaceTypeClass ClassInterfaceType = iota + 1
	ClassInterfaceTypeInterface
)

type ClassInterface struct {
	Class     *Class
	Interface *Interface
	Type      ClassInterfaceType
}

func (ci *ClassInterface) Accept(visitor Visitor) {

}
