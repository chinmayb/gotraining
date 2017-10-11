package aggregators

type SumAggregator struct {
	accumulator int
	count       int
}

func (ag *SumAggregator) Reset() {
	ag.accumulator, ag.count = 0, 0

}

func (ag *SumAggregator) Take(num int) {
	ag.accumulator += num
	ag.count++
}

func (ag *SumAggregator) Aggregate() (int, int) {
	return ag.accumulator, ag.count
}
