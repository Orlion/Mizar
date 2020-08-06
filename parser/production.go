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
	SymbolNew                                     = "NEW"
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
	SymbolThis                                    = "THIS"
	SymbolInterface                               = "INTERFACE"
	SymbolAbstract                                = "ASBTRACT"
	SymbolImplements                              = "IMPLEMENTS"
	SymbolExtends                                 = "EXTENDS"
	SymbolVoid                                    = "VOID"
	SymbolPublic                                  = "PUBLIC"
	SymbolPrivate                                 = "PRIVATE"
	SymbolProtected                               = "PROTECTED"
	SymbolExpression                              = "expression"
	SymbolArgumentList                            = "argument_list"
	SymbolMethodCallExpression                    = "method_call_expression"
	SymbolVarCallExpression                       = "var_call_expression"
	SymbolCallExpression                          = "call_expression"
	SymbolNewObjExpression                        = "new_obj_expression"
	SymbolVarDeclaration                          = "var_declaration"
	SymbolVarAssignStatement                      = "var_assign_statement"
	SymbolVarDeclarationStatement                 = "var_declaration_statement"
	SymbolReturnStatement                         = "return_statement"
	SymbolContinueStatement                       = "continue_statement"
	SymbolBreakStatement                          = "break_statement"
	SymbolForStatement                            = "for_statement"
	SymbolIfStatement                             = "if_statement"
	SymbolWhileStatement                          = "while_statement"
	SymbolStatement                               = "statement"
	SymbolStatementList                           = "statement_list"
	SymbolBlock                                   = "block"
	SymbolMethodModifier                          = "method_modifier"
	SymbolParameter                               = "parameter"
	SymbolParameterList                           = "parameter_list"
	SymbolReturnValType                           = "return_val_type"
	SymbolMethodDefinition                        = "method_definition"
	SymbolVarModifier                             = "var_modifier"
	SymbolClassStatement                          = "class_statement"
	SymbolClassStatementList                      = "class_statement_list"
	SymbolClassDeclaration                        = "class_declaration"
	SymbolImplementsDeclaration                   = "implements_declaration"
	SymbolExtendsDelcaration                      = "extends_declaration"
	SymbolInterfaceMethodDeclarationStatement     = "interface_method_declaration_statement"
	SymbolInterfaceMethodDeclarationStatementList = "interface_method_declaration_statement_list"
	SymbolInterfaceDeclaration                    = "interface_declaration"
	SymbolClassInterfaceDeclaration               = "class_interface_declaration"
	SymbolClassInterfaceDeclarationList           = "class_interface_declaration_list"
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

	p.str = fmt.Sprintf("%s ->   ", left)
	for k, v := range p.right {
		if p.dotPos == k {
			p.str += ".   "
		}
		p.str += (string(v) + "   ")
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

	pm.productionMap = make(map[Symbol][]*Production)

	pm.productionMap[SymbolArgumentList] = []*Production{
		newProduction(SymbolArgumentList, []Symbol{SymbolExpression}, 0),
		newProduction(SymbolArgumentList, []Symbol{SymbolArgumentList, SymbolComma, SymbolExpression}, 0),
	}

	pm.productionMap[SymbolNewObjExpression] = []*Production{
		newProduction(SymbolNewObjExpression, []Symbol{SymbolNew, SymbolIdentifier, SymbolLp, SymbolRp}, 0),
		newProduction(SymbolNewObjExpression, []Symbol{SymbolNew, SymbolIdentifier, SymbolLp, SymbolArgumentList, SymbolRp}, 0),
	}

	pm.productionMap[SymbolVarCallExpression] = []*Production{
		newProduction(SymbolVarCallExpression, []Symbol{SymbolIdentifier}, 0),
		newProduction(SymbolVarCallExpression, []Symbol{SymbolThis}, 0),
		newProduction(SymbolVarCallExpression, []Symbol{SymbolCallExpression, SymbolDot, SymbolIdentifier}, 0),
	}

	pm.productionMap[SymbolMethodCallExpression] = []*Production{
		newProduction(SymbolMethodCallExpression, []Symbol{SymbolCallExpression, SymbolDot, SymbolIdentifier, SymbolLp, SymbolRp}, 0),
		newProduction(SymbolMethodCallExpression, []Symbol{SymbolCallExpression, SymbolDot, SymbolIdentifier, SymbolArgumentList, SymbolLp, SymbolRp}, 0),
	}

	pm.productionMap[SymbolCallExpression] = []*Production{
		newProduction(SymbolCallExpression, []Symbol{SymbolMethodCallExpression}, 0),
		newProduction(SymbolCallExpression, []Symbol{SymbolVarCallExpression}, 0),
	}

	pm.productionMap[SymbolExpression] = []*Production{
		newProduction(SymbolExpression, []Symbol{SymbolStringLiteral}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolIntLiteral}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolDoubleLiteral}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolNull}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolTrue}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolFalse}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolIdentifier}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolNewObjExpression}, 0),
		newProduction(SymbolExpression, []Symbol{SymbolCallExpression}, 0),
	}

	pm.productionMap[SymbolVarDeclaration] = []*Production{
		newProduction(SymbolVarDeclaration, []Symbol{SymbolIdentifier, SymbolIdentifier}, 0),
	}

	pm.productionMap[SymbolVarAssignStatement] = []*Production{
		newProduction(SymbolVarAssignStatement, []Symbol{SymbolVarDeclaration, SymbolAssign, SymbolExpression, SymbolSemicolon}, 0),
		newProduction(SymbolVarAssignStatement, []Symbol{SymbolVarCallExpression, SymbolAssign, SymbolExpression, SymbolSemicolon}, 0),
	}

	pm.productionMap[SymbolVarDeclarationStatement] = []*Production{
		newProduction(SymbolVarDeclarationStatement, []Symbol{SymbolVarDeclaration, SymbolSemicolon}, 0),
	}

	pm.productionMap[SymbolReturnStatement] = []*Production{
		newProduction(SymbolReturnStatement, []Symbol{SymbolReturn, SymbolSemicolon}, 0),
		newProduction(SymbolReturnStatement, []Symbol{SymbolReturn, SymbolExpression, SymbolSemicolon}, 0),
	}

	pm.productionMap[SymbolContinueStatement] = []*Production{
		newProduction(SymbolContinueStatement, []Symbol{SymbolContinue, SymbolSemicolon}, 0),
		newProduction(SymbolContinueStatement, []Symbol{SymbolContinue, SymbolExpression, SymbolSemicolon}, 0),
	}

	pm.productionMap[SymbolBreakStatement] = []*Production{
		newProduction(SymbolBreakStatement, []Symbol{SymbolBreak, SymbolSemicolon}, 0),
		newProduction(SymbolBreakStatement, []Symbol{SymbolBreak, SymbolExpression, SymbolSemicolon}, 0),
	}

	pm.productionMap[SymbolForStatement] = []*Production{
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolRp, SymbolSemicolon, SymbolSemicolon, SymbolLp, SymbolBlock}, 0),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolRp, SymbolExpression, SymbolSemicolon, SymbolSemicolon, SymbolLp, SymbolBlock}, 0),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolRp, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolLp, SymbolBlock}, 0),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolRp, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolLp, SymbolBlock}, 0),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolRp, SymbolExpression, SymbolSemicolon, SymbolSemicolon, SymbolExpression, SymbolLp, SymbolBlock}, 0),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolRp, SymbolSemicolon, SymbolSemicolon, SymbolExpression, SymbolLp, SymbolBlock}, 0),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolRp, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolLp, SymbolBlock}, 0),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolRp, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolLp, SymbolBlock}, 0),
	}

	pm.productionMap[SymbolIfStatement] = []*Production{
		newProduction(SymbolIfStatement, []Symbol{SymbolIf, SymbolLp, SymbolExpression, SymbolRp, SymbolBlock}, 0),
		newProduction(SymbolIfStatement, []Symbol{SymbolIf, SymbolLp, SymbolExpression, SymbolRp, SymbolBlock, SymbolElse, SymbolBlock}, 0),
	}

	pm.productionMap[SymbolWhileStatement] = []*Production{
		newProduction(SymbolWhileStatement, []Symbol{SymbolWhile, SymbolLp, SymbolExpression, SymbolRp, SymbolBlock}, 0),
	}

	pm.productionMap[SymbolStatement] = []*Production{
		newProduction(SymbolStatement, []Symbol{SymbolExpression, SymbolSemicolon}, 0),
		newProduction(SymbolStatement, []Symbol{SymbolVarDeclarationStatement}, 0),
		newProduction(SymbolStatement, []Symbol{SymbolVarAssignStatement}, 0),
		newProduction(SymbolStatement, []Symbol{SymbolWhileStatement}, 0),
		newProduction(SymbolStatement, []Symbol{SymbolIfStatement}, 0),
		newProduction(SymbolStatement, []Symbol{SymbolForStatement}, 0),
		newProduction(SymbolStatement, []Symbol{SymbolBreakStatement}, 0),
		newProduction(SymbolStatement, []Symbol{SymbolContinueStatement}, 0),
		newProduction(SymbolStatement, []Symbol{SymbolReturnStatement}, 0),
	}

	pm.productionMap[SymbolStatementList] = []*Production{
		newProduction(SymbolStatementList, []Symbol{SymbolStatement}, 0),
		newProduction(SymbolStatementList, []Symbol{SymbolStatementList, SymbolStatement}, 0),
	}

	pm.productionMap[SymbolBlock] = []*Production{
		newProduction(SymbolBlock, []Symbol{SymbolLc, SymbolRc}, 0),
		newProduction(SymbolBlock, []Symbol{SymbolLc, SymbolStatementList, SymbolRc}, 0),
	}

	pm.productionMap[SymbolMethodModifier] = []*Production{
		newProduction(SymbolMethodModifier, []Symbol{SymbolPublic}, 0),
		newProduction(SymbolMethodModifier, []Symbol{SymbolProtected}, 0),
		newProduction(SymbolMethodModifier, []Symbol{SymbolPrivate}, 0),
		newProduction(SymbolMethodModifier, []Symbol{SymbolAbstract}, 0),
	}

	pm.productionMap[SymbolParameter] = []*Production{
		newProduction(SymbolParameter, []Symbol{SymbolIdentifier, SymbolIdentifier}, 0),
	}

	pm.productionMap[SymbolParameterList] = []*Production{
		newProduction(SymbolParameterList, []Symbol{SymbolParameter}, 0),
		newProduction(SymbolParameterList, []Symbol{SymbolParameterList, SymbolComma, SymbolParameter}, 0),
	}

	pm.productionMap[SymbolReturnValType] = []*Production{
		newProduction(SymbolReturnValType, []Symbol{SymbolVoid}, 0),
		newProduction(SymbolReturnValType, []Symbol{SymbolIdentifier}, 0),
	}

	pm.productionMap[SymbolMethodDefinition] = []*Production{
		newProduction(SymbolMethodDefinition, []Symbol{SymbolMethodModifier, SymbolReturnValType, SymbolIdentifier, SymbolRp, SymbolLp, SymbolBlock}, 0),
		newProduction(SymbolMethodDefinition, []Symbol{SymbolMethodModifier, SymbolReturnValType, SymbolIdentifier, SymbolRp, SymbolParameterList, SymbolLp, SymbolBlock}, 0),
	}

	pm.productionMap[SymbolVarModifier] = []*Production{
		newProduction(SymbolVarModifier, []Symbol{SymbolPublic}, 0),
		newProduction(SymbolVarModifier, []Symbol{SymbolProtected}, 0),
		newProduction(SymbolVarModifier, []Symbol{SymbolPrivate}, 0),
	}

	pm.productionMap[SymbolClassStatement] = []*Production{
		newProduction(SymbolClassStatement, []Symbol{SymbolVarModifier, SymbolVarDeclaration, SymbolSemicolon}, 0),
		newProduction(SymbolClassStatement, []Symbol{SymbolVarModifier, SymbolVarDeclaration, SymbolAssign, SymbolExpression, SymbolSemicolon}, 0),
	}

	pm.productionMap[SymbolClassStatementList] = []*Production{
		newProduction(SymbolClassStatementList, []Symbol{SymbolClassStatement}, 0),
		newProduction(SymbolClassStatementList, []Symbol{SymbolClassStatementList, SymbolClassStatement}, 0),
	}

	pm.productionMap[SymbolImplementsDeclaration] = []*Production{
		newProduction(SymbolImplementsDeclaration, []Symbol{SymbolImplements, SymbolIdentifier}, 0),
		newProduction(SymbolImplementsDeclaration, []Symbol{SymbolImplementsDeclaration, SymbolComma, SymbolIdentifier}, 0),
	}

	pm.productionMap[SymbolExtendsDelcaration] = []*Production{
		newProduction(SymbolExtendsDelcaration, []Symbol{SymbolExtends, SymbolIdentifier}, 0),
		newProduction(SymbolExtendsDelcaration, []Symbol{SymbolExtends, SymbolComma, SymbolIdentifier}, 0),
	}

	pm.productionMap[SymbolClassDeclaration] = []*Production{
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolLc, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolLc, SymbolClassStatementList, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolLc, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolLc, SymbolClassStatementList, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolRc}, 0),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0),
	}

	pm.productionMap[SymbolInterfaceMethodDeclarationStatement] = []*Production{
		newProduction(SymbolInterfaceMethodDeclarationStatement, []Symbol{SymbolReturnValType, SymbolIdentifier, SymbolLp, SymbolRp}, 0),
		newProduction(SymbolInterfaceMethodDeclarationStatement, []Symbol{SymbolReturnValType, SymbolIdentifier, SymbolLp, SymbolParameterList, SymbolRp}, 0),
	}

	pm.productionMap[SymbolInterfaceMethodDeclarationStatementList] = []*Production{
		newProduction(SymbolInterfaceMethodDeclarationStatementList, []Symbol{SymbolInterfaceMethodDeclarationStatement}, 0),
		newProduction(SymbolInterfaceMethodDeclarationStatementList, []Symbol{SymbolInterfaceMethodDeclarationStatementList, SymbolInterfaceMethodDeclarationStatement}, 0),
	}

	pm.productionMap[SymbolInterfaceDeclaration] = []*Production{
		newProduction(SymbolInterfaceDeclaration, []Symbol{SymbolInterface, SymbolIdentifier, SymbolLc, SymbolRc}, 0),
		newProduction(SymbolInterfaceDeclaration, []Symbol{SymbolInterface, SymbolIdentifier, SymbolLc, SymbolInterfaceMethodDeclarationStatementList, SymbolRc}, 0),
	}

	pm.productionMap[SymbolClassInterfaceDeclaration] = []*Production{
		newProduction(SymbolClassInterfaceDeclaration, []Symbol{SymbolClassDeclaration}, 0),
		newProduction(SymbolClassInterfaceDeclaration, []Symbol{SymbolInterfaceDeclaration}, 0),
	}

	pm.productionMap[SymbolClassInterfaceDeclarationList] = []*Production{
		newProduction(SymbolClassInterfaceDeclarationList, []Symbol{SymbolClassInterfaceDeclaration}, 0),
		newProduction(SymbolClassInterfaceDeclarationList, []Symbol{SymbolClassInterfaceDeclarationList, SymbolClassInterfaceDeclaration}, 0),
	}

	pm.productionMap[SymbolTranslationUnit] = []*Production{
		newProduction(SymbolTranslationUnit, []Symbol{SymbolClassInterfaceDeclarationList}, 0),
	}

	return
}

// 获取以左侧为left的生成式列表
func (pm *ProductionManager) getProductions(left Symbol) (productions []*Production) {
	return pm.productionMap[left]
}
