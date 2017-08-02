package aggregops

type Aggregator interface {
	Reset()
	Take(num int)
	Aggregate() (int, int)
}

type Multiplier interface {
	Factor() string
}

func Simple(nums []int, agg Aggregator) (int, int) {
	agg.Reset()
	for _,v := range nums {
		agg.Take(v)
	}
	result, count := agg.Aggregate()
	//conv, err := ag.(SumAggregater)
	//conv, ok := agg.(Multiplier)

	/*if ok {
		result = result * conv.Factor()
	}*/
	return result, count
}

func Partitioned(nums []int, agg Aggregator) (int, int){
    Aggregate
	return 0, 0
}