package parser

type GrammarState struct {
	stateNum            int
	productions         []*Production
	transition          map[Symbol]*GrammarState // 跳转关系，key为输入的字符，GrammarState为跳转到的状态节点
	closureSet          []*Production         // 当前节点做闭包操作产生的新生成式
	productionManager   *ProductionManager
	partition           map[Symbol][]*Production // 用来分区操作
}

func newGrammarState(stateNum int, productions []*Production) (gs *GrammarState) {
	gs = new(GrammarState)
	gs.stateNum = stateNum
	gs.productions = productions

	return
}

func (gs *GrammarState) closure() {
	var (
		production *Production
		ps []*Production
	)

	// 如果.右侧是非终结符则将其生成式递归加入进来
	i := 0
	for {
		if i >= len(gs.productions) {
			break
		}
		production = gs.productions[i]
		ps = gs.productionManager.getProductions(production.getDotSymbol())
		gs.productions = append(gs.productions, ps...)
		i++
	}

	// 分区

	// 生成下一节点
}

type GrammarStateManager struct {
	stateNumCount int
}

func newGrammarStateManager() (gsm *GrammarStateManager) {
	gsm = new(GrammarStateManager)
	gsm.stateNumCount = 0
	return
}

func (gsm *GrammarStateManager) build() {
	gs := newGrammarState(0, []*Production{newProduction(SymbolTranslationUnit, []Symbol{SymbolClassDeclarationList}, 0)})
	gs.closure()
}
