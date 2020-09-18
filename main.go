package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"mizar/lexer"
	"mizar/log"
	"mizar/parser"

	"github.com/sirupsen/logrus"
)

func main() {
	var (
		logLevel uint
		dumpAst  bool
	)

	// if len(os.Args) < 2 {
	// 	fmt.Println("请输入文件名")
	// 	return
	// }

	flag.UintVar(&logLevel, "log-level", uint(logrus.TraceLevel), "日志级别")
	flag.BoolVar(&dumpAst, "dumpast", false, "是否打印抽象语法树")
	flag.Parse()

	log.Init(logrus.Level(logLevel))

	// b, err := ioutil.ReadFile(os.Args[1])
	b, err := ioutil.ReadFile("demo/base.mi")
	if err != nil {
		fmt.Println(err)
		return
	}

	source := string(b)

	lexer := lexer.NewLexer(source)
	parserObj := parser.NewParser()
	ast, err := parserObj.Parse(lexer)
	if err != nil {
		fmt.Println(err)
	} else {
		bytes, _ := json.Marshal(ast)
		fmt.Println(string(bytes))
	}

	return
}
