package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("starting")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for key, value := range counts {
		fmt.Println("%s - %d", key, value)
	}
}