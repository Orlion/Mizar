package parser

import (
	"mizar/ast"
	"mizar/lexer"
)

type ProductionManager struct {
	productionMap   map[Symbol][]*Production
	firstSetBuilder *FirstSetBuilder
}

var productionManager *ProductionManager

func getProductionManager() (pm *ProductionManager) {
	if productionManager != nil {
		return productionManager
	}

	pm = new(ProductionManager)

	pm.firstSetBuilder = newFirstSetBuilder()
	pm.firstSetBuilder.runFirstSets()

	pm.productionMap = make(map[Symbol][]*Production)

	pm.productionMap[SymbolArgumentList] = []*Production{
		newProduction(SymbolArgumentList, []Symbol{SymbolExpression}, 0, func(args []interface{}) ast.Node {
			argumentList := new(ast.ArgumentList)
			expr := args[0].(*ast.Expression)
			argumentList.List = append(argumentList.List, expr)

			return argumentList
		}),

		newProduction(SymbolArgumentList, []Symbol{SymbolArgumentList, SymbolComma, SymbolExpression}, 0, func(args []interface{}) ast.Node {
			argumentList := args[0].(*ast.ArgumentList)

			expr := args[2].(*ast.Expression)

			argumentList.List = append(argumentList.List, expr)

			return argumentList
		}),
	}

	pm.productionMap[SymbolMethodCall] = []*Production{
		newProduction(SymbolMethodCall, []Symbol{SymbolIdentifier, SymbolLp, SymbolRp}, 0, func(args []interface{}) ast.Node {
			return &ast.MethodCall{Name: args[0].(*lexer.Token).V, ArgumentList: nil}
		}),
		newProduction(SymbolMethodCall, []Symbol{SymbolIdentifier, SymbolLp, SymbolArgumentList, SymbolRp}, 0, func(args []interface{}) ast.Node {
			return &ast.MethodCall{Name: args[0].(*lexer.Token).V, ArgumentList: args[2].(ast.ArgumentList).List}
		}),
	}

	pm.productionMap[SymbolNewObjExpression] = []*Production{
		newProduction(SymbolNewObjExpression, []Symbol{SymbolNew, SymbolMethodCall}, 0, func(args []interface{}) ast.Node {
			methodCall := args[1].(*ast.MethodCall)
			return &ast.NewObjectExpression{Name: methodCall.Name, ArgumentList: methodCall.ArgumentList}
		}),
	}

	pm.productionMap[SymbolVarCallExpression] = []*Production{
		newProduction(SymbolVarCallExpression, []Symbol{SymbolIdentifier}, 0, func(args []interface{}) ast.Node {
			return &ast.VarCallExpression{Var: args[0].(lexer.Token).V, Type: ast.VarCallExpressionTypeVar}
		}),
		newProduction(SymbolVarCallExpression, []Symbol{SymbolThis}, 0, func(args []interface{}) ast.Node {
			return &ast.VarCallExpression{This: args[0].(lexer.Token).V, Type: ast.VarCallExpressionTypeThis}
		}),
		newProduction(SymbolVarCallExpression, []Symbol{SymbolCallExpression, SymbolDot, SymbolIdentifier}, 0, func(args []interface{}) ast.Node {
			return &ast.VarCallExpression{CallExpression: args[0].(*ast.CallExpression), Var: args[2].(*lexer.Token).V, Type: ast.VarCallExpressionTypeCall}
		}),
	}

	pm.productionMap[SymbolMethodCallExpression] = []*Production{
		newProduction(SymbolMethodCallExpression, []Symbol{SymbolCallExpression, SymbolDot, SymbolMethodCall}, 0, func(args []interface{}) ast.Node {
			return &ast.MethodCallExpression{CallExpression: args[0].(*ast.CallExpression), Name: args[2].(*ast.MethodCall).Name, ArgumentList: args[2].(*ast.MethodCall).ArgumentList}
		}),
	}

	pm.productionMap[SymbolCallExpression] = []*Production{
		newProduction(SymbolCallExpression, []Symbol{SymbolMethodCallExpression}, 0, func(args []interface{}) ast.Node {
			return &ast.CallExpression{MethodCallExpression: args[0].(*ast.MethodCallExpression), Type: ast.CallExpressionTypeMethodCall}
		}),
		newProduction(SymbolCallExpression, []Symbol{SymbolVarCallExpression}, 0, func(args []interface{}) ast.Node {
			return &ast.CallExpression{VarCallExpression: args[0].(*ast.VarCallExpression), Type: ast.CallExpressionTypeValCall}
		}),
	}

	pm.productionMap[SymbolExpression] = []*Production{
		newProduction(SymbolExpression, []Symbol{SymbolStringLiteral}, 0, func(args []interface{}) ast.Node {
			return &ast.Expression{StringLiteral: args[0].(*ast.StringLiteral), Type: ast.ExpressionTypeString}
		}),
		newProduction(SymbolExpression, []Symbol{SymbolIntLiteral}, 0, func(args []interface{}) ast.Node {
			return &ast.Expression{IntLiteral: args[0].(*ast.IntLiteral), Type: ast.ExpressionTypeString}
		}),
		newProduction(SymbolExpression, []Symbol{SymbolDoubleLiteral}, 0, func(args []interface{}) ast.Node {
			return &ast.Expression{DoubleLiteral: args[0].(*ast.DoubleLiteral), Type: ast.ExpressionTypeDouble}
		}),
		newProduction(SymbolExpression, []Symbol{SymbolNull}, 0, func(args []interface{}) ast.Node {
			return &ast.Expression{NullLiteral: args[0].(*ast.NullLiteral), Type: ast.ExpressionTypeNull}
		}),
		newProduction(SymbolExpression, []Symbol{SymbolTrue}, 0, func(args []interface{}) ast.Node {
			return &ast.Expression{BoolLiteral: args[0].(*ast.BoolLiteral), Type: ast.ExpressionTypeBool}
		}),
		newProduction(SymbolExpression, []Symbol{SymbolFalse}, 0, func(args []interface{}) ast.Node {
			return &ast.Expression{BoolLiteral: args[0].(*ast.BoolLiteral), Type: ast.ExpressionTypeBool}
		}),
		newProduction(SymbolExpression, []Symbol{SymbolNewObjExpression}, 0, func(args []interface{}) ast.Node {
			return &ast.Expression{NewObjectExpression: args[0].(*ast.NewObjectExpression), Type: ast.ExpressionTypeNewObject}
		}),
		newProduction(SymbolExpression, []Symbol{SymbolCallExpression}, 0, func(args []interface{}) ast.Node {
			return &ast.Expression{CallExpression: args[0].(*ast.CallExpression), Type: ast.ExpressionTypeCall}
		}),
	}

	pm.productionMap[SymbolTypeVar] = []*Production{
		newProduction(SymbolTypeVar, []Symbol{SymbolVoid, SymbolIdentifier}, 0, func(args []interface{}) ast.Node {
			return &ast.TypeVar{Type: "void", Name: args[1].(*lexer.Token).V}
		}),
		newProduction(SymbolTypeVar, []Symbol{SymbolIdentifier, SymbolIdentifier}, 0, func(args []interface{}) ast.Node {
			return &ast.TypeVar{Type: args[1].(*lexer.Token).V, Name: args[1].(*lexer.Token).V}
		}),
	}

	pm.productionMap[SymbolExpressionStatement] = []*Production{
		newProduction(SymbolExpressionStatement, []Symbol{SymbolExpression, SymbolSemicolon}, 0, func(args []interface{}) ast.Node {
			expr := args[0].(*ast.Expression)
			return &ast.ExpressionStatement{Expression: expr}
		}),
	}

	pm.productionMap[SymbolVarAssignStatement] = []*Production{
		newProduction(SymbolVarAssignStatement, []Symbol{SymbolTypeVar, SymbolAssign, SymbolExpressionStatement}, 0, func(args []interface{}) ast.Node {
			varToken := args[0].(*lexer.Token)
			exprStmt := args[2].(*ast.ExpressionStatement)
			return &ast.VarAssignStatement{Name: varToken.V, Expression: exprStmt.Expression, Type: ast.VarAssignStatementTypeVar}
		}),
		newProduction(SymbolVarAssignStatement, []Symbol{SymbolVarCallExpression, SymbolAssign, SymbolExpressionStatement}, 0, func(args []interface{}) ast.Node {
			varCallExpr := args[0].(*ast.VarCallExpression)
			exprStmt := args[2].(*ast.ExpressionStatement)
			return &ast.VarAssignStatement{VarCallExpression: varCallExpr, Expression: exprStmt.Expression, Type: ast.VarAssignStatementTypeVarCall}
		}),
	}

	pm.productionMap[SymbolVarDeclarationStatement] = []*Production{
		newProduction(SymbolVarDeclarationStatement, []Symbol{SymbolTypeVar, SymbolSemicolon}, 0, func(args []interface{}) ast.Node {
			typeVar := args[0].(*ast.TypeVar)
			return &ast.VarDeclarationStatement{Type: typeVar.Type, Name: typeVar.Name}
		}),
	}

	pm.productionMap[SymbolReturnStatement] = []*Production{
		newProduction(SymbolReturnStatement, []Symbol{SymbolReturn, SymbolSemicolon}, 0, func(args []interface{}) ast.Node {
			return &ast.ReturnStatement{}
		}),
		newProduction(SymbolReturnStatement, []Symbol{SymbolReturn, SymbolExpressionStatement}, 0, func(args []interface{}) ast.Node {
			exprStmt := args[0].(*ast.ExpressionStatement)
			return &ast.ReturnStatement{Expression: exprStmt.Expression}
		}),
	}

	pm.productionMap[SymbolContinueStatement] = []*Production{
		newProduction(SymbolContinueStatement, []Symbol{SymbolContinue, SymbolSemicolon}, 0, func(args []interface{}) ast.Node {
			return &ast.ContinueStatement{}
		}),
		newProduction(SymbolContinueStatement, []Symbol{SymbolContinue, SymbolExpressionStatement}, 0, func(args []interface{}) ast.Node {
			exprStmt := args[0].(*ast.ExpressionStatement)
			return &ast.ContinueStatement{Expression: exprStmt.Expression}
		}),
	}

	pm.productionMap[SymbolBreakStatement] = []*Production{
		newProduction(SymbolBreakStatement, []Symbol{SymbolBreak, SymbolSemicolon}, 0, func(args []interface{}) ast.Node {
			return &ast.BreakStatement{}
		}),
		newProduction(SymbolBreakStatement, []Symbol{SymbolBreak, SymbolExpressionStatement}, 0, func(args []interface{}) ast.Node {
			exprStmt := args[0].(*ast.ExpressionStatement)
			return &ast.BreakStatement{Expression: exprStmt.Expression}
		}),
	}

	pm.productionMap[SymbolForStatement] = []*Production{
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolLp, SymbolSemicolon, SymbolSemicolon, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			block := args[5].(*ast.Block)
			return &ast.ForStatement{Block: block}
		}),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolLp, SymbolExpression, SymbolSemicolon, SymbolSemicolon, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			initExpr := args[2].(*ast.Expression)
			block := args[6].(*ast.Block)
			return &ast.ForStatement{InitExpression: initExpr, Block: block}
		}),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolLp, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			initExpr := args[2].(*ast.Expression)
			condExpr := args[4].(*ast.Expression)
			block := args[7].(*ast.Block)
			return &ast.ForStatement{InitExpression: initExpr, CondExpression: condExpr, Block: block}
		}),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolLp, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			initExpr := args[2].(*ast.Expression)
			condExpr := args[4].(*ast.Expression)
			postExpr := args[6].(*ast.Expression)
			block := args[8].(*ast.Block)
			return &ast.ForStatement{InitExpression: initExpr, CondExpression: condExpr, PostExpression: postExpr, Block: block}
		}),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolLp, SymbolExpression, SymbolSemicolon, SymbolSemicolon, SymbolExpression, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			initExpr := args[2].(*ast.Expression)
			postExpr := args[5].(*ast.Expression)
			block := args[7].(*ast.Block)
			return &ast.ForStatement{InitExpression: initExpr, PostExpression: postExpr, Block: block}
		}),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolLp, SymbolSemicolon, SymbolSemicolon, SymbolExpression, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			postExpr := args[4].(*ast.Expression)
			block := args[6].(*ast.Block)
			return &ast.ForStatement{PostExpression: postExpr, Block: block}
		}),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolLp, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			condExpr := args[3].(*ast.Expression)
			postExpr := args[5].(*ast.Expression)
			block := args[7].(*ast.Block)
			return &ast.ForStatement{CondExpression: condExpr, PostExpression: postExpr, Block: block}
		}),
		newProduction(SymbolForStatement, []Symbol{SymbolFor, SymbolLp, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			condExpr := args[4].(*ast.Expression)
			block := args[8].(*ast.Block)
			return &ast.ForStatement{CondExpression: condExpr, Block: block}
		}),
	}

	pm.productionMap[SymbolIfStatement] = []*Production{
		newProduction(SymbolIfStatement, []Symbol{SymbolIf, SymbolLp, SymbolExpression, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			expr := args[2].(*ast.Expression)
			ifBlock := args[4].(*ast.Block)
			return &ast.IfStatement{CondExpression: expr, IfBlock: ifBlock}
		}),
		newProduction(SymbolIfStatement, []Symbol{SymbolIf, SymbolLp, SymbolExpression, SymbolRp, SymbolBlock, SymbolElse, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			expr := args[2].(*ast.Expression)
			ifBlock := args[4].(*ast.Block)
			elseBlock := args[6].(*ast.Block)
			return &ast.IfStatement{CondExpression: expr, IfBlock: ifBlock, ElseBlock: elseBlock}
		}),
	}

	pm.productionMap[SymbolWhileStatement] = []*Production{
		newProduction(SymbolWhileStatement, []Symbol{SymbolWhile, SymbolLp, SymbolExpression, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			expr := args[2].(*ast.Expression)
			block := args[4].(*ast.Block)
			return &ast.WhileStatement{Expression: expr, Block: block}
		}),
	}

	pm.productionMap[SymbolStatement] = []*Production{
		newProduction(SymbolStatement, []Symbol{SymbolExpressionStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.ExpressionStatement)
			return &ast.Statement{ExpressionStatement: stmt, Type: ast.StatementTypeExpression}
		}),
		newProduction(SymbolStatement, []Symbol{SymbolVarDeclarationStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.VarDeclarationStatement)
			return &ast.Statement{VarDeclarationStatement: stmt, Type: ast.StatementTypeVarDeclaration}
		}),
		newProduction(SymbolStatement, []Symbol{SymbolVarAssignStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.VarAssignStatement)
			return &ast.Statement{VarAssignStatement: stmt, Type: ast.StatementTypeVarAssign}
		}),
		newProduction(SymbolStatement, []Symbol{SymbolWhileStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.WhileStatement)
			return &ast.Statement{WhileStatement: stmt, Type: ast.StatementTypeWhile}
		}),
		newProduction(SymbolStatement, []Symbol{SymbolIfStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.IfStatement)
			return &ast.Statement{IfStatement: stmt, Type: ast.StatementTypeIf}
		}),
		newProduction(SymbolStatement, []Symbol{SymbolForStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.ForStatement)
			return &ast.Statement{ForStatement: stmt, Type: ast.StatementTypeFor}
		}),
		newProduction(SymbolStatement, []Symbol{SymbolBreakStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.BreakStatement)
			return &ast.Statement{BreakStatement: stmt, Type: ast.StatementTypeBreak}
		}),
		newProduction(SymbolStatement, []Symbol{SymbolContinueStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.ContinueStatement)
			return &ast.Statement{ContinueStatement: stmt, Type: ast.StatementTypeContinue}
		}),
		newProduction(SymbolStatement, []Symbol{SymbolReturnStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.ReturnStatement)
			return &ast.Statement{ReturnStatement: stmt, Type: ast.StatementTypeReturn}
		}),
	}

	pm.productionMap[SymbolStatementList] = []*Production{
		newProduction(SymbolStatementList, []Symbol{SymbolStatement}, 0, func(args []interface{}) ast.Node {
			stmt := args[0].(*ast.Statement)
			stmtList := new(ast.StatementList)
			stmtList.List = append(stmtList.List, stmt)
			return stmtList
		}),
		newProduction(SymbolStatementList, []Symbol{SymbolStatementList, SymbolStatement}, 0, func(args []interface{}) ast.Node {
			stmtList := args[0].(*ast.StatementList)
			stmt := args[1].(*ast.Statement)
			stmtList.List = append(stmtList.List, stmt)
			return stmtList
		}),
	}

	pm.productionMap[SymbolEmptyBlock] = []*Production{
		newProduction(SymbolEmptyBlock, []Symbol{SymbolLc, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return &ast.Block{}
		}),
	}

	pm.productionMap[SymbolBlock] = []*Production{
		newProduction(SymbolBlock, []Symbol{SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return &ast.Block{}
		}),
		newProduction(SymbolBlock, []Symbol{SymbolLc, SymbolStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			stmtList := args[1].(*ast.StatementList)
			return &ast.Block{StatementList: stmtList.List}
		}),
	}

	pm.productionMap[SymbolParameterList] = []*Production{
		newProduction(SymbolParameterList, []Symbol{SymbolTypeVar}, 0, func(args []interface{}) ast.Node {
			typeVar := args[0].(*ast.TypeVar)
			param := new(ast.Parameter)
			param.Name = typeVar.Name
			param.Type = typeVar.Type
			paramList := new(ast.ParameterList)
			paramList.List = append(paramList.List, param)
			return paramList
		}),
		newProduction(SymbolParameterList, []Symbol{SymbolParameterList, SymbolComma, SymbolTypeVar}, 0, func(args []interface{}) ast.Node {
			paramList := args[0].(*ast.ParameterList)
			typeVar := args[2].(*ast.TypeVar)
			param := new(ast.Parameter)
			param.Name = typeVar.Name
			param.Type = typeVar.Type
			paramList.List = append(paramList.List, param)
			return paramList
		}),
	}

	pm.productionMap[SymbolMemberModifier] = []*Production{
		newProduction(SymbolMemberModifier, []Symbol{SymbolPublic}, 0, func(args []interface{}) ast.Node {
			return &ast.MemberModifier{Type: ast.ModifierPublic}
		}),
		newProduction(SymbolMemberModifier, []Symbol{SymbolProtected}, 0, func(args []interface{}) ast.Node {
			return &ast.MemberModifier{Type: ast.ModifierProtected}
		}),
		newProduction(SymbolMemberModifier, []Symbol{SymbolPrivate}, 0, func(args []interface{}) ast.Node {
			return &ast.MemberModifier{Type: ast.ModifierPrivate}
		}),
		newProduction(SymbolMemberModifier, []Symbol{SymbolAbstract}, 0, func(args []interface{}) ast.Node {
			return &ast.MemberModifier{Type: ast.ModifierAbstract}
		}),
	}

	pm.productionMap[SymbolMethodDefinition] = []*Production{
		newProduction(SymbolMethodDefinition, []Symbol{SymbolMemberModifier, SymbolTypeVar, SymbolLp, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			modifier := args[0].(*ast.MemberModifier)
			typeVar := args[1].(*ast.TypeVar)
			block := args[4].(*ast.Block)
			return &ast.MethodDefinition{ModifierType: modifier.Type, Name: typeVar.Name, Type: typeVar.Type, Block: block}
		}),
		newProduction(SymbolMethodDefinition, []Symbol{SymbolMemberModifier, SymbolTypeVar, SymbolLp, SymbolParameterList, SymbolRp, SymbolBlock}, 0, func(args []interface{}) ast.Node {
			modifier := args[0].(*ast.MemberModifier)
			typeVar := args[1].(*ast.TypeVar)
			paramList := args[3].(*ast.ParameterList)
			block := args[5].(*ast.Block)
			return &ast.MethodDefinition{ModifierType: modifier.Type, Name: typeVar.Name, Type: typeVar.Type, ParameterList: paramList.List, Block: block}
		}),
	}

	pm.productionMap[SymbolPropertyDefinition] = []*Production{
		newProduction(SymbolPropertyDefinition, []Symbol{SymbolMemberModifier, SymbolTypeVar, SymbolSemicolon}, 0, func(args []interface{}) ast.Node {
			modifier := args[0].(*ast.MemberModifier)
			typeVar := args[1].(*ast.TypeVar)
			return &ast.PropertyDefinition{ModifierType: modifier.Type, Name: typeVar.Name, Type: typeVar.Type}
		}),
		newProduction(SymbolPropertyDefinition, []Symbol{SymbolMemberModifier, SymbolTypeVar, SymbolAssign, SymbolExpressionStatement}, 0, func(args []interface{}) ast.Node {
			modifier := args[0].(*ast.MemberModifier)
			typeVar := args[1].(*ast.TypeVar)
			exprStmt := args[1].(*ast.ExpressionStatement)
			return &ast.PropertyDefinition{ModifierType: modifier.Type, Name: typeVar.Name, Type: typeVar.Type, Expr: exprStmt.Expression}
		}),
	}

	pm.productionMap[SymbolClassStatement] = []*Production{
		newProduction(SymbolClassStatement, []Symbol{SymbolMethodDefinition}, 0, func(args []interface{}) ast.Node {
			md := args[0].(*ast.MethodDefinition)
			return &ast.ClassStatement{MethodDefinition: md, Type: ast.ClassStatementTypeMethod}
		}),
		newProduction(SymbolClassStatement, []Symbol{SymbolPropertyDefinition}, 0, func(args []interface{}) ast.Node {
			pd := args[0].(*ast.PropertyDefinition)
			return &ast.ClassStatement{PropertyDefinition: pd, Type: ast.ClassStatementTypeProperty}
		}),
	}

	pm.productionMap[SymbolClassStatementList] = []*Production{
		newProduction(SymbolClassStatementList, []Symbol{SymbolClassStatement}, 0, func(args []interface{}) ast.Node {
			cs := args[0].(*ast.ClassStatement)
			csl := new(ast.ClassStatementList)
			csl.List = append(csl.List, cs)
			return csl
		}),
		newProduction(SymbolClassStatementList, []Symbol{SymbolClassStatementList, SymbolClassStatement}, 0, func(args []interface{}) ast.Node {
			csl := args[0].(*ast.ClassStatementList)
			cs := args[1].(*ast.ClassStatement)
			csl.List = append(csl.List, cs)
			return csl
		}),
	}

	pm.productionMap[SymbolImplementsDeclaration] = []*Production{
		newProduction(SymbolImplementsDeclaration, []Symbol{SymbolImplements, SymbolIdentifier}, 0, func(args []interface{}) ast.Node {
			t := args[1].(*lexer.Token)
			id := new(ast.ImplementsDeclaration)
			id.InterfaceNameList = append(id.InterfaceNameList, t.V)
			return id
		}),
		newProduction(SymbolImplementsDeclaration, []Symbol{SymbolImplementsDeclaration, SymbolComma, SymbolIdentifier}, 0, func(args []interface{}) ast.Node {
			id := args[0].(*ast.ImplementsDeclaration)
			t := args[2].(*lexer.Token)
			id.InterfaceNameList = append(id.InterfaceNameList, t.V)
			return id
		}),
	}

	pm.productionMap[SymbolExtendsDelcaration] = []*Production{
		newProduction(SymbolExtendsDelcaration, []Symbol{SymbolExtends, SymbolIdentifier}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolExtendsDelcaration, []Symbol{SymbolExtends, SymbolComma, SymbolIdentifier}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
	}

	pm.productionMap[SymbolClassDeclaration] = []*Production{
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolLc, SymbolClassStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolLc, SymbolClassStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassDeclaration, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
	}

	pm.productionMap[SymbolInterfaceMethodDeclarationStatement] = []*Production{
		newProduction(SymbolInterfaceMethodDeclarationStatement, []Symbol{SymbolTypeVar, SymbolLp, SymbolRp, SymbolSemicolon}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolInterfaceMethodDeclarationStatement, []Symbol{SymbolTypeVar, SymbolLp, SymbolParameterList, SymbolRp, SymbolSemicolon}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
	}

	pm.productionMap[SymbolInterfaceMethodDeclarationStatementList] = []*Production{
		newProduction(SymbolInterfaceMethodDeclarationStatementList, []Symbol{SymbolInterfaceMethodDeclarationStatement}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolInterfaceMethodDeclarationStatementList, []Symbol{SymbolInterfaceMethodDeclarationStatementList, SymbolInterfaceMethodDeclarationStatement}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
	}

	pm.productionMap[SymbolInterfaceDeclaration] = []*Production{
		newProduction(SymbolInterfaceDeclaration, []Symbol{SymbolInterface, SymbolIdentifier, SymbolEmptyBlock}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolInterfaceDeclaration, []Symbol{SymbolInterface, SymbolIdentifier, SymbolLc, SymbolInterfaceMethodDeclarationStatementList, SymbolRc}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
	}

	pm.productionMap[SymbolClassInterfaceDeclaration] = []*Production{
		newProduction(SymbolClassInterfaceDeclaration, []Symbol{SymbolClassDeclaration}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassInterfaceDeclaration, []Symbol{SymbolInterfaceDeclaration}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
	}

	pm.productionMap[SymbolClassInterfaceDeclarationList] = []*Production{
		newProduction(SymbolClassInterfaceDeclarationList, []Symbol{SymbolClassInterfaceDeclaration}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
		newProduction(SymbolClassInterfaceDeclarationList, []Symbol{SymbolClassInterfaceDeclarationList, SymbolClassInterfaceDeclaration}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
	}

	pm.productionMap[SymbolTranslationUnit] = []*Production{
		newProduction(SymbolTranslationUnit, []Symbol{SymbolClassInterfaceDeclarationList}, 0, func(args []interface{}) ast.Node {
			return nil
		}),
	}

	productionManager = pm

	return
}

// 获取以左侧为left的生成式列表
func (pm *ProductionManager) getProductions(left Symbol) (productions []*Production) {
	return pm.productionMap[left]
}

func (pm *ProductionManager) getFirstSetBuilder() *FirstSetBuilder {
	return pm.firstSetBuilder
}
