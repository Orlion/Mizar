package main

import (
	"fmt"
	"mizar/lexer"
)

func main() {
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
	for {
		token, err := lexer.NextToken()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(token.V)
	}
}
