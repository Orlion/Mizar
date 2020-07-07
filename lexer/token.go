package lexer

type Token struct {
	T TokenType
	V string
}

type TokenType string

const (
	TokenTypeAssign     TokenType = "="
	TokenTypeAdd                  = "+"
	TokenTypeSub                  = "-"
	TokenTypeMul                  = "*"
	TokenTypeDiv                  = "/"
	TokenTypeLc                   = "{"
	TokenTypeRc                   = "}"
	TokenTypeLp                   = "("
	TokenTypeRp                   = ")"
	TokenTypeSemicolon            = ";"
	TokenTypeComma                = ","
	TokenTypeNumber               = "number"
	TokenTypeString               = "string"
	TokenTypeIdentifier           = "identifier"
	TokenTypeContinue             = "continue"
	TokenTypeElseIf               = "elseif"
	TokenTypeReturn               = "return"
	TokenTypeWhile                = "while"
	TokenTypeBreak                = "break"
	TokenTypeElse                 = "else"
	TokenTypeFunc                 = "func"
	TokenTypeIf                   = "if"
)

// 按照长度从大到小排列，token识别先尝试识别最长的keyword
var Keywords = []TokenType{TokenTypeContinue, TokenTypeElseIf, TokenTypeReturn, TokenTypeWhile, TokenTypeBreak, TokenTypeElse, TokenTypeFunc, TokenTypeIf}
