package ast

// 形参
type Parameter struct {
	Type string
	Name string
}

type ParameterList struct {
	List []*Parameter
}
