package fizzbuzz

import (
	"strconv"
)

var (
	fizz = "Fizz"
	buzz = "Buzz"
)

// Convert FizzBuzzのルールに合わせて値を変換する
func Convert(i int) (string, error) {
	// 3と5の最小公倍数
	if 0 == (i % 15) {
		return fizz + " " + buzz, nil
	}

	if 0 == (i % 3) {
		return fizz, nil
	}

	if 0 == (i % 5) {
		return buzz, nil
	}

	return strconv.Itoa(i), nil
}
