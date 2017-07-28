package first
import "fmt"

func print_numbers(number int) {
	for i := number; i > 0; i-- {
		if i != 5 && i != 8 {
			fmt.Println(i)
		}
	}
}
