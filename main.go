package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"mizar/interpreter"
	"mizar/lexer"
	"mizar/log"
	"mizar/parser"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	var (
		logLevel uint
		dumpAst  bool
	)

	if len(os.Args) < 2 {
		fmt.Println("请输入文件名")
		return
	}

	flag.UintVar(&logLevel, "log-level", uint(logrus.ErrorLevel), "日志级别")
	flag.BoolVar(&dumpAst, "dumpast", false, "是否打印抽象语法树")
	flag.Parse()

	log.Init(logrus.Level(logLevel))

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	source := string(b)

	lexer := lexer.NewLexer(source)
	parserObj := parser.NewParser(lexer)
	ast, err := parserObj.Parse()
	if err != nil {
		fmt.Println("parse error: %w", err)
		return
	}

	if dumpAst {
		bytes, _ := json.Marshal(ast)
		fmt.Println(string(bytes))
	}

	interpreter := interpreter.New()
	err = interpreter.Exec(ast)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
