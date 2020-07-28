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
	TokenAssign        TokenType = "="
	TokenAddAssign               = "+="
	TokenSubAssign               = "-="
	TokenMulAssign               = "*="
	TokenDivAssign               = "/="
	TokenModAssign               = "%="
	TokenAdd                     = "+"
	TokenSub                     = "-"
	TokenMul                     = "*"
	TokenDiv                     = "/"
	TokenMod                     = "%"
	TokenIncrement               = "++"
	TokenDecrement               = "--"
	TokenLogicalAdd              = "&&"
	TokenLogicalOr               = "||"
	TokenEq                      = "=="
	TokenNe                      = "!="
	TokenGt                      = ">"
	TokenGe                      = ">="
	TokenLt                      = "<"
	TokenLe                      = "<="
	TokenLc                      = "{"
	TokenRc                      = "}"
	TokenLp                      = "("
	TokenRp                      = ")"
	TokenLb                      = "["
	TokenRb                      = "]"
	TokenSemicolon               = ";"
	TokenComma                   = ","
	TokenExclamation             = "!"
	TokenDot                     = "."
	TokenInt                     = "int"
	TokenString                  = "string"
	TokenBool                    = "bool"
	TokenDouble                  = "double"
	TokenChar                    = "char"
	TokenContinue                = "continue"
	TokenReturn                  = "return"
	TokenWhile                   = "while"
	TokenBreak                   = "break"
	TokenElse                    = "else"
	TokenFunction                = "function"
	TokenIf                      = "if"
	TokenFor                     = "for"
	TokenClass                   = "class"
	TokenInterface               = "interface"
	TokenAbstract                = "abstract"
	TokenPublic                  = "public"
	TokenPrivate                 = "private"
	TokenProtected               = "protected"
	TokenTrue                    = "true"
	TokenFalse                   = "false"
	TokenNull                    = "null"
	TokenTry                     = "try"
	TokenCatch                   = "catch"
	TokenFinally                 = "finally"
	TokenThrow                   = "throw"
	TokenThrows                  = "throws"
	TokenIdentifier              = "identifier"
	TokenStringLiteral           = "stringLiteral"
	TokenDoubleLiteral           = "doubleLiteral"
	TokenCharLiteral             = "charLiteral"
	TokenIntLiteral              = "intLiteral"
)

var reservedWords = []string{
	"=", "+=", "-=", "*=", "/=", "%=", "+", "-", "*", "/", "%", "++", "--", "&&", "||", "==", "!=", ">", ">=", "<", "<=", "{", "}", "(", ")", "[", "]", ";", ",", "!", ".", "int",
	"string", "bool", "double", "char", "continue", "elseif", "return", "while", "break", "else", "function", "if", "for", "class", "interface", "abstract", "public",
	"private", "protected", "true", "false", "null", "try", "catch", "finally", "throw", "throws",
}

var reservedWords2TokenTypeMap = map[string]TokenType{
	"=":         TokenAssign,
	"+=":        TokenAddAssign,
	"-=":        TokenSubAssign,
	"*=":        TokenMulAssign,
	"/=":        TokenDivAssign,
	"%=":        TokenModAssign,
	"+":         TokenAdd,
	"-":         TokenSub,
	"*":         TokenMul,
	"/":         TokenDiv,
	"%":         TokenMod,
	"++":        TokenIncrement,
	"--":        TokenDecrement,
	"&&":        TokenLogicalAdd,
	"||":        TokenLogicalOr,
	"==":        TokenEq,
	"!=":        TokenNe,
	">":         TokenGt,
	">=":        TokenGe,
	"<":         TokenLt,
	"<=":        TokenLe,
	"{":         TokenLc,
	"}":         TokenRc,
	"(":         TokenLp,
	")":         TokenRp,
	"[":         TokenLb,
	"]":         TokenRb,
	";":         TokenSemicolon,
	",":         TokenComma,
	"!":         TokenExclamation,
	".":         TokenDot,
	"int":       TokenInt,
	"string":    TokenString,
	"bool":      TokenBool,
	"double":    TokenDouble,
	"char":      TokenChar,
	"continue":  TokenContinue,
	"return":    TokenReturn,
	"while":     TokenWhile,
	"break":     TokenBreak,
	"else":      TokenElse,
	"function":  TokenFunction,
	"if":        TokenIf,
	"for":       TokenFor,
	"class":     TokenClass,
	"interface": TokenInterface,
	"abstract":  TokenAbstract,
	"public":    TokenPublic,
	"private":   TokenPrivate,
	"protected": TokenProtected,
	"true":      TokenTrue,
	"false":     TokenFalse,
	"null":      TokenNull,
	"try":       TokenTry,
	"catch":     TokenCatch,
	"finally":   TokenFinally,
	"throw":     TokenThrow,
	"throws":    TokenThrows,
}

func init() {
	sort.Sort(utils.SortByLength(reservedWords))
}
