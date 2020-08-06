package parser

import "testing"

func TestBuild(t *testing.T) {
	pm := newProductionManager()
	gsm := newGrammarStateManager(pm)
	gsm.build()
}
