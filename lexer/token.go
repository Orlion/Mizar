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
	TokenAssign TokenType = "ASSIGN"
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
	TokenLc = "LC"
	TokenRc = "RC"
	TokenLp = "LP"
	TokenRp = "RP"
	// TokenLb            = "[" // 暂时不实现泛型
	// TokenRb            = "]"
	TokenSemicolon = "SEMICOLON"
	TokenComma     = "COMMA"
	// TokenExclamation   = "!"
	TokenDot           = "DOT"
	TokenContinue      = "CONTINUE"
	TokenReturn        = "RETURN"
	TokenWhile         = "WHILE"
	TokenBreak         = "BREAK"
	TokenElse          = "ELSE"
	TokenVoid          = "VOID"
	TokenIf            = "IF"
	TokenFor           = "FOR"
	TokenClass         = "CLASS"
	TokenInterface     = "INTERFACE"
	TokenAbstract      = "ABSTRACT"
	TokenPublic        = "PUBLIC"
	TokenPrivate       = "PRIVATE"
	TokenProtected     = "PROTECTED"
	TokenImplements    = "IMPLEMENTS"
	TokenExtends       = "EXTENDS"
	TokenNew           = "NEW"
	TokenTrue          = "TRUE"
	TokenFalse         = "FALSE"
	TokenNull          = "NULL"
	TokenThis          = "THIS"
	TokenIdentifier    = "IDENTIFIER"
	TokenStringLiteral = "STRING_LITERAL"
	TokenDoubleLiteral = "DOUBLE_LITERAL"
	TokenIntLiteral    = "INT_LITERAL"
)

var reservedWords = []string{
	"=", "{", "}", "(", ")", "[", "]", ";", ",", ".",
	"continue", "return", "while", "break", "else", "void", "if", "for", "class", "interface", "abstract", "public",
	"private", "protected", "implements", "extends", "true", "false", "null", "this", "new",
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
	"this":       TokenThis,
	"new":        TokenNew,
}

func init() {
	sort.Sort(utils.SortByLength(reservedWords))
}
