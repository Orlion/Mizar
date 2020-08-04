package parser

import "fmt"

type Symbol string

const (
	SymbolStringLiteral                           = "STRING_LITERAL"
	SymbolIntLiteral                              = "INT_LITERAL"
	SymbolDoubleLiteral                           = "DOUBLE_LITERAL"
	SymbolNull                                    = "NULL"
	SymbolTrue                                    = "TRUE"
	SymbolFalse                                   = "FALSE"
	SymbolIdentifier                              = "IDENTIFIER"
	SymbolNew                                     = "New"
	SymbolLp                                      = "LP"
	SymbolRp                                      = "RP"
	SymbolDot                                     = "DOT"
	SymbolDoubleColon                             = "DOUBLE_COLON"
	SymbolLc                                      = "LC"
	SymbolRc                                      = "RC"
	SymbolComma                                   = "COMMA"
	SymbolReturn                                  = "RETURN"
	SymbolContinue                                = "CONTINUE"
	SymbolBreak                                   = "BREAK"
	SymbolSemicolon                               = "SEMICOLON"
	SymbolFor                                     = "FOR"
	SymbolIf                                      = "IF"
	SymbolElse                                    = "ELSE"
	SymbolWhile                                   = "WHILE"
	SymbolAssign                                  = "ASSIGN"
	SymbolClass                                   = "CLASS"
	SymbolInterface                               = "INTERFACE"
	SymbolAbstract                                = "ASBTRACT"
	SymbolImplements                              = "IMPLEMENTS"
	SymbolExtends                                 = "EXTENDS"
	SymbolVoid                                    = "VOID"
	SymbolPublic                                  = "PUBLIC"
	SymbolPrivate                                 = "PRIVATE"
	SymbolProtected                               = "PROTECTED"
	SymbolConst                                   = "CONST"
	SymbolExpression                              = "expression"
	SymbolBlock                                   = "block"
	SymbolArgumentList                            = "argument_list"
	SymbolMethodCallExpression                    = "method_call_expression"
	SymbolNewObjExpression                        = "new_obj_expression"
	SymbolExpressionOpt                           = "expression_opt"
	SymbolReturnStatement                         = "return_statement"
	SymbolContinueStatement                       = "continue_statement"
	SymbolBreakStatement                          = "break_statement"
	SymbolForStatement                            = "for_statement"
	SymbolIfStatement                             = "if_statement"
	SymbolWhileStatement                          = "while_statement"
	SymbolValueDeclarationStatement               = "value_declaration_statement"
	SymbolStatement                               = "statement"
	SymbolStatementList                           = "statement_list"
	SymbolValueDeclaration                        = "value_declaration"
	SymbolClassStatement                          = "class_statement"
	SymbolClassStatementList                      = "class_statement_list"
	SymbolClassDeclaration                        = "class_declaration"
	SymbolImplementsDeclaration                   = "implements_declaration"
	SymbolExtendsDelcaration                      = "extends_declaration"
	SymbolInterfaceMethonDeclarationStatement     = "interface_method_declaration_statement"
	SymbolInterfaceMethodDeclarationStatementList = "interface_method_declaration_statement_list"
	SymbolInterfaceDeclaration                    = "interface_declaration"
	SymbolParameterList                           = "parameter_list"
	SymbolMethodDeclaration                       = "method_declaration"
	SymbolClassDeclarationList                    = "class_declaration_list"
	SymbolTranslationUnit                         = "translation_unit"
)

type Production struct {
	str    string   // 转成字符串表示
	left   Symbol   // 左侧非终结符
	right  []Symbol // 右侧符号列表
	dotPos int      // .的位置
}

func newProduction(left Symbol, right []Symbol, dotPos int) (p *Production) {
	p = &Production{
		left:   left,
		right:  right,
		dotPos: dotPos,
	}

	p.str = fmt.Sprintf("%s->", left)
	for k, v := range p.right {
		if p.dotPos == k {
			p.str += "."
		}
		p.str += string(v)
	}

	return
}

// .前移
func (p *Production) dotForward() *Production {
	return newProduction(p.left, p.right, p.dotPos+1)
}

func (p *Production) getDotSymbol() Symbol {
	return p.right[p.dotPos]
}

func (p *Production) print() {
	fmt.Println(p.str)
}

type ProductionManager struct {
	productionMap map[Symbol][]*Production
}

func newProductionManager() (pm *ProductionManager) {
	pm = new(ProductionManager)

	pm.productionMap[SymbolExpression] = []*Production{
		newProduction(SymbolExpression, []Symbol{SymbolStringLiteral}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolIntLiteral}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolDoubleLiteral}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolNull}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolTrue}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolFalse}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolIdentifier}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolIdentifier, SymbolDot, SymbolIdentifier}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolIdentifier, SymbolDoubleColon, SymbolIdentifier}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolMethodCallExpression}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolNewObjExpression}, 0),
	}

	pm.productionMap[SymbolMethodCallExpression] = []*Production{
		newProduction(SymbolMethodCallExpression, []Symbol{SymbolNew, SymbolIdentifier, SymbolLp, SymbolArgumentList, SymbolRp}, 0),
		newProduction(SymbolMethodCallExpression, []Symbol{SymbolNew, SymbolIdentifier, SymbolLp, SymbolRp}, 0),
	}

	return
}

// 获取以左侧为left的生成式列表
func (pm *ProductionManager) getProductions(left Symbol) (productions []*Production) {
	return
}
