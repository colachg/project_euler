package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	sum := 0
	for i := 0; fib(i) <= 4000000; i++ {
		if fib(i)%2 == 0 {
			sum += fib(i)
		}
	}
	fmt.Println("sum is ===>", sum)
}
