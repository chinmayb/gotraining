package main

import (
	"bufio"
	"fmt"
	"os"
)
import ops "github.com/chinmayb/gotraining/matrix/operations"

const (
	ADD    = "add"
	CREATE = "create"
	SUB    = "subtract"
	SHOW   = "show"
	EXIT   = "exit"
)

func main() {
	//matrices := make(ops.Matrices, 0) //runtime call
	matrices := ops.Matrices{} //compile
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nEnter the operation to perform : [ ",
			CREATE, SUB, ADD, SHOW, EXIT, " ]")
		oper, _ := reader.ReadString('\n')
		oper = oper[:len(oper)-1]
		switch oper {
		case CREATE:
			m1 := ops.New(3, 3)
			m1.Insert()
			matrices = append(matrices, m1)
		case SHOW:
			for _, m := range matrices {
				m.Show()
				//fmt.Println(m.ToString())
			}
		case ADD:
			fmt.Println("\n Enter the ids of  first matrix to add")
			showid1, _ := reader.ReadString('\n')
			showid1 = showid1[:len(showid1)-1]
			fmt.Println("\n Enter the ids of  second matrix to add")
			showid2, _ := reader.ReadString('\n')
			showid2 = showid2[:len(showid1)-1]

			matrices = append(matrices, ops.Add(matrices[0], matrices[1]))
		case SUB:
			matrices = append(matrices, ops.Subtract(matrices[0], matrices[1]))
		case EXIT:
			os.Exit(0)
		default:
			fmt.Printf("Invlid choice")
		}

	}
}
