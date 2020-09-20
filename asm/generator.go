package asm

import "mizar/ast"

func Generate(tu *ast.TranslationUnit) string {
	visitor := &Generator{}
	tu.Accept(visitor)
	return ""
}

type Generator struct {
}

func (g *Generator) Visit(node ast.Node) {

}

func (g *Generator) visitTranslationUnit(tu *ast.TranslationUnit) {
	for _, class := range tu.ClassList {
		g.visitClass(class)
	}
}

func (g *Generator) visitClass(class *ast.Class) {

}
