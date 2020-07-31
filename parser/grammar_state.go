package parser

type GrammarState struct {
	stateNumCount       int
	stateNum            int
	grammarStateManager *GrammarStateManager
	productions         []*Production
	transition          map[int]*GrammarState // 跳转关系，key为输入的字符，GrammarState为跳转到的状态节点
	closureSet          []*Production         // 当前节点做闭包操作产生的新生成式
	productionManager   *ProductionManager
	partition           map[int][]*Production // 用来分区操作
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

}
