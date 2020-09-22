package asm

import (
	"errors"
	"fmt"
	"mizar/ast"
)

var (
	ClassNotDefineErr  = errors.New("Class not define, class name:")
	MethodNotDefineErr = errors.New("Method not define, method name: ")
)

func Generate(tu *ast.TranslationUnit) string {
	visitor := &Generator{}
	tu.Accept(visitor)
	return ""
}

type Generator struct {
}

func (g *Generator) Visit(node ast.Node) {

}

func (g *Generator) visitTranslationUnit(tu *ast.TranslationUnit) (err error) {
	mainClass, exists := tu.ClassMap["Main"]
	if !exists {
		err = fmt.Errorf("%w Main", ClassNotDefineErr)
		return
	}

	mainClass.Accept(g)

	return
}

func (g *Generator) visitClass(class *ast.Class) (err error) {
	if class.Name == "Main" {
		if method, exists := class.MethodDefinitionMap["main"][""]; exists {
			method.Accept(g)
		} else {
			err = fmt.Errorf("%w main", MethodNotDefineErr)
		}
	} else {

	}

	return
}

func (g *Generator) visitMethod(method *ast.MethodDefinition) (err error) {
	if method.Name == "main" {
		method.Block.Accept(g)
	} else {

	}

	return
}

func (g *Generator) visitBlock(block *ast.Block) (err error) {
	for _, stmt := range block.StatementList {
		stmt.Accept(g)
	}

	return
}

func (g *Generator) visitStatement(stmt *ast.Statement) (err error) {
	switch stmt.Type {
	case ast.StatementTypeBreak:
		stmt.BreakStatement.Accept(g)
	}

	return
}

func (g *Generator) visitBreakStatement(stmt *ast.Statement) (err error) {
	

	return
}
