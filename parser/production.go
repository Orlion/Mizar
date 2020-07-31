package parser

import "fmt"

type Production struct {
	left   string   // 左侧非终结符
	right  []string // 右侧符号列表
	dotPos int      // .的位置
}

type ProductionManager struct {
	productionMap map[int][]*Production
}

func newProduction(left string, right []string, dotPos int) *Production {
	return &Production{
		left:   left,
		right:  right,
		dotPos: dotPos,
	}
}

// .前移
func (p *Production) dotForward() *Production {
	return newProduction(p.left, p.right, p.dotPos+1)
}

func (p *Production) getDotSymbol() string {
	return p.right[p.dotPos]
}

func (p *Production) print() {
	fmt.Printf("%s -> ", p.left)
	for k, v := range p.right {
		if p.dotPos == k {
			fmt.Print(". ")
		}
		fmt.Print("" + v)
	}
	fmt.Println()
}

func newProductionManager() *ProductionManager {
	pm := new(ProductionManager)

	pm.productionMap[1] = newProduction(1, nil, 0)

	return pm
}
