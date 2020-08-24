package parser

import (
	"fmt"
	"mizar/utils"
	"strconv"
)

type GrammarState struct {
	gsm            *GrammarStateManager
	stateNum       int
	productions    []*Production
	transition     map[Symbol]*GrammarState // 跳转关系，key为输入的字符，GrammarState为跳转到的状态节点
	closureSet     []*Production            // 当前节点做闭包操作产生的新生成式
	closureKeySet  map[string]struct{}
	partition      map[Symbol][]*Production // 用来分区操作
	transitionDone bool
}

func newGrammarState(gsm *GrammarStateManager, stateNum int, productions []*Production) (gs *GrammarState) {
	gs = new(GrammarState)
	gs.gsm = gsm
	gs.stateNum = stateNum
	gs.productions = productions
	gs.closureKeySet = make(map[string]struct{})

	return
}

func (gs *GrammarState) makeClosure() {
	pStack := utils.NewStack()
	// 先将当前节点的所有生成式加入到栈中
	for _, p := range gs.productions {
		pStack.Push(p)
	}

	for !pStack.Empty() {
		// 弹出栈顶生成式
		pInter := pStack.Pop()
		production := pInter.(*Production)
		symbol := production.getDotSymbol()
		// 如果生成式.的右边的符号是终结符则直接跳过
		if symbol.isTerminals() {
			continue
		}

		// 从pm中查出以该符号为目标符号的所有生成式
		closures := getProductionManager().getProductions(symbol)
		// 获取当前生成式的lookAhead集合
		lookAhead := production.computeFirstSetOfBetaAndC()

		for _, oldProduct := range closures {
			newProduct := oldProduct.cloneSelf()
			newProduct.addLookAheadSet(lookAhead)

			if _, exists := gs.closureKeySet[newProduct.getCode()]; !exists {
				gs.closureSet = append(gs.closureSet, newProduct)
				gs.closureKeySet[newProduct.getCode()] = struct{}{}

				pStack.Push(newProduct)

				gs.removeRedundantProduction(newProduct)
			}
		}
	}
}

func (gs *GrammarState) removeRedundantProduction(p *Production) {
	target := gs.closureSet[:0]
	for _, item := range gs.closureSet {
		if p.coverUp(item) {
			break
		}

		target = append(target, item)
	}

	gs.closureSet = target
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
		newGsPs := []*Production{}
		for _, p := range ps {
			newGsPs = append(newGsPs, p.dotForward())
		}
		newGs = gs.gsm.getGrammarState(newGsPs)
		gs.transition[symbol] = newGs
	}

	gs.transitionDone = true
}

// 扩展下一个节点
func (gs *GrammarState) extendTransition() {
	for _, childGs := range gs.transition {
		if childGs.transitionDone == false {
			childGs.createTransition()
		}
	}
}

func (gs *GrammarState) createTransition() {
	gs.makeClosure()
	gs.makePartition()
	gs.makeTransition()
	gs.extendTransition()
}

func (gs *GrammarState) print() {
	fmt.Println("GrammarState: " + strconv.Itoa(gs.stateNum))

	for _, p := range gs.productions {
		p.print()
	}

	fmt.Println()
	fmt.Println()

	for _, child := range gs.transition {
		child.print()
	}
}
