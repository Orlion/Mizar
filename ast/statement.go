package ast

type Block struct {
	StatementList []*Statement
}

func (block *Block) Accept(visitor Visitor) {

}

type StatementList struct {
	List []*Statement
}

func (stmtList *StatementList) Accept(visitor Visitor) {

}

type StatementType int8

const (
	StatementTypeExpression StatementType = iota + 1
	StatementTypeVarDeclaration
	StatementTypeVarAssign
	StatementTypeWhile
	StatementTypeIf
	StatementTypeFor
	StatementTypeBreak
	StatementTypeContinue
	StatementTypeReturn
)

type Statement struct {
	ExpressionStatement     *ExpressionStatement
	VarDeclarationStatement *VarDeclarationStatement
	VarAssignStatement      *VarAssignStatement
	WhileStatement          *WhileStatement
	IfStatement             *IfStatement
	ForStatement            *ForStatement
	BreakStatement          *BreakStatement
	ContinueStatement       *ContinueStatement
	ReturnStatement         *ReturnStatement
	Type                    StatementType
}

func (stmt *Statement) Accept(visitor Visitor) {

}

type ExpressionStatement struct {
	Expression *Expression
}

func (exprStmt *ExpressionStatement) Accept(visitor Visitor) {

}

type VarDeclarationStatement struct {
	Type string
	Name string
}

func (varDeclStmt *VarDeclarationStatement) Accept(visitor Visitor) {

}

type VarAssignStatementType int8

const (
	VarAssignStatementTypeVar VarAssignStatementType = iota + 1
	VarAssignStatementTypeVarCall
)

type VarAssignStatement struct {
	VarType           string
	VarName           string
	VarCallExpression *VarCallExpression
	Expression        *Expression
	Type              VarAssignStatementType
}

func (varAssignStmt *VarAssignStatement) Accept(visitor Visitor) {

}

type WhileStatement struct {
	Expression *Expression
	Block      *Block
}

func (whileStmt *WhileStatement) Accept(visitor Visitor) {

}

type IfStatement struct {
	CondExpression *Expression
	IfBlock        *Block
	ElseBlock      *Block
}

func (ifStmt *IfStatement) Accept(visitor Visitor) {

}

type ForStatement struct {
	InitExpression *Expression
	CondExpression *Expression
	PostExpression *Expression
	Block          *Block
}

func (forStmt *ForStatement) Accept(visitor Visitor) {

}

type BreakStatement struct {
	Expression *Expression
}

func (breakStmt *BreakStatement) Accept(visitor Visitor) {

}

type ContinueStatement struct {
	Expression *Expression
}

func (continueStmt *ContinueStatement) Accept(visitor Visitor) {

}

type ReturnStatement struct {
	Expression *Expression
}

func (returnStmt *ReturnStatement) Accept(visitor Visitor) {

}
