package lexer

import (
	"mizar/utils"
	"sort"

	"github.com/Orlion/merak/symbol"
)

type TokenType string

const (
	EoiToken           TokenType = "EOI"
	TokenAssign                  = "ASSIGN"
	TokenLc                      = "LC"
	TokenRc                      = "RC"
	TokenLp                      = "LP"
	TokenRp                      = "RP"
	TokenSemicolon               = "SEMICOLON"
	TokenComma                   = "COMMA"
	TokenDot                     = "DOT"
	TokenContinue                = "CONTINUE"
	TokenReturn                  = "RETURN"
	TokenWhile                   = "WHILE"
	TokenBreak                   = "BREAK"
	TokenElse                    = "ELSE"
	TokenVoid                    = "VOID"
	TokenIf                      = "IF"
	TokenFor                     = "FOR"
	TokenClass                   = "CLASS"
	TokenInterface               = "INTERFACE"
	TokenAbstract                = "ABSTRACT"
	TokenPublic                  = "PUBLIC"
	TokenPrivate                 = "PRIVATE"
	TokenProtected               = "PROTECTED"
	TokenImplements              = "IMPLEMENTS"
	TokenExtends                 = "EXTENDS"
	TokenNew                     = "NEW"
	TokenTrue                    = "TRUE"
	TokenFalse                   = "FALSE"
	TokenNull                    = "NULL"
	TokenThis                    = "THIS"
	TokenIdentifier              = "IDENTIFIER"
	TokenStringLiteral           = "STRING_LITERAL"
	TokenDoubleLiteral           = "DOUBLE_LITERAL"
	TokenIntLiteral              = "INT_LITERAL"
)

type Token struct {
	T           TokenType
	Lexeme      string
	StartLine   int
	StartColumn int
	EndLine     int
	EndColumn   int
	FileName    string // token所在文件名
}

func (t *Token) ToSymbol() symbol.Symbol {
	return Symbol(t.T)
}

func (t *Token) ToString() string {
	return t.Lexeme
}

var reservedWords = []string{
	"=", "{", "}", "(", ")", ";", ",", ".",
	"continue", "return", "while", "break", "else", "void", "if", "for", "class", "interface", "abstract", "public",
	"private", "protected", "implements", "extends", "true", "false", "null", "this", "new",
}

var reservedWords2TokenTypeMap = map[string]TokenType{
	"=":          TokenAssign,
	"{":          TokenLc,
	"}":          TokenRc,
	"(":          TokenLp,
	")":          TokenRp,
	";":          TokenSemicolon,
	",":          TokenComma,
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
