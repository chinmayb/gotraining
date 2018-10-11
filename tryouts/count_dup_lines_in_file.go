package main

import "os"
import "bufio"

func countLines(f *os.File, filename string, count map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

	}
}

func main() {
	files := os.Args[1:]
	count := map[string]int{}
	for _, file := range files {
		f, _ := os.Open(file)
		countLines(f, file, count)
		defer f.Close()
	}
}
