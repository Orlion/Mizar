package parser

import (
	"fmt"
	"strconv"
)

type GrammarState struct {
	gsm           *GrammarStateManager
	stateNum      int
	productions   []*Production
	transition    map[Symbol]*GrammarState // 跳转关系，key为输入的字符，GrammarState为跳转到的状态节点
	closureSet    []*Production            // 当前节点做闭包操作产生的新生成式
	closureKeySet map[string]struct{}
	partition     map[Symbol][]*Production // 用来分区操作
}

func newGrammarState(gsm *GrammarStateManager, stateNum int, productions []*Production) (gs *GrammarState) {
	gs = new(GrammarState)
	gs.gsm = gsm
	gs.stateNum = stateNum
	gs.productions = productions

	fmt.Println("newGrammarState: " + strconv.Itoa(stateNum))
	for _, p := range productions {
		p.print()
	}
	fmt.Println()
	fmt.Println()

	return
}

func (gs *GrammarState) closure() {
	var (
		production *Production
		ps         []*Production
	)

	gs.closureKeySet = make(map[string]struct{})
	for _, p := range gs.productions {
		if _, exists := gs.closureKeySet[p.str]; !exists {
			gs.closureSet = append(gs.closureSet, p)
			gs.closureKeySet[p.str] = struct{}{}
		}
	}

	// 如果.右侧是非终结符则将其生成式递归加入进来
	i := 0
	for {
		if i >= len(gs.closureSet) {
			break
		}
		production = gs.closureSet[i]
		ps = gs.gsm.pm.getProductions(production.getDotSymbol())
		for _, p := range ps {
			if _, exists := gs.closureKeySet[p.str]; !exists {
				gs.closureSet = append(gs.closureSet, p)
				gs.closureKeySet[p.str] = struct{}{}
			}
		}
		i++
	}
}

// 分区, 将.右侧相同非终结符的生成式划分到同一分区
func (gs *GrammarState) makePartition() {
	gs.partition = make(map[Symbol][]*Production)
	for _, p := range gs.closureSet {
		gs.partition[p.getDotSymbol()] = append(gs.partition[p.getDotSymbol()], p)
	}
}

// .右移一位生成下一节点
func (gs *GrammarState) makeTransition() {
	var newGs *GrammarState

	gs.transition = make(map[Symbol]*GrammarState)

	for symbol, ps := range gs.partition {
		newGs = gs.gsm.getGrammarState(ps)
		gs.transition[symbol] = newGs
	}
}

// 扩展下一个节点
func (gs *GrammarState) extendTransition() {
	for _, childGs := range gs.transition {
		childGs.createTransition()
	}
}

func (gs *GrammarState) createTransition() {
	gs.closure()
	gs.makePartition()
	gs.makeTransition()
	gs.extendTransition()
}

type GrammarStateManager struct {
	stateNumCount int
	pm            *ProductionManager
}

func newGrammarStateManager(pm *ProductionManager) (gsm *GrammarStateManager) {
	gsm = new(GrammarStateManager)
	gsm.pm = pm
	gsm.stateNumCount = 0
	return
}

func (gsm *GrammarStateManager) getGrammarState(ps []*Production) (gs *GrammarState) {
	gsm.stateNumCount++
	gs = newGrammarState(gsm, gsm.stateNumCount, ps)
	return
}

func (gsm *GrammarStateManager) build() *GrammarState {
	gsm.stateNumCount++
	gs := newGrammarState(gsm, gsm.stateNumCount, gsm.pm.getProductions(SymbolTranslationUnit))
	gs.createTransition()

	return gs
}
