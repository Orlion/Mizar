package parser

type TranslationUnit struct {
	DefinitionOrStatementList []*DefinitionOrStatement
}

type DefinitionOrStatement struct {
	T                  string // FunctionDefinition, Statement
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

type Statement struct {
	T                 string // Expression,	WhileStatement, IfStatement, BreakStatement, ContinueStatement, ReturnStatement
	WhileStatement    *WhileStatement
	IfStatement       *IfStatement
	BreakStatement    *BreakStatement
	ContinueStatement *ContinueStatement
	ReturnStatement   *ReturnStatement
}

type MultiplicativeExpression struct {
	T                        string // "", "Mul", "Div"
	PrimaryExpression        *PrimaryExpression
	MultiplicativeExpression *MultiplicativeExpression
}

type PrimaryExpression struct {
	T                  string // String, Number, Identifier, Expression FuncCallExpression
	String             string
	Number             string
	Identifier         string
	Expression         *Expression
	FuncCallExpression *FuncCallExpression
}

type Expression struct {
	T                  string // FuncCallExpression, AdditiveExpression, Identifier
	AdditiveExpression *AdditiveExpression
	Identifier         string
	FuncCallExpression *FuncCallExpression
}

type AdditiveExpression struct {
	T                        string // "" "Add", "Sub"
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
