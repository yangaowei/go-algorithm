package strategy

type calculator struct {
	alg IAlg
}

func NewCalculator() *calculator {

	return &calculator{}
}

// 改变算法
func (calculator *calculator) ChangeAlg(alg IAlg) {
	calculator.alg = alg
}

func (self *calculator) Cal() {
	self.alg.used()
}
