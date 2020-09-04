package parser

import (
	"sort"
	"strings"
)

type GrammarStateManager struct {
	stateNumCount int
	states        map[string]*GrammarState
	gs            *GrammarState
}

func newGrammarStateManager() (gsm *GrammarStateManager) {
	gsm = new(GrammarStateManager)
	gsm.stateNumCount = -1
	gsm.states = make(map[string]*GrammarState)
	return
}

func (gsm *GrammarStateManager) getGrammarState(ps []*Production, fromStateNum int, edge Symbol) (gs *GrammarState) {
	keyList := make([]string, 0)
	for _, p := range ps {
		keyList = append(keyList, p.GetCode())
	}

	sort.Strings(keyList)
	key := strings.Join(keyList, " | ")

	if s, exists := gsm.states[key]; exists {
		gs = s
	} else {
		gsm.stateNumCount++
		gs = newGrammarState(gsm, gsm.stateNumCount, ps, fromStateNum, edge)
		gsm.states[key] = gs
	}

	return
}

func (gsm *GrammarStateManager) build() *GrammarState {
	gsm.stateNumCount++
	ps := getProductionManager().getProductions(SymbolTranslationUnit)
	gs := newGrammarState(gsm, gsm.stateNumCount, getProductionManager().getProductions(SymbolTranslationUnit), 0, "")

	keyList := make([]string, 0)
	for _, p := range ps {
		keyList = append(keyList, p.GetCode())
	}

	sort.Strings(keyList)
	key := strings.Join(keyList, " | ")

	gsm.states[key] = gs

	gs.createTransition()

	return gs
}

// map[currentState]map[currentInput]action
func (gsm *GrammarStateManager) getLRStateTable() map[int]map[Symbol]*Action {
	m := make(map[int]map[Symbol]*Action)

	for _, gs := range gsm.states {
		jump := make(map[Symbol]*Action)
		for symbol, childGs := range gs.transition {
			jump[symbol] = newShiftAction(childGs.stateNum)
		}

		reduceMap := gs.makeReduce()
		for symbol, action := range reduceMap {
			jump[symbol] = action
		}

		m[gs.stateNum] = jump
	}

	return m
}
