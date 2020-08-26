package parser

import (
	"sort"
	"strings"
)

type GrammarStateManager struct {
	stateNumCount int
	states        map[string]*GrammarState
}

func newGrammarStateManager() (gsm *GrammarStateManager) {
	gsm = new(GrammarStateManager)
	gsm.stateNumCount = 0
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
	gs := newGrammarState(gsm, gsm.stateNumCount, getProductionManager().getProductions(SymbolTranslationUnit), -1, "")
	gs.createTransition()

	return gs
}
