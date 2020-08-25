package parser

import (
	"fmt"
	"strings"
)

type Production struct {
	code      string   // 转成字符串表示, 用来标识生成式唯一
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

	p.lookAhead = map[Symbol]struct{}{EOISymbol: {}}

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
	fmt.Println(p.getCode())
}

/**
 * 计算First(β C)
 * 将β与C前后相连再计算他们的First集合，如果β里面的每一项都是nullable的，那么First(β C)就是First(β) 并上First(C)
 * 由于C必定是终结符的组合，所以First(C)等于C的第一个终结符，例如C = {+, *, EOI} 那么First(C) = {+}
 */
func (p *Production) computeFirstSetOfBetaAndC() map[Symbol]struct{} {
	set := []Symbol{}
	setKeys := map[Symbol]struct{}{}

	for i := p.dotPos + 1; i < len(p.right); i++ {
		if _, exists := setKeys[p.right[i]]; !exists {
			set = append(set, p.right[i])
		}
	}

	for s, _ := range p.lookAhead {
		if _, exists := setKeys[s]; !exists {
			set = append(set, s)
		}
	}

	pm := getProductionManager()

	firstSet := make(map[Symbol]struct{})

	for _, s := range set {
		lookAhead := pm.getFirstSetBuilder().getFirstSet(s)
		for s1, _ := range lookAhead {
			firstSet[s1] = struct{}{}
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

	for symbol, _ := range p.lookAhead {
		if _, exists := production.lookAhead[symbol]; !exists {
			return -1
		}
	}

	return 0
}

func (p *Production) addLookAheadSet(lookAhead map[Symbol]struct{}) {
	p.lookAhead = lookAhead
}

func (p *Production) cloneSelf() *Production {
	product := newProduction(p.left, p.right, p.dotPos)
	product.lookAhead = make(map[Symbol]struct{})
	for s, v := range p.lookAhead {
		product.lookAhead[s] = v
	}

	return product
}

func (p *Production) equals(production *Production) bool {
	return p.getCode() == production.getCode() && 0 == p.lookAheadCompare(production)
}

func (p *Production) getCode() string {
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
		for k, _ := range p.lookAhead {
			codeBuilder.WriteString(string(k))
			codeBuilder.WriteString(" ")
		}
		codeBuilder.WriteString(")")

		p.code = codeBuilder.String()
	}

	return p.code
}
