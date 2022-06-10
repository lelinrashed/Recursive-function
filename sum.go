package main

func sum(n int) int {
	if n == 0 {
		return 0
	}

	// 10 + sum(9)55
	// 9 + sum(8) 45
	// 8 + sum(7) 36
	// 7 + sum(6) 28
	// 6 + sum(5) 21
	// 5 + sum(4) 15
	// 4 + sum(3) 10
	// 3 + sum(2) 6
	// 2 + sum(1) 3
	// 1 + sum(0) 1

	return n + sum(n-1)
}
