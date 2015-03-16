package pascal

func Triangle(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	prevTriangle := Triangle(n - 1)

	lastRow := prevTriangle[len(prevTriangle)-1]
	newRow := make([]int, len(lastRow)+1)

	newRow[0] = 1
	for i := 1; i < len(lastRow); i++ {
		newRow[i] = lastRow[i-1] + lastRow[i]
	}
	newRow[len(newRow)-1] = 1

	return append(prevTriangle, newRow)
}
