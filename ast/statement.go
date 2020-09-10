package ast

type Block struct {
	StatementList []*Statement
}

func (block *Block) accept(vistor ASTVistor) {

}

type StatementList struct {
	List []*Statement
}

func (stmtList *StatementList) accept(vistor ASTVistor) {

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

func (stmt *Statement) accept(vistor ASTVistor) {

}

type ExpressionStatement struct {
	Expression *Expression
}

func (exprStmt *ExpressionStatement) accept(vistor ASTVistor) {

}

type VarDeclarationStatement struct {
	Type string
	Name string
}

func (varDeclStmt *VarDeclarationStatement) accept(vistor ASTVistor) {

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

func (varAssignStmt *VarAssignStatement) accept(vistor ASTVistor) {

}

type WhileStatement struct {
	Expression *Expression
	Block      *Block
}

func (whileStmt *WhileStatement) accept(vistor ASTVistor) {

}

type IfStatement struct {
	CondExpression *Expression
	IfBlock        *Block
	ElseBlock      *Block
}

func (ifStmt *IfStatement) accept(vistor ASTVistor) {

}

type ForStatement struct {
	InitExpression *Expression
	CondExpression *Expression
	PostExpression *Expression
	Block          *Block
}

func (forStmt *ForStatement) accept(vistor ASTVistor) {

}

type BreakStatement struct {
	Expression *Expression
}

func (breakStmt *BreakStatement) accept(vistor ASTVistor) {

}

type ContinueStatement struct {
	Expression *Expression
}

func (continueStmt *ContinueStatement) accept(vistor ASTVistor) {

}

type ReturnStatement struct {
	Expression *Expression
}

func (returnStmt *ReturnStatement) accept(vistor ASTVistor) {

}
