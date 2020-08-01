package parser

import "fmt"

type Symbol string

const (
	SymbolStringLiteral = "STRING_LITERAL"
	SymbolIntLiteral = "INT_LITERAL"
	SymbolDoubleLiteral = "DOUBLE_LITERAL"
	SymbolNullL = "NULL"
	SymbolTrue = "TRUE"
	SymbolFalse = "FALSE"
	SymbolIdentifier = "IDENTIFIER"
	SymbolNew = "New"
	SymbolLp = "LP"
	SymbolRp = "RP"
	SymbolDot = "DOT"
	SymbolDoubleColon = "DOUBLE_COLON"
	SymbolLc = "LC"
	SymbolRc = "RC"
	SymbolComma = "COMMA"
	SymbolReturn = "RETURN"
	SymbolContinue = "CONTINUE"
	SymbolBreak = "BREAK"
	SymbolSemicolon = "SEMICOLON"
	SymbolFor = "FOR"
	SymbolIf = "IF"
	SymbolElse = "ELSE"
	SymbolWhile = "WHILE"
	SymbolAssign = "ASSIGN"
	SymbolClass = "CLASS"
	SymbolInterface = "INTERFACE"
	SymbolAbstract = "ASBTRACT"
	SymbolImplements = "IMPLEMENTS"
	SymbolExtends = "EXTENDS"
	SymbolVoid = "VOID"
	SymbolPublic = "PUBLIC"
	SymbolPrivate = "PRIVATE"
	SymbolProtected = "PROTECTED"
	SymbolConst = "CONST"
	SymbolExpression = "expression"
	SymbolBlock = "block"
	SymbolArgumentList = "argument_list"
	SymbolMethodCallExpression = "method_call_expression"
)

type Production struct {
	left   string   // 左侧非终结符
	right  []string // 右侧符号列表
	dotPos int      // .的位置
}

type ProductionManager struct {
	productionMap map[int][]*Production
}

func newProduction(left string, right []string, dotPos int) *Production {
	return &Production{
		left:   left,
		right:  right,
		dotPos: dotPos,
	}
}

// .前移
func (p *Production) dotForward() *Production {
	return newProduction(p.left, p.right, p.dotPos+1)
}

func (p *Production) getDotSymbol() string {
	return p.right[p.dotPos]
}

func (p *Production) print() {
	fmt.Printf("%s -> ", p.left)
	for k, v := range p.right {
		if p.dotPos == k {
			fmt.Print(". ")
		}
		fmt.Print("" + v)
	}
	fmt.Println()
}

func newProductionManager() *ProductionManager {
	pm := new(ProductionManager)

	pm.productionMap[1] = newProduction(1, nil, 0)

	return pm
}
