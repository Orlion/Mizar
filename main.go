package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"mizar/lexer"
	"mizar/log"
	"mizar/parser"

	"github.com/sirupsen/logrus"
)

func init() {
	var logLevel uint
	flag.Parse()
	flag.UintVar(&logLevel, "log-level", uint(logrus.TraceLevel), "default log level")
	log.Init(logrus.Level(logLevel))
}

func main() {
	const source = `
	abc = 123;
	def = "你好啊";
	if (abc) {
		print(abc);
	}
	`

	//lexer1 := lexer.NewLexer(source)
	//for {
	//	t, err := lexer1.NextToken()
	//	fmt.Println(t)
	//	if err != nil {
	//		break
	//	}
	//}
	//
	//return

	lexer := lexer.NewLexer(source)
	parserObj := parser.NewParser(lexer)
	ast, err := parserObj.Parse()
	if err != nil {
		fmt.Println("parse error: %w", err)
		return
	}

	bytes, err := json.Marshal(ast)
	fmt.Println(err)
	fmt.Println(string(bytes))
}
