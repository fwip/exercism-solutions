package queenattack

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	posRegex     = regexp.MustCompile("^(?P<row>[a-h])(?P<col>[1-8])$")
	letterDecode = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}
)

type pos struct {
	row, col int
}

func parseAlgebraic(in string) (out pos, err error) {
	match := posRegex.FindStringSubmatch(in)
	if match == nil {
		err = errors.New("Not valid syntax")
		return
	}
	out.row = letterDecode[match[1]]
	out.col, _ = strconv.Atoi(match[2])
	return
}

func CanQueenAttack(w, b string) (canAttack bool, err error) {

	wPos, wErr := parseAlgebraic(w)
	if wErr != nil {
		err = errors.New("Couldn't parse white queen position: " + w)
		return
	}
	bPos, bErr := parseAlgebraic(b)
	if bErr != nil {
		err = errors.New("Couldn't parse black queen position: " + b)
		return
	}

	if wPos == bPos {
		err = errors.New("Can't have both queens in the same spot.")
		return
	}

	canAttack = canSee(wPos, bPos)
	return canAttack, err
}

func canSee(w, b pos) bool {
	xDiff := w.row - b.row
	yDiff := w.col - b.col

	return xDiff == 0 || yDiff == 0 || xDiff == yDiff || xDiff == -yDiff
}
