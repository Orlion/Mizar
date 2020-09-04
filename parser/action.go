package parser

type Action struct {
	isReduce         bool        // true: reduce操作，false: shift操作
	reduceProduction *Production // reduce的生成式
	shiftStateNum    int         // shift到的状态节点的状态码
}

func newReduceAction(reduceProduction *Production) *Action {
	return &Action{isReduce: true, reduceProduction: reduceProduction}
}

func newShiftAction(shiftStateNum int) *Action {
	return &Action{isReduce: false, shiftStateNum: shiftStateNum}
}
