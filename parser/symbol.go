package parser

import "unicode"

type Symbol string

const (
	NilSymbol                                     Symbol = ""
	SymbolStringLiteral                           Symbol = "STRING_LITERAL"
	SymbolIntLiteral                              Symbol = "INT_LITERAL"
	SymbolDoubleLiteral                           Symbol = "DOUBLE_LITERAL"
	SymbolNull                                    Symbol = "NULL"
	SymbolTrue                                    Symbol = "TRUE"
	SymbolFalse                                   Symbol = "FALSE"
	SymbolIdentifier                              Symbol = "IDENTIFIER"
	SymbolNew                                     Symbol = "NEW"
	SymbolLp                                      Symbol = "LP"
	SymbolRp                                      Symbol = "RP"
	SymbolDot                                     Symbol = "DOT"
	SymbolLc                                      Symbol = "LC"
	SymbolRc                                      Symbol = "RC"
	SymbolComma                                   Symbol = "COMMA"
	SymbolReturn                                  Symbol = "RETURN"
	SymbolContinue                                Symbol = "CONTINUE"
	SymbolBreak                                   Symbol = "BREAK"
	SymbolSemicolon                               Symbol = "SEMICOLON"
	SymbolFor                                     Symbol = "FOR"
	SymbolIf                                      Symbol = "IF"
	SymbolElse                                    Symbol = "ELSE"
	SymbolWhile                                   Symbol = "WHILE"
	SymbolAssign                                  Symbol = "ASSIGN"
	SymbolClass                                   Symbol = "CLASS"
	SymbolThis                                    Symbol = "THIS"
	SymbolInterface                               Symbol = "INTERFACE"
	SymbolAbstract                                Symbol = "ASBTRACT"
	SymbolImplements                              Symbol = "IMPLEMENTS"
	SymbolExtends                                 Symbol = "EXTENDS"
	SymbolVoid                                    Symbol = "VOID"
	SymbolPublic                                  Symbol = "PUBLIC"
	SymbolPrivate                                 Symbol = "PRIVATE"
	SymbolProtected                               Symbol = "PROTECTED"
	SymbolExpression                              Symbol = "expression"
	SymbolArgumentList                            Symbol = "argument_list"
	SymbolMethodCallExpression                    Symbol = "method_call_expression"
	SymbolVarCallExpression                       Symbol = "var_call_expression"
	SymbolCallExpression                          Symbol = "calSymboll_expression"
	SymbolNewObjExpression                        Symbol = "new_obj_expression"
	SymbolVarDeclaration                          Symbol = "var_declaration"
	SymbolVarAssignStatement                      Symbol = "var_assign_statement"
	SymbolVarDeclarationStatement                 Symbol = "var_declaration_statement"
	SymbolReturnStatement                         Symbol = "return_statement"
	SymbolContinueStatement                       Symbol = "continue_statement"
	SymbolBreakStatement                          Symbol = "break_statement"
	SymbolForStatement                            Symbol = "for_statement"
	SymbolIfStatement                             Symbol = "if_statement"
	SymbolWhileStatement                          Symbol = "while_statement"
	SymbolStatement                               Symbol = "statement"
	SymbolStatementList                           Symbol = "statement_list"
	SymbolBlock                                   Symbol = "block"
	SymbolMethodModifier                          Symbol = "method_modifier"
	SymbolParameter                               Symbol = "parameter"
	SymbolParameterList                           Symbol = "parameter_list"
	SymbolReturnValType                           Symbol = "return_val_type"
	SymbolMethodDefinition                        Symbol = "method_definition"
	SymbolVarModifier                             Symbol = "var_modifier"
	SymbolClassStatement                          Symbol = "class_statement"
	SymbolClassStatementList                      Symbol = "class_statement_list"
	SymbolClassDeclaration                        Symbol = "class_declaration"
	SymbolImplementsDeclaration                   Symbol = "implements_declaration"
	SymbolExtendsDelcaration                      Symbol = "extends_declaration"
	SymbolInterfaceMethodDeclarationStatement     Symbol = "interface_method_declaration_statement"
	SymbolInterfaceMethodDeclarationStatementList Symbol = "interface_method_declaration_statement_list"
	SymbolInterfaceDeclaration                    Symbol = "interface_declaration"
	SymbolClassInterfaceDeclaration               Symbol = "class_interface_declaration"
	SymbolClassInterfaceDeclarationList           Symbol = "class_interface_declaration_list"
	SymbolTranslationUnit                         Symbol = "translation_unit"
)

// 是否是终结符
func (s Symbol) isTerminals() bool {
	return unicode.IsUpper([]rune(s)[0])
}

type Symbols struct {
	value       Symbol
	productions [][]Symbol
	firstSet    map[Symbol]struct{}
	isNullable  bool
}

func newSymbols(symbol Symbol, nullable bool, productions [][]Symbol) *Symbols {
	symbols := new(Symbols)
	symbols.value = symbol
	symbols.isNullable = nullable
	symbols.productions = productions
	symbols.firstSet = make(map[Symbol]struct{})

	if symbol.isTerminals() {
		// 终结符的first set是它自己
		symbols.firstSet[symbol] = struct{}{}
	}

	return symbols
}
