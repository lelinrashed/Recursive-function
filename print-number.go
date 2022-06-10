package main

import "fmt"

func printNumber(n int) {
	if n == 0 {
		return
	}
	printNumber(n - 1)
	fmt.Println(n)
}
