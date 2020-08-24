package parser

// 用来计算每个表达式的FirstSet
type FirstSetBuilder struct {
	symbolMap       map[Symbol]*Symbols
	symbolArr       []*Symbols
	runFirstSetPass bool
}

func newFirstSetBuilder() *FirstSetBuilder {
	firstSetBuilder := new(FirstSetBuilder)
	firstSetBuilder.runFirstSetPass = true
	firstSetBuilder.initProductions()
	return firstSetBuilder
}

func (fsb *FirstSetBuilder) getFirstSet(s Symbol) map[Symbol]struct{} {
	for _, symbols := range fsb.symbolArr {
		if symbols.value == s {
			return symbols.firstSet
		}
	}

	return nil
}

func (fsb *FirstSetBuilder) isSymbolNullable(symbol Symbol) bool {
	if symbols, exists := fsb.symbolMap[symbol]; !exists {
		return false
	} else {
		return symbols.isNullable
	}
}

func (fsb *FirstSetBuilder) initProductions() {
	fsb.symbolMap = make(map[Symbol]*Symbols)

	productions := [][]Symbol{}
	productions = append(productions, []Symbol{SymbolClassInterfaceDeclarationList})
	translationUnit := newSymbols(SymbolTranslationUnit, false, productions)
	fsb.symbolMap[SymbolTranslationUnit] = translationUnit
	fsb.symbolArr = append(fsb.symbolArr, translationUnit)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolClassInterfaceDeclarationList})
	productions = append(productions, []Symbol{SymbolClassInterfaceDeclarationList, SymbolClassInterfaceDeclaration})
	classInterfaceDeclarationList := newSymbols(SymbolClassInterfaceDeclarationList, false, productions)
	fsb.symbolMap[SymbolClassInterfaceDeclarationList] = classInterfaceDeclarationList
	fsb.symbolArr = append(fsb.symbolArr, classInterfaceDeclarationList)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolClassDeclaration})
	productions = append(productions, []Symbol{SymbolInterfaceDeclaration})
	classInterfaceDeclaration := newSymbols(SymbolClassInterfaceDeclaration, false, productions)
	fsb.symbolMap[SymbolClassInterfaceDeclaration] = classInterfaceDeclaration
	fsb.symbolArr = append(fsb.symbolArr, classInterfaceDeclaration)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolInterface, SymbolIdentifier, SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolInterface, SymbolIdentifier, SymbolLc, SymbolInterfaceMethodDeclarationStatementList, SymbolRc})
	interfaceDeclaration := newSymbols(SymbolInterfaceDeclaration, false, productions)
	fsb.symbolMap[SymbolInterfaceDeclaration] = interfaceDeclaration
	fsb.symbolArr = append(fsb.symbolArr, interfaceDeclaration)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolInterfaceMethodDeclarationStatement})
	productions = append(productions, []Symbol{SymbolInterfaceMethodDeclarationStatementList, SymbolInterfaceMethodDeclarationStatement})
	interfaceMethodDeclarationStatementList := newSymbols(SymbolInterfaceMethodDeclarationStatementList, false, productions)
	fsb.symbolMap[SymbolInterfaceMethodDeclarationStatementList] = interfaceMethodDeclarationStatementList
	fsb.symbolArr = append(fsb.symbolArr, interfaceMethodDeclarationStatementList)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolReturnValType, SymbolIdentifier, SymbolLp, SymbolRp})
	productions = append(productions, []Symbol{SymbolReturnValType, SymbolIdentifier, SymbolLp, SymbolParameterList, SymbolRp})
	interfaceMethodDeclarationStatement := newSymbols(SymbolInterfaceMethodDeclarationStatement, false, productions)
	fsb.symbolMap[SymbolInterfaceMethodDeclarationStatement] = interfaceMethodDeclarationStatement
	fsb.symbolArr = append(fsb.symbolArr, interfaceMethodDeclarationStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolClass, SymbolIdentifier, SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolClass, SymbolIdentifier, SymbolLc, SymbolClassStatementList, SymbolRc})
	productions = append(productions, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolLc, SymbolClassStatementList, SymbolRc})
	productions = append(productions, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolClassStatementList, SymbolRc})
	productions = append(productions, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolLc, SymbolClassStatementList, SymbolRc})
	productions = append(productions, []Symbol{SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc})
	productions = append(productions, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc})
	productions = append(productions, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc})
	productions = append(productions, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolAbstract, SymbolClass, SymbolIdentifier, SymbolExtendsDelcaration, SymbolImplementsDeclaration, SymbolLc, SymbolClassStatementList, SymbolRc})
	classDeclaration := newSymbols(SymbolClassDeclaration, false, productions)
	fsb.symbolMap[SymbolClassDeclaration] = classDeclaration
	fsb.symbolArr = append(fsb.symbolArr, classDeclaration)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolExtends, SymbolIdentifier})
	productions = append(productions, []Symbol{SymbolExtends, SymbolComma, SymbolIdentifier})
	extendsDelcaration := newSymbols(SymbolExtendsDelcaration, false, productions)
	fsb.symbolMap[SymbolExtendsDelcaration] = extendsDelcaration
	fsb.symbolArr = append(fsb.symbolArr, extendsDelcaration)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolImplements, SymbolIdentifier})
	productions = append(productions, []Symbol{SymbolImplementsDeclaration, SymbolComma, SymbolIdentifier})
	implementsDeclaration := newSymbols(SymbolImplementsDeclaration, false, productions)
	fsb.symbolMap[SymbolImplementsDeclaration] = implementsDeclaration
	fsb.symbolArr = append(fsb.symbolArr, implementsDeclaration)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolClassStatement})
	productions = append(productions, []Symbol{SymbolClassStatementList, SymbolClassStatement})
	classStatementList := newSymbols(SymbolClassStatementList, false, productions)
	fsb.symbolMap[SymbolClassStatementList] = classStatementList
	fsb.symbolArr = append(fsb.symbolArr, classStatementList)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolVarModifier, SymbolVarDeclaration, SymbolSemicolon})
	productions = append(productions, []Symbol{SymbolVarModifier, SymbolVarDeclaration, SymbolAssign, SymbolExpression, SymbolSemicolon})
	classStatement := newSymbols(SymbolClassStatement, false, productions)
	fsb.symbolMap[SymbolClassStatement] = classStatement
	fsb.symbolArr = append(fsb.symbolArr, classStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolPublic})
	productions = append(productions, []Symbol{SymbolProtected})
	productions = append(productions, []Symbol{SymbolPrivate})
	varModifier := newSymbols(SymbolVarModifier, false, productions)
	fsb.symbolMap[SymbolVarModifier] = varModifier
	fsb.symbolArr = append(fsb.symbolArr, varModifier)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolMethodModifier, SymbolReturnValType, SymbolIdentifier, SymbolRp, SymbolLp, SymbolBlock})
	productions = append(productions, []Symbol{SymbolMethodModifier, SymbolReturnValType, SymbolIdentifier, SymbolRp, SymbolParameterList, SymbolLp, SymbolBlock})
	methodDefinition := newSymbols(SymbolMethodDefinition, false, productions)
	fsb.symbolMap[SymbolMethodDefinition] = methodDefinition
	fsb.symbolArr = append(fsb.symbolArr, methodDefinition)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolVoid})
	productions = append(productions, []Symbol{SymbolIdentifier})
	returnValType := newSymbols(SymbolReturnValType, false, productions)
	fsb.symbolMap[SymbolReturnValType] = returnValType
	fsb.symbolArr = append(fsb.symbolArr, returnValType)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolVoid})
	productions = append(productions, []Symbol{SymbolIdentifier})
	parameterList := newSymbols(SymbolParameterList, false, productions)
	fsb.symbolMap[SymbolParameterList] = parameterList
	fsb.symbolArr = append(fsb.symbolArr, parameterList)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolIdentifier, SymbolIdentifier})
	parameter := newSymbols(SymbolParameter, false, productions)
	fsb.symbolMap[SymbolParameter] = parameter
	fsb.symbolArr = append(fsb.symbolArr, parameter)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolPublic})
	productions = append(productions, []Symbol{SymbolProtected})
	productions = append(productions, []Symbol{SymbolPrivate})
	productions = append(productions, []Symbol{SymbolAbstract})
	methodModifier := newSymbols(SymbolMethodModifier, false, productions)
	fsb.symbolMap[SymbolMethodModifier] = methodModifier
	fsb.symbolArr = append(fsb.symbolArr, methodModifier)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolLc, SymbolRc})
	productions = append(productions, []Symbol{SymbolLc, SymbolStatementList, SymbolRc})
	block := newSymbols(SymbolBlock, false, productions)
	fsb.symbolMap[SymbolBlock] = block
	fsb.symbolArr = append(fsb.symbolArr, block)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolStatement})
	productions = append(productions, []Symbol{SymbolStatementList, SymbolStatement})
	statementList := newSymbols(SymbolStatementList, false, productions)
	fsb.symbolMap[SymbolStatementList] = statementList
	fsb.symbolArr = append(fsb.symbolArr, statementList)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolExpression, SymbolSemicolon})
	productions = append(productions, []Symbol{SymbolVarDeclarationStatement})
	productions = append(productions, []Symbol{SymbolVarAssignStatement})
	productions = append(productions, []Symbol{SymbolWhileStatement})
	productions = append(productions, []Symbol{SymbolIfStatement})
	productions = append(productions, []Symbol{SymbolForStatement})
	productions = append(productions, []Symbol{SymbolBreakStatement})
	productions = append(productions, []Symbol{SymbolContinueStatement})
	productions = append(productions, []Symbol{SymbolReturnStatement})
	statement := newSymbols(SymbolStatement, false, productions)
	fsb.symbolMap[SymbolStatement] = statement
	fsb.symbolArr = append(fsb.symbolArr, statement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolWhile, SymbolLp, SymbolExpression, SymbolRp, SymbolBlock})
	whileStatement := newSymbols(SymbolWhileStatement, false, productions)
	fsb.symbolMap[SymbolWhileStatement] = whileStatement
	fsb.symbolArr = append(fsb.symbolArr, whileStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolIf, SymbolLp, SymbolExpression, SymbolRp, SymbolBlock})
	productions = append(productions, []Symbol{SymbolIf, SymbolLp, SymbolExpression, SymbolRp, SymbolBlock, SymbolElse, SymbolBlock})
	ifStatement := newSymbols(SymbolIfStatement, false, productions)
	fsb.symbolMap[SymbolIfStatement] = ifStatement
	fsb.symbolArr = append(fsb.symbolArr, ifStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolFor, SymbolRp, SymbolSemicolon, SymbolSemicolon, SymbolLp, SymbolBlock})
	productions = append(productions, []Symbol{SymbolFor, SymbolRp, SymbolExpression, SymbolSemicolon, SymbolSemicolon, SymbolLp, SymbolBlock})
	productions = append(productions, []Symbol{SymbolFor, SymbolRp, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolLp, SymbolBlock})
	productions = append(productions, []Symbol{SymbolFor, SymbolRp, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolLp, SymbolBlock})
	productions = append(productions, []Symbol{SymbolFor, SymbolRp, SymbolExpression, SymbolSemicolon, SymbolSemicolon, SymbolExpression, SymbolLp, SymbolBlock})
	productions = append(productions, []Symbol{SymbolFor, SymbolRp, SymbolSemicolon, SymbolSemicolon, SymbolExpression, SymbolLp, SymbolBlock})
	productions = append(productions, []Symbol{SymbolFor, SymbolRp, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolExpression, SymbolLp, SymbolBlock})
	productions = append(productions, []Symbol{SymbolFor, SymbolRp, SymbolSemicolon, SymbolExpression, SymbolSemicolon, SymbolLp, SymbolBlock})
	forStatement := newSymbols(SymbolForStatement, false, productions)
	fsb.symbolMap[SymbolForStatement] = forStatement
	fsb.symbolArr = append(fsb.symbolArr, forStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolBreak, SymbolSemicolon})
	productions = append(productions, []Symbol{SymbolBreak, SymbolExpression, SymbolSemicolon})
	breakStatement := newSymbols(SymbolBreakStatement, false, productions)
	fsb.symbolMap[SymbolBreakStatement] = breakStatement
	fsb.symbolArr = append(fsb.symbolArr, breakStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolContinue, SymbolSemicolon})
	productions = append(productions, []Symbol{SymbolContinue, SymbolExpression, SymbolSemicolon})
	continueStatement := newSymbols(SymbolContinueStatement, false, productions)
	fsb.symbolMap[SymbolContinueStatement] = continueStatement
	fsb.symbolArr = append(fsb.symbolArr, continueStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolReturn, SymbolSemicolon})
	productions = append(productions, []Symbol{SymbolReturn, SymbolExpression, SymbolSemicolon})
	returnStatement := newSymbols(SymbolReturnStatement, false, productions)
	fsb.symbolMap[SymbolReturnStatement] = returnStatement
	fsb.symbolArr = append(fsb.symbolArr, returnStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolVarDeclaration, SymbolSemicolon})
	varDeclarationStatement := newSymbols(SymbolVarDeclarationStatement, false, productions)
	fsb.symbolMap[SymbolVarDeclarationStatement] = varDeclarationStatement
	fsb.symbolArr = append(fsb.symbolArr, varDeclarationStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolVarDeclaration, SymbolAssign, SymbolExpression, SymbolSemicolon})
	productions = append(productions, []Symbol{SymbolVarCallExpression, SymbolAssign, SymbolExpression, SymbolSemicolon})
	varAssignStatement := newSymbols(SymbolVarAssignStatement, false, productions)
	fsb.symbolMap[SymbolVarAssignStatement] = varAssignStatement
	fsb.symbolArr = append(fsb.symbolArr, varAssignStatement)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolIdentifier, SymbolIdentifier})
	varDeclaration := newSymbols(SymbolVarDeclaration, false, productions)
	fsb.symbolMap[SymbolVarDeclaration] = varDeclaration
	fsb.symbolArr = append(fsb.symbolArr, varDeclaration)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolStringLiteral})
	productions = append(productions, []Symbol{SymbolIntLiteral})
	productions = append(productions, []Symbol{SymbolDoubleLiteral})
	productions = append(productions, []Symbol{SymbolNull})
	productions = append(productions, []Symbol{SymbolTrue})
	productions = append(productions, []Symbol{SymbolFalse})
	productions = append(productions, []Symbol{SymbolIdentifier})
	productions = append(productions, []Symbol{SymbolNewObjExpression})
	productions = append(productions, []Symbol{SymbolCallExpression})
	expression := newSymbols(SymbolExpression, false, productions)
	fsb.symbolMap[SymbolExpression] = expression
	fsb.symbolArr = append(fsb.symbolArr, expression)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolMethodCallExpression})
	productions = append(productions, []Symbol{SymbolVarCallExpression})
	callExpression := newSymbols(SymbolCallExpression, false, productions)
	fsb.symbolMap[SymbolCallExpression] = callExpression
	fsb.symbolArr = append(fsb.symbolArr, callExpression)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolCallExpression, SymbolDot, SymbolIdentifier, SymbolLp, SymbolRp})
	productions = append(productions, []Symbol{SymbolCallExpression, SymbolDot, SymbolIdentifier, SymbolArgumentList, SymbolLp, SymbolRp})
	methodCallExpression := newSymbols(SymbolMethodCallExpression, false, productions)
	fsb.symbolMap[SymbolMethodCallExpression] = methodCallExpression
	fsb.symbolArr = append(fsb.symbolArr, methodCallExpression)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolIdentifier})
	productions = append(productions, []Symbol{SymbolThis})
	productions = append(productions, []Symbol{SymbolCallExpression, SymbolDot, SymbolIdentifier})
	varCallExpression := newSymbols(SymbolVarCallExpression, false, productions)
	fsb.symbolMap[SymbolVarCallExpression] = varCallExpression
	fsb.symbolArr = append(fsb.symbolArr, varCallExpression)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolNew, SymbolIdentifier, SymbolLp, SymbolRp})
	productions = append(productions, []Symbol{SymbolNew, SymbolIdentifier, SymbolLp, SymbolArgumentList, SymbolRp})
	newObjExpression := newSymbols(SymbolNewObjExpression, false, productions)
	fsb.symbolMap[SymbolNewObjExpression] = newObjExpression
	fsb.symbolArr = append(fsb.symbolArr, newObjExpression)

	productions = [][]Symbol{}
	productions = append(productions, []Symbol{SymbolExpression})
	productions = append(productions, []Symbol{SymbolArgumentList, SymbolComma, SymbolExpression})
	argumentList := newSymbols(SymbolArgumentList, false, productions)
	fsb.symbolMap[SymbolArgumentList] = argumentList
	fsb.symbolArr = append(fsb.symbolArr, argumentList)

	stringLiteral := newSymbols(SymbolStringLiteral, false, nil)
	fsb.symbolMap[SymbolStringLiteral] = stringLiteral
	fsb.symbolArr = append(fsb.symbolArr, stringLiteral)

	intLiteral := newSymbols(SymbolIntLiteral, false, nil)
	fsb.symbolMap[SymbolIntLiteral] = intLiteral
	fsb.symbolArr = append(fsb.symbolArr, intLiteral)

	doubleLiteral := newSymbols(SymbolDoubleLiteral, false, nil)
	fsb.symbolMap[SymbolDoubleLiteral] = doubleLiteral
	fsb.symbolArr = append(fsb.symbolArr, doubleLiteral)

	null := newSymbols(SymbolNull, false, nil)
	fsb.symbolMap[SymbolNull] = null
	fsb.symbolArr = append(fsb.symbolArr, null)

	symbolsTrue := newSymbols(SymbolTrue, false, nil)
	fsb.symbolMap[SymbolTrue] = symbolsTrue
	fsb.symbolArr = append(fsb.symbolArr, symbolsTrue)

	symbolsFalse := newSymbols(SymbolFalse, false, nil)
	fsb.symbolMap[SymbolFalse] = symbolsFalse
	fsb.symbolArr = append(fsb.symbolArr, symbolsFalse)

	identifier := newSymbols(SymbolIdentifier, false, nil)
	fsb.symbolMap[SymbolIdentifier] = identifier
	fsb.symbolArr = append(fsb.symbolArr, identifier)

	symbolsNew := newSymbols(SymbolNew, false, nil)
	fsb.symbolMap[SymbolNew] = symbolsNew
	fsb.symbolArr = append(fsb.symbolArr, symbolsNew)

	lp := newSymbols(SymbolLp, false, nil)
	fsb.symbolMap[SymbolLp] = lp
	fsb.symbolArr = append(fsb.symbolArr, lp)

	rp := newSymbols(SymbolRp, false, nil)
	fsb.symbolMap[SymbolRp] = rp
	fsb.symbolArr = append(fsb.symbolArr, rp)

	dot := newSymbols(SymbolDot, false, nil)
	fsb.symbolMap[SymbolDot] = dot
	fsb.symbolArr = append(fsb.symbolArr, dot)

	lc := newSymbols(SymbolLc, false, nil)
	fsb.symbolMap[SymbolLc] = lc
	fsb.symbolArr = append(fsb.symbolArr, lc)

	rc := newSymbols(SymbolRc, false, nil)
	fsb.symbolMap[SymbolRc] = rc
	fsb.symbolArr = append(fsb.symbolArr, rc)

	comma := newSymbols(SymbolComma, false, nil)
	fsb.symbolMap[SymbolComma] = comma
	fsb.symbolArr = append(fsb.symbolArr, comma)

	symbolsReturn := newSymbols(SymbolReturn, false, nil)
	fsb.symbolMap[SymbolReturn] = symbolsReturn
	fsb.symbolArr = append(fsb.symbolArr, symbolsReturn)

	symbolsContinue := newSymbols(SymbolContinue, false, nil)
	fsb.symbolMap[SymbolContinue] = symbolsContinue
	fsb.symbolArr = append(fsb.symbolArr, symbolsContinue)

	symbolsBreak := newSymbols(SymbolBreak, false, nil)
	fsb.symbolMap[SymbolBreak] = symbolsBreak
	fsb.symbolArr = append(fsb.symbolArr, symbolsBreak)

	semicolon := newSymbols(SymbolSemicolon, false, nil)
	fsb.symbolMap[SymbolSemicolon] = semicolon
	fsb.symbolArr = append(fsb.symbolArr, semicolon)

	symbolsFor := newSymbols(SymbolFor, false, nil)
	fsb.symbolMap[SymbolFor] = symbolsFor
	fsb.symbolArr = append(fsb.symbolArr, symbolsFor)

	symbolsIf := newSymbols(SymbolIf, false, nil)
	fsb.symbolMap[SymbolIf] = symbolsIf
	fsb.symbolArr = append(fsb.symbolArr, symbolsIf)

	symbolsElse := newSymbols(SymbolElse, false, nil)
	fsb.symbolMap[SymbolElse] = symbolsElse
	fsb.symbolArr = append(fsb.symbolArr, symbolsElse)

	symbolsWhile := newSymbols(SymbolWhile, false, nil)
	fsb.symbolMap[SymbolWhile] = symbolsWhile
	fsb.symbolArr = append(fsb.symbolArr, symbolsWhile)

	assign := newSymbols(SymbolAssign, false, nil)
	fsb.symbolMap[SymbolAssign] = assign
	fsb.symbolArr = append(fsb.symbolArr, assign)

	class := newSymbols(SymbolClass, false, nil)
	fsb.symbolMap[SymbolClass] = class
	fsb.symbolArr = append(fsb.symbolArr, class)

	this := newSymbols(SymbolThis, false, nil)
	fsb.symbolMap[SymbolThis] = this
	fsb.symbolArr = append(fsb.symbolArr, this)

	symbolsInterface := newSymbols(SymbolInterface, false, nil)
	fsb.symbolMap[SymbolInterface] = symbolsInterface
	fsb.symbolArr = append(fsb.symbolArr, symbolsInterface)

	abstract := newSymbols(SymbolAbstract, false, nil)
	fsb.symbolMap[SymbolAbstract] = abstract
	fsb.symbolArr = append(fsb.symbolArr, abstract)

	implements := newSymbols(SymbolImplements, false, nil)
	fsb.symbolMap[SymbolImplements] = implements
	fsb.symbolArr = append(fsb.symbolArr, implements)

	extends := newSymbols(SymbolExtends, false, nil)
	fsb.symbolMap[SymbolExtends] = extends
	fsb.symbolArr = append(fsb.symbolArr, extends)

	void := newSymbols(SymbolVoid, false, nil)
	fsb.symbolMap[SymbolVoid] = void
	fsb.symbolArr = append(fsb.symbolArr, void)

	public := newSymbols(SymbolPublic, false, nil)
	fsb.symbolMap[SymbolPublic] = public
	fsb.symbolArr = append(fsb.symbolArr, public)

	private := newSymbols(SymbolPrivate, false, nil)
	fsb.symbolMap[SymbolPrivate] = private
	fsb.symbolArr = append(fsb.symbolArr, private)

	protected := newSymbols(SymbolProtected, false, nil)
	fsb.symbolMap[SymbolProtected] = protected
	fsb.symbolArr = append(fsb.symbolArr, protected)
}

func (fsb *FirstSetBuilder) runFirstSets() {
	for fsb.runFirstSetPass {
		fsb.runFirstSetPass = false
		for _, symbols := range fsb.symbolArr {
			fsb.addSymbolFirstSet(symbols)
		}
	}
}

func (fsb *FirstSetBuilder) addSymbolFirstSet(symbols *Symbols) {
	// 如果符号是终结符那它的firstSet就是它自己
	if symbols.value.isTerminals() {
		if _, exists := symbols.firstSet[symbols.value]; !exists {
			symbols.firstSet[symbols.value] = struct{}{}
		}

		return
	}

	// 遍历该符号的所有生成式
	for _, p := range symbols.productions {
		if len(p) == 0 {
			continue
		}

		if p[0].isTerminals() {
			// 如果生成式的第一个符号是终结符并且该符号不在当前符号的firstSet中则加入进来
			if _, exists := symbols.firstSet[p[0]]; !exists {
				fsb.runFirstSetPass = true
				symbols.firstSet[p[0]] = struct{}{}
			}
		} else if !p[0].isTerminals() {
			// 如果生成式的第一个符号是非终结符则遍历生成式的每个符号
			for _, curSymbol := range p {
				curSymbols := fsb.symbolMap[curSymbol]
				// 将每个符号的firstSet中的符号添加到该符号的firstSet中并标识runFirstSetPass为false
				for s, _ := range curSymbols.firstSet {
					if _, exists := symbols.firstSet[s]; !exists {
						symbols.firstSet[s] = struct{}{}
						fsb.runFirstSetPass = false
					}
				}
				if !curSymbols.isNullable {
					break
				}
			}
		}
	}
}
