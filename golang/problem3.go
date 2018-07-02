//Largest prime factor
//Problem 3 
//The prime factors of 13195 are 5, 7, 13 and 29.

//What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

func main() {
	factor, n := 2, 600851475143
	for n > 1 {
		if n%factor == 0 {
			n = n / factor
		}
		for n%factor == 0 {
			n = n / factor
		}
		factor += 1
	}
	fmt.Println("The largest prime factor is ==>", factor-1)
}
