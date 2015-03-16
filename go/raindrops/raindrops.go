package raindrops

import (
	"strconv"
)

func Convert(n int) string {
	answer := ""
	if n%3 == 0 {
		answer += "Pling"
	}
	if n%5 == 0 {
		answer += "Plang"
	}
	if n%7 == 0 {
		answer += "Plong"
	}
	if len(answer) == 0 {
		answer = strconv.Itoa(n)
	}
	return answer
}
