package parser

import "fmt"

type Production struct {
	str       string   // 转成字符串表示
	left      Symbol   // 左侧非终结符
	right     []Symbol // 右侧符号列表
	dotPos    int      // .的位置
	lookAhead map[Symbol]struct{}
}

func newProduction(left Symbol, right []Symbol, dotPos int) (p *Production) {
	p = &Production{
		left:   left,
		right:  right,
		dotPos: dotPos,
	}

	p.str = fmt.Sprintf("%s ->   ", left)
	for k, v := range p.right {
		if p.dotPos == k {
			p.str += ".   "
		}
		p.str += (string(v) + "   ")
	}

	return
}

// .前移
func (p *Production) dotForward() *Production {
	return newProduction(p.left, p.right, p.dotPos+1)
}

func (p *Production) getDotSymbol() Symbol {
	if p.dotPos >= len(p.right) {
		return NilSymbol
	}
	return p.right[p.dotPos]
}

func (p *Production) print() {
	fmt.Println(p.str)
}

// 计算β与C并集
func (p *Production) computeFirstSetOfBetaAndC() map[Symbol]struct{} {
	pm := getProductionManager()

	set := p.lookAhead

	for _, item := range p.right {
		firstSet := pm.getFirstSetBuilder().getFirstSet(item)
		for symbol, _ := range firstSet {
			if _, exists := set[symbol]; !exists {
				set[symbol] = struct{}{}
			}
		}

		if !pm.getFirstSetBuilder().isSymbolNullable(item) {
			break
		}
	}

	return set
}

// equals还要判断lookAhead是否相同

// 判断是否能够覆盖老表达式, 新表达式的lookAhead大于旧表达式的lookAhead
func (p *Production) coverUp(oldProduction *Production) bool {
	if p.productionEquals(oldProduction) && p.lookAheadCompare(oldProduction) > 0 {
		return true
	}

	return false
}

func (p *Production) productionEquals(production *Production) bool {
	if p.left != production.left {
		return false
	}

	if p.dotPos != production.dotPos {
		return false
	}

	if len(p.right) != len(production.right) {
		return false
	}

	for k, v := range p.right {
		if v != production.right[k] {
			return false
		}
	}

	return true
}

func (p *Production) lookAheadCompare(production *Production) int {
	if len(p.lookAhead) < len(production.lookAhead) {
		return -1
	}

	if len(p.lookAhead) > len(production.lookAhead) {
		return 1
	}

	for symbol, _ := range p.lookAhead {
		if _, exists := production.lookAhead[symbol]; !exists {
			return -1
		}
	}

	return 0
}

func (p *Production) addLookAhead(lookAhead []Symbol) {

}
