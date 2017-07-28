package main

import "fmt"

import "github.com/chinmayb/gotraining/first"
import "github.com/chinmayb/gotraining/counter"

func counterFunc() {
	fmt.Printf("\n[-] Counter %d", counter.Next())
	fmt.Printf("\n[-] Counter %d", counter.Next())
	fmt.Printf("\n[-] Counter %d", counter.Next())
	counter.Init()
	fmt.Printf("\n[-] Counter %d", counter.Next())
}

func main() {
	fmt.Println("\nWelcome\n")
	// fmt.Println("[-] Int and float types")
	// first.IntFloatTypes()
	// fmt.Println("\n[-] Addition", first.Add(1, 2))
	//counterFunc()
	//counter.COUNT = 5
	//counter.Show()
	//first.Arrayoperations()
	//first.SliceOperations()
	first.MadMaps()
}
