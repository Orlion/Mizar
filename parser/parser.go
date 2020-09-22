package parser

import (
	"mizar/ast"
	"mizar/lexer"
	"strconv"

	"github.com/Orlion/merak"
	merak_ast "github.com/Orlion/merak/ast"
	"github.com/Orlion/merak/symbol"
)

type Parser struct {
	p *merak.Parser
}

func NewParser() *Parser {
	parser := &Parser{p: merak.NewParser()}
	parser.initProductions()
	return parser
}

// 自底向上分析
func (parser *Parser) Parse(l *lexer.Lexer) (tu *ast.TranslationUnit, err error) {
	astInter, err := parser.p.Build(lexer.SymbolTranslationUnit, lexer.EOISymbol).SetLexer(l).Parse()
	if err != nil {
		return
	}

	tu = astInter.(*ast.TranslationUnit)

	return
}

func (parser *Parser) initProductions() {
	parser.p.RegisterProduction(lexer.SymbolArgumentList, []symbol.Symbol{lexer.SymbolExpression}, false, func(args []interface{}) merak_ast.Node {
		argumentList := new(ast.ArgumentList)
		expr := args[0].(*ast.Expression)
		argumentList.List = append(argumentList.List, expr)

		return argumentList
	})

	parser.p.RegisterProduction(lexer.SymbolArgumentList, []symbol.Symbol{lexer.SymbolArgumentList, lexer.SymbolComma, lexer.SymbolExpression}, false, func(args []interface{}) merak_ast.Node {
		argumentList := args[0].(*ast.ArgumentList)

		expr := args[2].(*ast.Expression)

		argumentList.List = append(argumentList.List, expr)

		return argumentList
	})

	parser.p.RegisterProduction(lexer.SymbolMethodCall, []symbol.Symbol{lexer.SymbolIdentifier, lexer.SymbolLp, lexer.SymbolRp}, false, func(args []interface{}) merak_ast.Node {
		return &ast.MethodCall{Name: args[0].(*lexer.Token).Lexeme, ArgumentList: nil}
	})
	parser.p.RegisterProduction(lexer.SymbolMethodCall, []symbol.Symbol{lexer.SymbolIdentifier, lexer.SymbolLp, lexer.SymbolArgumentList, lexer.SymbolRp}, false, func(args []interface{}) merak_ast.Node {
		return &ast.MethodCall{Name: args[0].(*lexer.Token).Lexeme, ArgumentList: args[2].(*ast.ArgumentList).List}
	})

	parser.p.RegisterProduction(lexer.SymbolNewObjExpression, []symbol.Symbol{lexer.SymbolNew, lexer.SymbolMethodCall}, false, func(args []interface{}) merak_ast.Node {
		methodCall := args[1].(*ast.MethodCall)
		return &ast.NewObjectExpression{Name: methodCall.Name, ArgumentList: methodCall.ArgumentList}
	})

	parser.p.RegisterProduction(lexer.SymbolVarCallExpression, []symbol.Symbol{lexer.SymbolIdentifier}, false, func(args []interface{}) merak_ast.Node {
		return &ast.VarCallExpression{Var: args[0].(*lexer.Token).Lexeme, Type: ast.VarCallExpressionTypeVar}
	})
	parser.p.RegisterProduction(lexer.SymbolVarCallExpression, []symbol.Symbol{lexer.SymbolThis}, false, func(args []interface{}) merak_ast.Node {
		return &ast.VarCallExpression{This: args[0].(*lexer.Token).Lexeme, Type: ast.VarCallExpressionTypeThis}
	})
	parser.p.RegisterProduction(lexer.SymbolVarCallExpression, []symbol.Symbol{lexer.SymbolCallExpression, lexer.SymbolDot, lexer.SymbolIdentifier}, false, func(args []interface{}) merak_ast.Node {
		return &ast.VarCallExpression{CallExpression: args[0].(*ast.CallExpression), Var: args[2].(*lexer.Token).Lexeme, Type: ast.VarCallExpressionTypeCall}
	})

	parser.p.RegisterProduction(lexer.SymbolMethodCallExpression, []symbol.Symbol{lexer.SymbolCallExpression, lexer.SymbolDot, lexer.SymbolMethodCall}, false, func(args []interface{}) merak_ast.Node {
		return &ast.MethodCallExpression{CallExpression: args[0].(*ast.CallExpression), Name: args[2].(*ast.MethodCall).Name, ArgumentList: args[2].(*ast.MethodCall).ArgumentList}
	})

	parser.p.RegisterProduction(lexer.SymbolCallExpression, []symbol.Symbol{lexer.SymbolMethodCallExpression}, false, func(args []interface{}) merak_ast.Node {
		return &ast.CallExpression{MethodCallExpression: args[0].(*ast.MethodCallExpression), Type: ast.CallExpressionTypeMethodCall}
	})
	parser.p.RegisterProduction(lexer.SymbolCallExpression, []symbol.Symbol{lexer.SymbolVarCallExpression}, false, func(args []interface{}) merak_ast.Node {
		return &ast.CallExpression{VarCallExpression: args[0].(*ast.VarCallExpression), Type: ast.CallExpressionTypeValCall}
	})

	parser.p.RegisterProduction(lexer.SymbolExpression, []symbol.Symbol{lexer.SymbolStringLiteral}, false, func(args []interface{}) merak_ast.Node {
		stringToken := args[0].(*lexer.Token)
		return &ast.Expression{StringLiteral: stringToken.Lexeme, Type: ast.ExpressionTypeString}
	})
	parser.p.RegisterProduction(lexer.SymbolExpression, []symbol.Symbol{lexer.SymbolIntLiteral}, false, func(args []interface{}) merak_ast.Node {
		intToken := args[0].(*lexer.Token)
		intVal, err := strconv.ParseInt(intToken.Lexeme, 10, 64)
		if err != nil {
			panic(err)
		}
		return &ast.Expression{IntLiteral: intVal, Type: ast.ExpressionTypeInt}
	})
	parser.p.RegisterProduction(lexer.SymbolExpression, []symbol.Symbol{lexer.SymbolDoubleLiteral}, false, func(args []interface{}) merak_ast.Node {
		doubleToken := args[0].(*lexer.Token)
		floatVal, err := strconv.ParseFloat(doubleToken.Lexeme, 10)
		if err != nil {
			panic(err)
		}
		return &ast.Expression{DoubleLiteral: floatVal, Type: ast.ExpressionTypeDouble}
	})
	parser.p.RegisterProduction(lexer.SymbolExpression, []symbol.Symbol{lexer.SymbolNull}, false, func(args []interface{}) merak_ast.Node {
		return &ast.Expression{NullLiteral: nil, Type: ast.ExpressionTypeNull}
	})
	parser.p.RegisterProduction(lexer.SymbolExpression, []symbol.Symbol{lexer.SymbolTrue}, false, func(args []interface{}) merak_ast.Node {
		return &ast.Expression{BoolLiteral: true, Type: ast.ExpressionTypeBool}
	})
	parser.p.RegisterProduction(lexer.SymbolExpression, []symbol.Symbol{lexer.SymbolFalse}, false, func(args []interface{}) merak_ast.Node {
		return &ast.Expression{BoolLiteral: false, Type: ast.ExpressionTypeBool}
	})
	parser.p.RegisterProduction(lexer.SymbolExpression, []symbol.Symbol{lexer.SymbolNewObjExpression}, false, func(args []interface{}) merak_ast.Node {
		return &ast.Expression{NewObjectExpression: args[0].(*ast.NewObjectExpression), Type: ast.ExpressionTypeNewObject}
	})
	parser.p.RegisterProduction(lexer.SymbolExpression, []symbol.Symbol{lexer.SymbolCallExpression}, false, func(args []interface{}) merak_ast.Node {
		return &ast.Expression{CallExpression: args[0].(*ast.CallExpression), Type: ast.ExpressionTypeCall}
	})

	parser.p.RegisterProduction(lexer.SymbolTypeVar, []symbol.Symbol{lexer.SymbolVoid, lexer.SymbolIdentifier}, false, func(args []interface{}) merak_ast.Node {
		return &ast.TypeVar{Type: "void", Name: args[1].(*lexer.Token).Lexeme}
	})
	parser.p.RegisterProduction(lexer.SymbolTypeVar, []symbol.Symbol{lexer.SymbolIdentifier, lexer.SymbolIdentifier}, false, func(args []interface{}) merak_ast.Node {
		return &ast.TypeVar{Type: args[0].(*lexer.Token).Lexeme, Name: args[1].(*lexer.Token).Lexeme}
	})

	parser.p.RegisterProduction(lexer.SymbolExpressionStatement, []symbol.Symbol{lexer.SymbolExpression, lexer.SymbolSemicolon}, false, func(args []interface{}) merak_ast.Node {
		expr := args[0].(*ast.Expression)
		return &ast.ExpressionStatement{Expression: expr}
	})

	parser.p.RegisterProduction(lexer.SymbolVarAssignStatement, []symbol.Symbol{lexer.SymbolTypeVar, lexer.SymbolAssign, lexer.SymbolExpressionStatement}, false, func(args []interface{}) merak_ast.Node {
		typeVar := args[0].(*ast.TypeVar)
		exprStmt := args[2].(*ast.ExpressionStatement)
		return &ast.VarAssignStatement{VarName: typeVar.Name, VarType: typeVar.Type, Expression: exprStmt.Expression, Type: ast.VarAssignStatementTypeVar}
	})
	parser.p.RegisterProduction(lexer.SymbolVarAssignStatement, []symbol.Symbol{lexer.SymbolVarCallExpression, lexer.SymbolAssign, lexer.SymbolExpressionStatement}, false, func(args []interface{}) merak_ast.Node {
		varCallExpr := args[0].(*ast.VarCallExpression)
		exprStmt := args[2].(*ast.ExpressionStatement)
		return &ast.VarAssignStatement{VarCallExpression: varCallExpr, Expression: exprStmt.Expression, Type: ast.VarAssignStatementTypeVarCall}
	})

	parser.p.RegisterProduction(lexer.SymbolVarDeclarationStatement, []symbol.Symbol{lexer.SymbolTypeVar, lexer.SymbolSemicolon}, false, func(args []interface{}) merak_ast.Node {
		typeVar := args[0].(*ast.TypeVar)
		return &ast.VarDeclarationStatement{Type: typeVar.Type, Name: typeVar.Name}
	})

	parser.p.RegisterProduction(lexer.SymbolReturnStatement, []symbol.Symbol{lexer.SymbolReturn, lexer.SymbolSemicolon}, false, func(args []interface{}) merak_ast.Node {
		return &ast.ReturnStatement{}
	})
	parser.p.RegisterProduction(lexer.SymbolReturnStatement, []symbol.Symbol{lexer.SymbolReturn, lexer.SymbolExpressionStatement}, false, func(args []interface{}) merak_ast.Node {
		exprStmt := args[1].(*ast.ExpressionStatement)
		return &ast.ReturnStatement{Expression: exprStmt.Expression}
	})

	parser.p.RegisterProduction(lexer.SymbolContinueStatement, []symbol.Symbol{lexer.SymbolContinue, lexer.SymbolSemicolon}, false, func(args []interface{}) merak_ast.Node {
		return &ast.ContinueStatement{}
	})
	parser.p.RegisterProduction(lexer.SymbolContinueStatement, []symbol.Symbol{lexer.SymbolContinue, lexer.SymbolExpressionStatement}, false, func(args []interface{}) merak_ast.Node {
		exprStmt := args[0].(*ast.ExpressionStatement)
		return &ast.ContinueStatement{Expression: exprStmt.Expression}
	})

	parser.p.RegisterProduction(lexer.SymbolBreakStatement, []symbol.Symbol{lexer.SymbolBreak, lexer.SymbolSemicolon}, false, func(args []interface{}) merak_ast.Node {
		return &ast.BreakStatement{}
	})
	parser.p.RegisterProduction(lexer.SymbolBreakStatement, []symbol.Symbol{lexer.SymbolBreak, lexer.SymbolExpressionStatement}, false, func(args []interface{}) merak_ast.Node {
		exprStmt := args[0].(*ast.ExpressionStatement)
		return &ast.BreakStatement{Expression: exprStmt.Expression}
	})

	parser.p.RegisterProduction(lexer.SymbolForStatement, []symbol.Symbol{lexer.SymbolFor, lexer.SymbolLp, lexer.SymbolSemicolon, lexer.SymbolSemicolon, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		block := args[5].(*ast.Block)
		return &ast.ForStatement{Block: block}
	})
	parser.p.RegisterProduction(lexer.SymbolForStatement, []symbol.Symbol{lexer.SymbolFor, lexer.SymbolLp, lexer.SymbolExpression, lexer.SymbolSemicolon, lexer.SymbolSemicolon, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		initExpr := args[2].(*ast.Expression)
		block := args[6].(*ast.Block)
		return &ast.ForStatement{InitExpression: initExpr, Block: block}
	})
	parser.p.RegisterProduction(lexer.SymbolForStatement, []symbol.Symbol{lexer.SymbolFor, lexer.SymbolLp, lexer.SymbolExpression, lexer.SymbolSemicolon, lexer.SymbolExpression, lexer.SymbolSemicolon, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		initExpr := args[2].(*ast.Expression)
		condExpr := args[4].(*ast.Expression)
		block := args[7].(*ast.Block)
		return &ast.ForStatement{InitExpression: initExpr, CondExpression: condExpr, Block: block}
	})
	parser.p.RegisterProduction(lexer.SymbolForStatement, []symbol.Symbol{lexer.SymbolFor, lexer.SymbolLp, lexer.SymbolExpression, lexer.SymbolSemicolon, lexer.SymbolExpression, lexer.SymbolSemicolon, lexer.SymbolExpression, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		initExpr := args[2].(*ast.Expression)
		condExpr := args[4].(*ast.Expression)
		postExpr := args[6].(*ast.Expression)
		block := args[8].(*ast.Block)
		return &ast.ForStatement{InitExpression: initExpr, CondExpression: condExpr, PostExpression: postExpr, Block: block}
	})

	parser.p.RegisterProduction(lexer.SymbolForStatement, []symbol.Symbol{lexer.SymbolFor, lexer.SymbolLp, lexer.SymbolExpression, lexer.SymbolSemicolon, lexer.SymbolSemicolon, lexer.SymbolExpression, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		initExpr := args[2].(*ast.Expression)
		postExpr := args[5].(*ast.Expression)
		block := args[7].(*ast.Block)
		return &ast.ForStatement{InitExpression: initExpr, PostExpression: postExpr, Block: block}
	})
	parser.p.RegisterProduction(lexer.SymbolForStatement, []symbol.Symbol{lexer.SymbolFor, lexer.SymbolLp, lexer.SymbolSemicolon, lexer.SymbolSemicolon, lexer.SymbolExpression, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		postExpr := args[4].(*ast.Expression)
		block := args[6].(*ast.Block)
		return &ast.ForStatement{PostExpression: postExpr, Block: block}
	})
	parser.p.RegisterProduction(lexer.SymbolForStatement, []symbol.Symbol{lexer.SymbolFor, lexer.SymbolLp, lexer.SymbolSemicolon, lexer.SymbolExpression, lexer.SymbolSemicolon, lexer.SymbolExpression, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		condExpr := args[3].(*ast.Expression)
		postExpr := args[5].(*ast.Expression)
		block := args[7].(*ast.Block)
		return &ast.ForStatement{CondExpression: condExpr, PostExpression: postExpr, Block: block}
	})
	parser.p.RegisterProduction(lexer.SymbolForStatement, []symbol.Symbol{lexer.SymbolFor, lexer.SymbolLp, lexer.SymbolSemicolon, lexer.SymbolExpression, lexer.SymbolSemicolon, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		condExpr := args[4].(*ast.Expression)
		block := args[8].(*ast.Block)
		return &ast.ForStatement{CondExpression: condExpr, Block: block}
	})

	parser.p.RegisterProduction(lexer.SymbolIfStatement, []symbol.Symbol{lexer.SymbolIf, lexer.SymbolLp, lexer.SymbolExpression, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		expr := args[2].(*ast.Expression)
		ifBlock := args[4].(*ast.Block)
		return &ast.IfStatement{CondExpression: expr, IfBlock: ifBlock}
	})
	parser.p.RegisterProduction(lexer.SymbolIfStatement, []symbol.Symbol{lexer.SymbolIf, lexer.SymbolLp, lexer.SymbolExpression, lexer.SymbolRp, lexer.SymbolBlock, lexer.SymbolElse, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		expr := args[2].(*ast.Expression)
		ifBlock := args[4].(*ast.Block)
		elseBlock := args[6].(*ast.Block)
		return &ast.IfStatement{CondExpression: expr, IfBlock: ifBlock, ElseBlock: elseBlock}
	})

	parser.p.RegisterProduction(lexer.SymbolWhileStatement, []symbol.Symbol{lexer.SymbolWhile, lexer.SymbolLp, lexer.SymbolExpression, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		expr := args[2].(*ast.Expression)
		block := args[4].(*ast.Block)
		return &ast.WhileStatement{Expression: expr, Block: block}
	})

	parser.p.RegisterProduction(lexer.SymbolStatement, []symbol.Symbol{lexer.SymbolExpressionStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.ExpressionStatement)
		return &ast.Statement{ExpressionStatement: stmt, Type: ast.StatementTypeExpression}
	})
	parser.p.RegisterProduction(lexer.SymbolStatement, []symbol.Symbol{lexer.SymbolVarDeclarationStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.VarDeclarationStatement)
		return &ast.Statement{VarDeclarationStatement: stmt, Type: ast.StatementTypeVarDeclaration}
	})
	parser.p.RegisterProduction(lexer.SymbolStatement, []symbol.Symbol{lexer.SymbolVarAssignStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.VarAssignStatement)
		return &ast.Statement{VarAssignStatement: stmt, Type: ast.StatementTypeVarAssign}
	})
	parser.p.RegisterProduction(lexer.SymbolStatement, []symbol.Symbol{lexer.SymbolWhileStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.WhileStatement)
		return &ast.Statement{WhileStatement: stmt, Type: ast.StatementTypeWhile}
	})
	parser.p.RegisterProduction(lexer.SymbolStatement, []symbol.Symbol{lexer.SymbolIfStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.IfStatement)
		return &ast.Statement{IfStatement: stmt, Type: ast.StatementTypeIf}
	})
	parser.p.RegisterProduction(lexer.SymbolStatement, []symbol.Symbol{lexer.SymbolForStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.ForStatement)
		return &ast.Statement{ForStatement: stmt, Type: ast.StatementTypeFor}
	})
	parser.p.RegisterProduction(lexer.SymbolStatement, []symbol.Symbol{lexer.SymbolBreakStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.BreakStatement)
		return &ast.Statement{BreakStatement: stmt, Type: ast.StatementTypeBreak}
	})
	parser.p.RegisterProduction(lexer.SymbolStatement, []symbol.Symbol{lexer.SymbolContinueStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.ContinueStatement)
		return &ast.Statement{ContinueStatement: stmt, Type: ast.StatementTypeContinue}
	})
	parser.p.RegisterProduction(lexer.SymbolStatement, []symbol.Symbol{lexer.SymbolReturnStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.ReturnStatement)
		return &ast.Statement{ReturnStatement: stmt, Type: ast.StatementTypeReturn}
	})

	parser.p.RegisterProduction(lexer.SymbolStatementList, []symbol.Symbol{lexer.SymbolStatement}, false, func(args []interface{}) merak_ast.Node {
		stmt := args[0].(*ast.Statement)
		stmtList := new(ast.StatementList)
		stmtList.List = append(stmtList.List, stmt)
		return stmtList
	})
	parser.p.RegisterProduction(lexer.SymbolStatementList, []symbol.Symbol{lexer.SymbolStatementList, lexer.SymbolStatement}, false, func(args []interface{}) merak_ast.Node {
		stmtList := args[0].(*ast.StatementList)
		stmt := args[1].(*ast.Statement)
		stmtList.List = append(stmtList.List, stmt)
		return stmtList
	})

	parser.p.RegisterProduction(lexer.SymbolEmptyBlock, []symbol.Symbol{lexer.SymbolLc, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		return &ast.Block{}
	})

	parser.p.RegisterProduction(lexer.SymbolBlock, []symbol.Symbol{lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		return &ast.Block{}
	})
	parser.p.RegisterProduction(lexer.SymbolBlock, []symbol.Symbol{lexer.SymbolLc, lexer.SymbolStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		stmtList := args[1].(*ast.StatementList)
		return &ast.Block{StatementList: stmtList.List}
	})

	parser.p.RegisterProduction(lexer.SymbolParameterList, []symbol.Symbol{lexer.SymbolTypeVar}, false, func(args []interface{}) merak_ast.Node {
		typeVar := args[0].(*ast.TypeVar)
		param := new(ast.Parameter)
		param.Name = typeVar.Name
		param.Type = typeVar.Type
		paramList := new(ast.ParameterList)
		paramList.List = append(paramList.List, param)
		return paramList
	})
	parser.p.RegisterProduction(lexer.SymbolParameterList, []symbol.Symbol{lexer.SymbolParameterList, lexer.SymbolComma, lexer.SymbolTypeVar}, false, func(args []interface{}) merak_ast.Node {
		paramList := args[0].(*ast.ParameterList)
		typeVar := args[2].(*ast.TypeVar)
		param := new(ast.Parameter)
		param.Name = typeVar.Name
		param.Type = typeVar.Type
		paramList.List = append(paramList.List, param)
		return paramList
	})

	parser.p.RegisterProduction(lexer.SymbolMemberModifier, []symbol.Symbol{lexer.SymbolPublic}, false, func(args []interface{}) merak_ast.Node {
		return &ast.MemberModifier{Type: ast.ModifierPublic}
	})
	parser.p.RegisterProduction(lexer.SymbolMemberModifier, []symbol.Symbol{lexer.SymbolProtected}, false, func(args []interface{}) merak_ast.Node {
		return &ast.MemberModifier{Type: ast.ModifierProtected}
	})
	parser.p.RegisterProduction(lexer.SymbolMemberModifier, []symbol.Symbol{lexer.SymbolPrivate}, false, func(args []interface{}) merak_ast.Node {
		return &ast.MemberModifier{Type: ast.ModifierPrivate}
	})
	parser.p.RegisterProduction(lexer.SymbolMemberModifier, []symbol.Symbol{lexer.SymbolAbstract}, false, func(args []interface{}) merak_ast.Node {
		return &ast.MemberModifier{Type: ast.ModifierAbstract}
	})

	parser.p.RegisterProduction(lexer.SymbolMethodDefinition, []symbol.Symbol{lexer.SymbolMemberModifier, lexer.SymbolTypeVar, lexer.SymbolLp, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		modifier := args[0].(*ast.MemberModifier)
		typeVar := args[1].(*ast.TypeVar)
		block := args[4].(*ast.Block)
		return &ast.MethodDefinition{ModifierType: modifier.Type, Name: typeVar.Name, Type: typeVar.Type, Block: block}
	})
	parser.p.RegisterProduction(lexer.SymbolMethodDefinition, []symbol.Symbol{lexer.SymbolMemberModifier, lexer.SymbolTypeVar, lexer.SymbolLp, lexer.SymbolParameterList, lexer.SymbolRp, lexer.SymbolBlock}, false, func(args []interface{}) merak_ast.Node {
		modifier := args[0].(*ast.MemberModifier)
		typeVar := args[1].(*ast.TypeVar)
		paramList := args[3].(*ast.ParameterList)
		block := args[5].(*ast.Block)
		return &ast.MethodDefinition{ModifierType: modifier.Type, Name: typeVar.Name, Type: typeVar.Type, ParameterList: paramList.List, Block: block}
	})

	parser.p.RegisterProduction(lexer.SymbolPropertyDefinition, []symbol.Symbol{lexer.SymbolMemberModifier, lexer.SymbolTypeVar, lexer.SymbolSemicolon}, false, func(args []interface{}) merak_ast.Node {
		modifier := args[0].(*ast.MemberModifier)
		typeVar := args[1].(*ast.TypeVar)
		return &ast.PropertyDefinition{ModifierType: modifier.Type, Name: typeVar.Name, Type: typeVar.Type}
	})
	parser.p.RegisterProduction(lexer.SymbolPropertyDefinition, []symbol.Symbol{lexer.SymbolMemberModifier, lexer.SymbolTypeVar, lexer.SymbolAssign, lexer.SymbolExpressionStatement}, false, func(args []interface{}) merak_ast.Node {
		modifier := args[0].(*ast.MemberModifier)
		typeVar := args[1].(*ast.TypeVar)
		exprStmt := args[3].(*ast.ExpressionStatement)
		return &ast.PropertyDefinition{ModifierType: modifier.Type, Name: typeVar.Name, Type: typeVar.Type, Expr: exprStmt.Expression}
	})

	parser.p.RegisterProduction(lexer.SymbolClassStatement, []symbol.Symbol{lexer.SymbolMethodDefinition}, false, func(args []interface{}) merak_ast.Node {
		md := args[0].(*ast.MethodDefinition)
		return &ast.ClassStatement{MethodDefinition: md, Type: ast.ClassStatementTypeMethod}
	})
	parser.p.RegisterProduction(lexer.SymbolClassStatement, []symbol.Symbol{lexer.SymbolPropertyDefinition}, false, func(args []interface{}) merak_ast.Node {
		pd := args[0].(*ast.PropertyDefinition)
		return &ast.ClassStatement{PropertyDefinition: pd, Type: ast.ClassStatementTypeProperty}
	})

	parser.p.RegisterProduction(lexer.SymbolClassStatementList, []symbol.Symbol{lexer.SymbolClassStatement}, false, func(args []interface{}) merak_ast.Node {
		cs := args[0].(*ast.ClassStatement)
		csl := new(ast.ClassStatementList)
		csl.MethodDefinitionMap = make(map[string]map[string]*ast.MethodDefinition)
		csl.AbstractMethodDefinitionMap = make(map[string]map[string]*ast.MethodDefinition)
		switch cs.Type {
		case ast.ClassStatementTypeMethod:
			csl.MethodDefinitionMap[cs.MethodDefinition.Name] = cs.MethodDefinition
		case ast.ClassStatementTypeAbstractMethod:
			csl.AbstractMethodDefinitionList = append(csl.AbstractMethodDefinitionList, cs.MethodDefinition)
		case ast.ClassStatementTypeProperty:
			csl.PropertyDefinitionList = append(csl.PropertyDefinitionList, cs.PropertyDefinition)
		}

		return csl
	})
	parser.p.RegisterProduction(lexer.SymbolClassStatementList, []symbol.Symbol{lexer.SymbolClassStatementList, lexer.SymbolClassStatement}, false, func(args []interface{}) merak_ast.Node {
		csl := args[0].(*ast.ClassStatementList)
		cs := args[1].(*ast.ClassStatement)
		switch cs.Type {
		case ast.ClassStatementTypeMethod:
			csl.MethodDefinitionList = append(csl.MethodDefinitionList, cs.MethodDefinition)
		case ast.ClassStatementTypeAbstractMethod:
			csl.AbstractMethodDefinitionList = append(csl.AbstractMethodDefinitionList, cs.MethodDefinition)
		case ast.ClassStatementTypeProperty:
			csl.PropertyDefinitionList = append(csl.PropertyDefinitionList, cs.PropertyDefinition)
		}
		return csl
	})

	parser.p.RegisterProduction(lexer.SymbolImplementsDeclaration, []symbol.Symbol{lexer.SymbolImplements, lexer.SymbolIdentifier}, false, func(args []interface{}) merak_ast.Node {
		t := args[1].(*lexer.Token)
		impl := new(ast.Implements)
		impl.InterfaceNameList = append(impl.InterfaceNameList, t.Lexeme)
		return impl
	})
	parser.p.RegisterProduction(lexer.SymbolImplementsDeclaration, []symbol.Symbol{lexer.SymbolImplementsDeclaration, lexer.SymbolComma, lexer.SymbolIdentifier}, false, func(args []interface{}) merak_ast.Node {
		impl := args[0].(*ast.Implements)
		t := args[2].(*lexer.Token)
		impl.InterfaceNameList = append(impl.InterfaceNameList, t.Lexeme)
		return impl
	})

	parser.p.RegisterProduction(lexer.SymbolExtendsDelcaration, []symbol.Symbol{lexer.SymbolExtends, lexer.SymbolIdentifier}, false, func(args []interface{}) merak_ast.Node {
		t := args[1].(*lexer.Token)
		extends := new(ast.Extends)
		extends.ClassNameList = append(extends.ClassNameList, t.Lexeme)
		return extends
	})
	parser.p.RegisterProduction(lexer.SymbolExtendsDelcaration, []symbol.Symbol{lexer.SymbolExtends, lexer.SymbolComma, lexer.SymbolIdentifier}, false, func(args []interface{}) merak_ast.Node {
		extends := args[0].(*ast.Extends)
		t := args[2].(*lexer.Token)
		extends.ClassNameList = append(extends.ClassNameList, t.Lexeme)
		return extends
	})

	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		return &ast.Class{Name: nameT.Lexeme}
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolLc, lexer.SymbolClassStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		csl := args[3].(*ast.ClassStatementList)
		class := new(ast.Class)
		class.Name = nameT.Lexeme
		class.MethodDefinitionList = csl.MethodDefinitionList
		class.AbstractMethodDefinitionList = csl.AbstractMethodDefinitionList
		class.PropertyDefinitionList = csl.PropertyDefinitionList
		return class
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolAbstract, lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[2].(*lexer.Token)
		class := new(ast.Class)
		class.IsAbstract = true
		class.Name = nameT.Lexeme
		return class
	})

	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolAbstract, lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolLc, lexer.SymbolClassStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[2].(*lexer.Token)
		csl := args[4].(*ast.ClassStatementList)
		class := new(ast.Class)
		class.IsAbstract = true
		class.Name = nameT.Lexeme
		class.MethodDefinitionList = csl.MethodDefinitionList
		class.AbstractMethodDefinitionList = csl.AbstractMethodDefinitionList
		class.PropertyDefinitionList = csl.PropertyDefinitionList
		return class
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolExtendsDelcaration, lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		extends := args[2].(*ast.Extends)
		class := new(ast.Class)
		class.Name = nameT.Lexeme
		class.Extends = extends.ClassNameList
		return class
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolExtendsDelcaration, lexer.SymbolLc, lexer.SymbolClassStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		extends := args[2].(*ast.Extends)
		csl := args[4].(*ast.ClassStatementList)
		class := new(ast.Class)
		class.Name = nameT.Lexeme
		class.Extends = extends.ClassNameList
		class.MethodDefinitionList = csl.MethodDefinitionList
		class.AbstractMethodDefinitionList = csl.AbstractMethodDefinitionList
		class.PropertyDefinitionList = csl.PropertyDefinitionList
		return class
	})

	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolAbstract, lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolExtendsDelcaration, lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[2].(*lexer.Token)
		extends := args[3].(*ast.Extends)
		class := new(ast.Class)
		class.IsAbstract = true
		class.Name = nameT.Lexeme
		class.Extends = extends.ClassNameList
		return class
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolAbstract, lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolExtendsDelcaration, lexer.SymbolLc, lexer.SymbolClassStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[2].(*lexer.Token)
		extends := args[3].(*ast.Extends)
		csl := args[5].(*ast.ClassStatementList)
		class := new(ast.Class)
		class.IsAbstract = true
		class.Name = nameT.Lexeme
		class.Extends = extends.ClassNameList
		class.MethodDefinitionList = csl.MethodDefinitionList
		class.AbstractMethodDefinitionList = csl.AbstractMethodDefinitionList
		class.PropertyDefinitionList = csl.PropertyDefinitionList
		return class
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolImplementsDeclaration, lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		implements := args[2].(*ast.Implements)
		class := new(ast.Class)
		class.Name = nameT.Lexeme
		class.Implements = implements.InterfaceNameList
		return class
	})

	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolImplementsDeclaration, lexer.SymbolLc, lexer.SymbolClassStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		implements := args[2].(*ast.Implements)
		csl := args[4].(*ast.ClassStatementList)
		class := new(ast.Class)
		class.Name = nameT.Lexeme
		class.Implements = implements.InterfaceNameList
		class.MethodDefinitionMap = csl.MethodDefinitionMap
		class.AbstractMethodDefinitionMap = csl.AbstractMethodDefinitionMap
		class.PropertyDefinitionMap = csl.PropertyDefinitionMap
		return class
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolAbstract, lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolImplementsDeclaration, lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[2].(*lexer.Token)
		implements := args[3].(*ast.Implements)
		class := new(ast.Class)
		class.IsAbstract = true
		class.Name = nameT.Lexeme
		class.Implements = implements.InterfaceNameList
		return class
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolAbstract, lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolImplementsDeclaration, lexer.SymbolLc, lexer.SymbolClassStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[2].(*lexer.Token)
		implements := args[3].(*ast.Implements)
		csl := args[5].(*ast.ClassStatementList)
		class := new(ast.Class)
		class.IsAbstract = true
		class.Name = nameT.Lexeme
		class.Implements = implements.InterfaceNameList
		class.MethodDefinitionList = csl.MethodDefinitionList
		class.AbstractMethodDefinitionList = csl.AbstractMethodDefinitionList
		class.PropertyDefinitionList = csl.PropertyDefinitionList
		return class
	})

	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolExtendsDelcaration, lexer.SymbolImplementsDeclaration, lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		extends := args[2].(*ast.Extends)
		implements := args[3].(*ast.Implements)
		class := new(ast.Class)
		class.Name = nameT.Lexeme
		class.Extends = extends.ClassNameList
		class.Implements = implements.InterfaceNameList
		return class
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolExtendsDelcaration, lexer.SymbolImplementsDeclaration, lexer.SymbolLc, lexer.SymbolClassStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		extends := args[2].(*ast.Extends)
		implements := args[3].(*ast.Implements)
		csl := args[5].(*ast.ClassStatementList)
		class := new(ast.Class)
		class.Name = nameT.Lexeme
		class.Extends = extends.ClassNameList
		class.Implements = implements.InterfaceNameList
		class.MethodDefinitionList = csl.MethodDefinitionList
		class.AbstractMethodDefinitionList = csl.AbstractMethodDefinitionList
		class.PropertyDefinitionList = csl.PropertyDefinitionList
		return class
	})
	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolAbstract, lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolExtendsDelcaration, lexer.SymbolImplementsDeclaration, lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[2].(*lexer.Token)
		extends := args[3].(*ast.Extends)
		implements := args[4].(*ast.Implements)
		class := new(ast.Class)
		class.IsAbstract = true
		class.Name = nameT.Lexeme
		class.Extends = extends.ClassNameList
		class.Implements = implements.InterfaceNameList
		return class
	})

	parser.p.RegisterProduction(lexer.SymbolClassDeclaration, []symbol.Symbol{lexer.SymbolAbstract, lexer.SymbolClass, lexer.SymbolIdentifier, lexer.SymbolExtendsDelcaration, lexer.SymbolImplementsDeclaration, lexer.SymbolLc, lexer.SymbolClassStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[2].(*lexer.Token)
		extends := args[3].(*ast.Extends)
		implements := args[4].(*ast.Implements)
		csl := args[6].(*ast.ClassStatementList)
		class := new(ast.Class)
		class.IsAbstract = true
		class.Name = nameT.Lexeme
		class.Extends = extends.ClassNameList
		class.Implements = implements.InterfaceNameList
		class.MethodDefinitionList = csl.MethodDefinitionList
		class.AbstractMethodDefinitionList = csl.AbstractMethodDefinitionList
		class.PropertyDefinitionList = csl.PropertyDefinitionList
		return class
	})

	parser.p.RegisterProduction(lexer.SymbolInterfaceMethodDeclarationStatement, []symbol.Symbol{lexer.SymbolTypeVar, lexer.SymbolLp, lexer.SymbolRp, lexer.SymbolSemicolon}, false, func(args []interface{}) merak_ast.Node {
		typeVar := args[0].(*ast.TypeVar)
		return &ast.InterfaceMethod{Type: typeVar.Type, Name: typeVar.Name}
	})
	parser.p.RegisterProduction(lexer.SymbolInterfaceMethodDeclarationStatement, []symbol.Symbol{lexer.SymbolTypeVar, lexer.SymbolLp, lexer.SymbolParameterList, lexer.SymbolRp, lexer.SymbolSemicolon}, false, func(args []interface{}) merak_ast.Node {
		typeVar := args[0].(*ast.TypeVar)
		paramList := args[2].(*ast.ParameterList)
		return &ast.InterfaceMethod{Type: typeVar.Type, Name: typeVar.Name, ParameterList: paramList.List}
	})

	parser.p.RegisterProduction(lexer.SymbolInterfaceMethodDeclarationStatementList, []symbol.Symbol{lexer.SymbolInterfaceMethodDeclarationStatement}, false, func(args []interface{}) merak_ast.Node {
		im := args[0].(*ast.InterfaceMethod)
		return &ast.InterfaceMethodList{List: []*ast.InterfaceMethod{im}}
	})
	parser.p.RegisterProduction(lexer.SymbolInterfaceMethodDeclarationStatementList, []symbol.Symbol{lexer.SymbolInterfaceMethodDeclarationStatementList, lexer.SymbolInterfaceMethodDeclarationStatement}, false, func(args []interface{}) merak_ast.Node {
		iml := args[0].(*ast.InterfaceMethodList)
		im := args[1].(*ast.InterfaceMethod)
		iml.List = append(iml.List, im)
		return iml
	})

	parser.p.RegisterProduction(lexer.SymbolInterfaceDeclaration, []symbol.Symbol{lexer.SymbolInterface, lexer.SymbolIdentifier, lexer.SymbolEmptyBlock}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		return &ast.Interface{Name: nameT.Lexeme}
	})
	parser.p.RegisterProduction(lexer.SymbolInterfaceDeclaration, []symbol.Symbol{lexer.SymbolInterface, lexer.SymbolIdentifier, lexer.SymbolLc, lexer.SymbolInterfaceMethodDeclarationStatementList, lexer.SymbolRc}, false, func(args []interface{}) merak_ast.Node {
		nameT := args[1].(*lexer.Token)
		iml := args[3].(*ast.InterfaceMethodList)
		return &ast.Interface{Name: nameT.Lexeme, MethodList: iml.List}
	})

	parser.p.RegisterProduction(lexer.SymbolClassInterfaceDeclaration, []symbol.Symbol{lexer.SymbolClassDeclaration}, false, func(args []interface{}) merak_ast.Node {
		class := args[0].(*ast.Class)
		return &ast.ClassInterface{Class: class, Type: ast.ClassInterfaceTypeClass}
	})
	parser.p.RegisterProduction(lexer.SymbolClassInterfaceDeclaration, []symbol.Symbol{lexer.SymbolInterfaceDeclaration}, false, func(args []interface{}) merak_ast.Node {
		inter := args[0].(*ast.Interface)
		return &ast.ClassInterface{Interface: inter, Type: ast.ClassInterfaceTypeInterface}
	})

	parser.p.RegisterProduction(lexer.SymbolClassInterfaceDeclarationList, []symbol.Symbol{lexer.SymbolClassInterfaceDeclaration}, false, func(args []interface{}) merak_ast.Node {
		ci := args[0].(*ast.ClassInterface)
		tu := new(ast.TranslationUnit)
		switch ci.Type {
		case ast.ClassInterfaceTypeClass:
			tu.ClassList = append(tu.ClassList, ci.Class)
		case ast.ClassInterfaceTypeInterface:
			tu.InterfaceList = append(tu.InterfaceList, ci.Interface)
		}
		return tu
	})
	parser.p.RegisterProduction(lexer.SymbolClassInterfaceDeclarationList, []symbol.Symbol{lexer.SymbolClassInterfaceDeclarationList, lexer.SymbolClassInterfaceDeclaration}, false, func(args []interface{}) merak_ast.Node {
		tu := args[0].(*ast.TranslationUnit)
		ci := args[1].(*ast.ClassInterface)
		switch ci.Type {
		case ast.ClassInterfaceTypeClass:
			tu.ClassList = append(tu.ClassList, ci.Class)
		case ast.ClassInterfaceTypeInterface:
			tu.InterfaceList = append(tu.InterfaceList, ci.Interface)
		}
		return tu
	})

	parser.p.RegisterProduction(lexer.SymbolTranslationUnit, []symbol.Symbol{lexer.SymbolClassInterfaceDeclarationList}, false, func(args []interface{}) merak_ast.Node {
		tu := args[0].(*ast.TranslationUnit)
		return tu
	})
}
