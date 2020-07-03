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
	TokenTypeLc                   = "{"
	TokenTypeRc                   = "}"
	TokenTypeLp                   = "("
	TokenTypeRp                   = ")"
	TokenTypeSemicolon            = ";"
	TokenTypeUint                 = "uint"
	TokenTypeString               = "string"
	TokenTypeIdentifier           = "identifier"
	TokenTypeContinue             = "continue"
	TokenTypeElseIf               = "elseif"
	TokenTypeWhile                = "while"
	TokenTypeBreak                = "break"
	TokenTypeElse                 = "else"
	TokenTypeIf                   = "if"
)

// 按照长度从大到小排列，token识别先尝试识别最长的keyword
var Keywords = [6]TokenType{TokenTypeContinue, TokenTypeElseIf, TokenTypeWhile, TokenTypeBreak, TokenTypeElse, TokenTypeIf}
