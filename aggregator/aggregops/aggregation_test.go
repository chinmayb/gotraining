package aggregops

import "testing"
import ag1 "github.com/chinmayb/gotraining/aggregator/aggregators"


type FakeAggregator int

func (f *FakeAggregator) Reset() {
	*f = 0
}

func (f *FakeAggregator) Take(n int) {
	_ = n
}


func (f *FakeAggregator) Aggregate() (int, int) {
	return 0,0

}

//TestSimple does unit testing of Simple
func TestSimple(t *testing.T) {
	nums := []int {1,2,3,4,5}
	//someaggregator := FakeAggregator{}
	someaggregator := new(ag1.SumAggregator)
	result, cnt := Simple(nums, someaggregator)

	if result != 15 || cnt != 5 {
		t.Errorf("Aggregator returned res-%d, cnt-%d", result, cnt)
	}
}

func TestPartition(t *testing.T) {
	t.Skip("Skipping it for now")
}