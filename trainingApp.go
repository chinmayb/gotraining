package main

import "fmt"

import "github.com/chinmayb/gotraining/first"
import "github.com/chinmayb/gotraining/counter"

func main() {
	fmt.Println("\nWelcome\n")
	fmt.Println("[-] Int and float types")
	first.IntFloatTypes()
	fmt.Println("\n[-] Addition", first.Add(1,2))

	fmt.Printf("\n[-] Counter %d", counter.Next())
	fmt.Printf("\n[-] Counter %d", counter.Next())

	counter.COUNT = 5
	counter.Show()
}