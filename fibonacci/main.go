package main

import (
	"fmt"
)

var fibMap = make(map[int]int)

func main() {
	fmt.Println(fib(100))
}

func fib(n int) int {
	if val, ok := fibMap[n]; ok {
		return val
	}
	if n <= 2 {
		return 1
	}

	fibMap[n] = fib(n-1) + fib(n-2)

	return fibMap[n]
}
