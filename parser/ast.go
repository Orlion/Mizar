package parser

type TranslationUnit struct {
	DefinitionOrStatementList []*DefinitionOrStatement
}

type DefinitionOrStatementType string

const (
	DefinitionOrStatementTypeFunctionDefinition DefinitionOrStatementType = "FunctionDefinition"
	DefinitionOrStatementTypeStatement          DefinitionOrStatementType = "Statement"
)

type DefinitionOrStatement struct {
	T                  DefinitionOrStatementType // FunctionDefinition, Statement
	FunctionDefinition *FunctionDefinition
	Statement          *Statement
}

type FunctionDefinition struct {
	Name          string
	ParameterList *ParameterList
	Block         *Block
}

type ParameterList struct {
	ParameterList []*Parameter
}

type Parameter struct {
	Name string
}

type StatementType string

const (
	StatementTypeExpression        StatementType = "Expression"
	StatementTypeWhileStatement    StatementType = "WhileStatement"
	StatementTypeIfStatement       StatementType = "IfStatement"
	StatementTypeBreakStatement    StatementType = "BreakStatement"
	StatementTypeContinueStatement StatementType = "ContinueStatement"
	StatementTypeReturnStatement   StatementType = "ReturnStatement"
)

type Statement struct {
	T                 StatementType // Expression,	WhileStatement, IfStatement, BreakStatement, ContinueStatement, ReturnStatement
	Expression        *Expression
	WhileStatement    *WhileStatement
	IfStatement       *IfStatement
	BreakStatement    *BreakStatement
	ContinueStatement *ContinueStatement
	ReturnStatement   *ReturnStatement
}

type MultiplicativeExpressionType string

const (
	MultiplicativeExpressionTypeNull MultiplicativeExpressionType = ""
	MultiplicativeExpressionTypeMul  MultiplicativeExpressionType = "Mul"
	MultiplicativeExpressionTypeDiv  MultiplicativeExpressionType = "Div"
)

type MultiplicativeExpression struct {
	T                        MultiplicativeExpressionType // "", "Mul", "Div"
	PrimaryExpression        *PrimaryExpression
	MultiplicativeExpression *MultiplicativeExpression
}

type PrimaryExpressionType string

const (
	PrimaryExpressionTypeString             PrimaryExpressionType = "String"
	PrimaryExpressionTypeNumber             PrimaryExpressionType = "Number"
	PrimaryExpressionTypeIdentifier         PrimaryExpressionType = "Identifier"
	PrimaryExpressionTypeExpression         PrimaryExpressionType = "Expression"
	PrimaryExpressionTypeFuncCallExpression PrimaryExpressionType = "FuncCallExpression"
)

type PrimaryExpression struct {
	T                  PrimaryExpressionType // String, Number, Identifier, Expression FuncCallExpression
	String             string
	Number             string
	Identifier         string
	Expression         *Expression
	FuncCallExpression *FuncCallExpression
}

type ExpressionType string

const (
	ExpressionTypeAdditiveExpression ExpressionType = "AdditiveExpression"
	ExpressionTypeAssignment         ExpressionType = "Assignment"
)

type Expression struct {
	T                  ExpressionType // AdditiveExpression, Assignment
	AdditiveExpression *AdditiveExpression
	Assignment         *Assignment
	Expression         *Expression
}

type Assignment struct {
	Identifier string
}

type AdditiveExpressionType string

const (
	AdditiveExpressionTypeNull AdditiveExpressionType = ""
	AdditiveExpressionTypeAdd  AdditiveExpressionType = "Add"
	AdditiveExpressionTypeSub  AdditiveExpressionType = "Sub"
)

type AdditiveExpression struct {
	T                        AdditiveExpressionType // "" "Add", "Sub"
	MultiplicativeExpression *MultiplicativeExpression
	AdditiveExpression       *AdditiveExpression
}

type FuncCallExpression struct {
	Identifier   string
	ArgumentList *ArgumentList
}

type ArgumentList struct {
	ExpressionList []*Expression
}

type StatementList struct {
	StatementList []*Statement
}

type Block struct {
	StatementList *StatementList
}

type WhileStatement struct {
	Expression *Expression
	Block      *Block
}

type IfStatement struct {
	Expression *Expression
	Block      *Block
	ElseIfList *ElseIfList
	ElseBlock  *Block
}

type ElseIfList struct {
	ElseIfList []*ElseIf
}

type ElseIf struct {
	Expression *Expression
	Block      *Block
}

type BreakStatement struct {
}

type ContinueStatement struct {
}

type ReturnStatement struct {
	Expression *Expression
}
