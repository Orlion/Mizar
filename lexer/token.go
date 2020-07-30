package lexer

import (
	"mizar/utils"
	"sort"
)

type Token struct {
	T         TokenType
	V         string
	LineNum   int    // token所在行数
	ColumnNum int    // token所在列数
	FileName  string // token所在文件名
}

type TokenType string

const (
	TokenAssign TokenType = "="
	// TokenAddAssign               = "+=" // 运算完全由方法来实现，纯粹的面向对象
	// TokenSubAssign               = "-="
	// TokenMulAssign               = "*="
	// TokenDivAssign               = "/="
	// TokenModAssign               = "%="
	// TokenAdd                     = "+"
	// TokenSub                     = "-"
	// TokenMul                     = "*"
	// TokenDiv                     = "/"
	// TokenMod                     = "%"
	// TokenIncrement               = "++"
	// TokenDecrement               = "--"
	// TokenLogicalAdd              = "&&"
	// TokenLogicalOr               = "||"
	// TokenEq                      = "=="
	// TokenNe                      = "!="
	// TokenGt                      = ">"
	// TokenGe                      = ">="
	// TokenLt                      = "<"
	// TokenLe                      = "<="
	TokenLc = "{"
	TokenRc = "}"
	TokenLp = "("
	TokenRp = ")"
	// TokenLb            = "[" // 暂时不实现泛型
	// TokenRb            = "]"
	TokenSemicolon = ";"
	TokenComma     = ","
	// TokenExclamation   = "!"
	TokenDot           = "."
	TokenDoubleColon   = "::"
	TokenContinue      = "continue"
	TokenReturn        = "return"
	TokenWhile         = "while"
	TokenBreak         = "break"
	TokenElse          = "else"
	TokenVoid          = "void"
	TokenIf            = "if"
	TokenFor           = "for"
	TokenClass         = "class"
	TokenInterface     = "interface"
	TokenAbstract      = "abstract"
	TokenPublic        = "public"
	TokenPrivate       = "private"
	TokenProtected     = "protected"
	TokenImplements    = "implements"
	TokenExtends       = "extends"
	TokenTrue          = "true"
	TokenFalse         = "false"
	TokenNull          = "null"
	TokenConst         = "const"
	TokenThis          = "this"
	TokenSelf          = "self"
	TokenIdentifier    = "identifier"
	TokenStringLiteral = "stringLiteral"
	TokenDoubleLiteral = "doubleLiteral"
	TokenIntLiteral    = "intLiteral"
)

var reservedWords = []string{
	"=", "{", "}", "(", ")", "[", "]", ";", ",", ".", "::",
	"continue", "return", "while", "break", "else", "void", "if", "for", "class", "interface", "abstract", "public",
	"private", "protected", "implements", "extends", "true", "false", "null", "const", "this", "self",
}

var reservedWords2TokenTypeMap = map[string]TokenType{
	"=": TokenAssign,
	// "+=":        TokenAddAssign,
	// "-=":        TokenSubAssign,
	// "*=":        TokenMulAssign,
	// "/=":        TokenDivAssign,
	// "%=":        TokenModAssign,
	// "+":         TokenAdd,
	// "-":         TokenSub,
	// "*":         TokenMul,
	// "/":         TokenDiv,
	// "%":         TokenMod,
	// "++":        TokenIncrement,
	// "--":        TokenDecrement,
	// "&&":        TokenLogicalAdd,
	// "||":        TokenLogicalOr,
	// "==":        TokenEq,
	// "!=":        TokenNe,
	// ">":         TokenGt,
	// ">=":        TokenGe,
	// "<":         TokenLt,
	// "<=":        TokenLe,
	"{": TokenLc,
	"}": TokenRc,
	"(": TokenLp,
	")": TokenRp,
	// "[":         TokenLb,
	// "]":         TokenRb,
	";": TokenSemicolon,
	",": TokenComma,
	// "!":         TokenExclamation,
	".":          TokenDot,
	"continue":   TokenContinue,
	"return":     TokenReturn,
	"while":      TokenWhile,
	"break":      TokenBreak,
	"else":       TokenElse,
	"void":       TokenVoid,
	"if":         TokenIf,
	"for":        TokenFor,
	"class":      TokenClass,
	"interface":  TokenInterface,
	"abstract":   TokenAbstract,
	"public":     TokenPublic,
	"private":    TokenPrivate,
	"protected":  TokenProtected,
	"implements": TokenImplements,
	"extends":    TokenExtends,
	"true":       TokenTrue,
	"false":      TokenFalse,
	"null":       TokenNull,
	"const":      TokenConst,
	"::":         TokenDoubleColon,
	"this":       TokenThis,
	"self":       TokenSelf,
}

func init() {
	sort.Sort(utils.SortByLength(reservedWords))
}
