package main

import (
	"fmt"
	"math/rand"
	//"time"
)

import ag1 "github.com/chinmayb/gotraining/aggregator/aggregators"
import ao "github.com/chinmayb/gotraining/aggregator/aggregops"

func generate(elecount int) []int {
	nums := make([]int, elecount)
	for i := 0; i < elecount; i++ {
		nums[i] = rand.Intn(100)
	}
	return nums
}

func timeTrack() {

}

func callSimple(casename string, nums []int, ag ao.Aggregator) {
	//defer timeTrack(time.Now(), casename)
	result, count := ao.Simple(nums, ag)

	fmt.Println(result, count)

}
func main() {
	summer := new(ag1.SumAggregator)
	//numbers := [] int {10,12,14,16}
	nums := generate(1000)
	callSimple("case :1 Simple", nums, summer)

}
