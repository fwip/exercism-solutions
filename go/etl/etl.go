package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for score, letters := range in {
		for _, l := range letters {
			out[strings.ToLower(l)] = score
		}
	}
	return out
}
