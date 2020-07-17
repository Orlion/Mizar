package main

import (
	"fmt"
	"mizar/lexer"
	"mizar/parser"
)

func a() {
	const source = `
	abc = "我"
	def = 2
	ghi = abc + def
	if (ghi > 0) {
		jkl = "jkl"
	} else if {
		mn = "mn"
	} else {
		op = 123
	}
	print(ghi)
	`

	lexer := lexer.NewLexer(source)
	parserObj := parser.NewParser(lexer)
	ast, err := parserObj.Parse()
	if err != nil {
		fmt.Println("parse error: %w", err)
		return
	}

	fmt.Println(ast)
}
