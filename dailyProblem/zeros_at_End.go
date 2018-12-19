package main

import "fmt"

//  O(n) and extra space is O(1)
//Input :  arr[] = {1, 2, 0, 4, 3, 0, 5, 0};
//Output : arr[] = {1, 2, 4, 3, 5, 0, 0};

func main(){
	count := 0
	arr := []int{1, 2, 0, 0, 3, 0, 5, 0}
	for i, a := range arr {
		if a > 0 {
			arr[count] = arr[i]
			count ++
		}
	}
	for count < len(arr) {
		arr[count] = 0
		count ++
	}
	fmt.Println(arr)
}
