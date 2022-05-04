package fizzbuzz

import (
	"errors"
	"strconv"
)

var (
	fizz = "Fizz"
	buzz = "Buzz"
)

// Convert FizzBuzzのルールに合わせて値を変換する
func Convert(i int) (string, error) {
	if i <= 0 {
		return strconv.Itoa(i), errors.New("argument is only natural number")
	}

	a := (i % 3) == 0
	b := (i % 5) == 0

	if a && b {
		return fizz + " " + buzz, nil
	}

	if a {
		return fizz, nil
	}

	if b {
		return buzz, nil
	}

	return strconv.Itoa(i), nil
}
