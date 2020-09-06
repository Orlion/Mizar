package parser

import (
	"github.com/sirupsen/logrus"
	"mizar/log"
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
			if oldAction, exists := jump[symbol]; exists {
				log.Warn(logrus.Fields{
					"gsStateNum":        gs.stateNum,
					"symbol":            symbol,
					"oldAction":         oldAction,
					"newActionStateNum": childGs.stateNum,
				}, "shift conflict")
			}
			jump[symbol] = newShiftAction(childGs.stateNum)
		}

		reduceMap := gs.makeReduce()
		for symbol, action := range reduceMap {
			if oldAction, exists := jump[symbol]; exists {
				log.Warn(logrus.Fields{
					"gsStateNum":                gs.stateNum,
					"symbol":                    symbol,
					"oldActionIsReduce":         oldAction.isReduce,
					"oldActionShiftStateNum":    oldAction.shiftStateNum,
					"oldActionReduceProduction": oldAction.reduceProduction,
					"newAction":                 action.reduceProduction.GetCode(),
				}, "shift reduce conflict")
			}
			jump[symbol] = action
		}

		m[gs.stateNum] = jump
	}

	return m
}
