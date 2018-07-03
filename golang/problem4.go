//Largest palindrome product

//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

//Find the largest palindrome made from the product of two 3-digit numbers.

package main

import "fmt"

func main() {
	var n int
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			if i*j < 100000 {
				if i*j/10000 == i*j%10 && (i*j/1000)%10 == (i*j/10)%10 {
					fmt.Println(i * j)
				}
			} else {
				if i*j/100000 == i*j%10 && (i*j/10000)%10 == (i*j/10)%10 && (i*j/1000)%10 == (i*j/100)%10 {
					if i*j > n {
						n = i * j
					}
				}
			}
		}
	}
	fmt.Println("The largest palindrome is ===>", n)
}
