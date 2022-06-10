package main

func sumBetween(n int, m int) int {
	if n > m {
		return 0
	}

	// 2 + sumBetween(2+1, 8)	35
	// 3 + sumBetween(3+1, 8)	33
	// 4 + sumBetween(4+1, 8)	30
	// 5 + sumBetween(5+1, 8)	26
	// 6 + sumBetween(6+1, 8)	21
	// 7 + sumBetween(7+1, 8)	15
	// 8 + sumBetween(0)		 8

	return n + sumBetween(n+1, m)
}
