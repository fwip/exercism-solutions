package diffsquares

func SquareOfSums(n int) int {
	// For a sequence of numbers 1..N: seq[x] + seq[N-x] = N+1
	sums := ((n + 1) * n) / 2
	return sums * sums
}

func SumOfSquares(n int) (total int) {
	for i := 1; i <= n; i++ {
		total += i * i
	}
	return
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
