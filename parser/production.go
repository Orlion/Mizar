package parser

import (
	"fmt"
	"strings"
)

type Production struct {
	code          string   // 转成字符串表示, 用来标识生成式唯一
	left          Symbol   // 左侧非终结符
	right         []Symbol // 右侧符号列表
	dotPos        int      // .的位置
	lookAhead     []Symbol
	lookAheadKeys map[Symbol]struct{}
	productionNum int
}

func newProduction(left Symbol, right []Symbol, dotPos int) (p *Production) {
	p = &Production{
		left:   left,
		right:  right,
		dotPos: dotPos,
	}

	p.lookAhead = []Symbol{EOISymbol}
	p.lookAheadKeys = map[Symbol]struct{}{EOISymbol: {}}

	return
}

// .前移
func (p *Production) dotForward() *Production {
	newProduct := newProduction(p.left, p.right, p.dotPos+1)
	for s, v := range p.lookAhead {
		newProduct.lookAhead[s] = v
	}

	return newProduct
}

func (p *Production) getDotSymbol() Symbol {
	if p.dotPos >= len(p.right) {
		return NilSymbol
	}

	return p.right[p.dotPos]
}

func (p *Production) print() {
	fmt.Println(p.GetCode())
}

/**
 * 计算First(β C)
 * 将β与C前后相连再计算他们的First集合，如果β里面的每一项都是nullable的，那么First(β C)就是First(β) 并上First(C)
 * 由于C必定是终结符的组合，所以First(C)等于C的第一个终结符，例如C = {+, *, EOI} 那么First(C) = {+}
 */
func (p *Production) computeFirstSetOfBetaAndC() []Symbol {
	set := []Symbol{}
	setKeys := map[Symbol]struct{}{}

	for i := p.dotPos + 1; i < len(p.right); i++ {
		if _, exists := setKeys[p.right[i]]; !exists {
			set = append(set, p.right[i])
		}
	}

	for _, s := range p.lookAhead {
		if _, exists := setKeys[s]; !exists {
			set = append(set, s)
		}
	}

	pm := getProductionManager()

	firstSet := []Symbol{}

	for _, s := range set {
		lookAhead := pm.getFirstSetBuilder().getFirstSet(s)
		for _, s1 := range lookAhead {
			firstSet = append(firstSet, s1)
		}

		if !pm.getFirstSetBuilder().isSymbolNullable(s) {
			break
		}
	}

	return firstSet
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

	for _, symbol := range p.lookAhead {
		if _, exists := production.lookAheadKeys[symbol]; !exists {
			return -1
		}
	}

	return 0
}

func (p *Production) addLookAheadSet(lookAhead []Symbol) {
	p.lookAhead = lookAhead
}

func (p *Production) cloneSelf() *Production {
	product := newProduction(p.left, p.right, p.dotPos)
	product.lookAheadKeys = make(map[Symbol]struct{})
	for _, s := range p.lookAhead {
		if _, exists := product.lookAheadKeys[s]; !exists {
			product.lookAhead = append(product.lookAhead, s)
			product.lookAheadKeys[s] = struct{}{}
		}
	}

	return product
}

func (p *Production) GetCode() string {
	if p.code == "" {
		var codeBuilder strings.Builder

		codeBuilder.WriteString(fmt.Sprintf("%s ->   ", p.left))
		for k, v := range p.right {
			if p.dotPos == k {
				codeBuilder.WriteString(".   ")
			}
			codeBuilder.WriteString(string(v))
			codeBuilder.WriteString("   ")
		}

		codeBuilder.WriteString("(")

		// 将无序的set转为有序的list
		list := make([]string, 0)

		for _, k := range p.lookAhead {
			list = append(list, string(k))
		}

		// 对list排序
		for _, s := range list {
			codeBuilder.WriteString(s)
			codeBuilder.WriteString(" ")
		}

		codeBuilder.WriteString(")")

		p.code = codeBuilder.String()
	}

	return p.code
}

// 判断能否reduce
func (p *Production) canBeReduce() bool {
	return p.dotPos >= len(p.right)
}
