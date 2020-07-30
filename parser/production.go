package parser

type Production struct {
	left   string   // 左侧非终结符
	right  []string // 右侧符号列表
	dotPos int      // .的位置
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
