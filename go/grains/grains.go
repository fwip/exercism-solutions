package grains

import "errors"

func Square(n int) (uint64, error) {
	if n > 64 || n < 1 {
		return 0, errors.New("There's 64 squares on a chessboard.")
	}
	return 1 << uint(n-1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
