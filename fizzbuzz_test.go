// package main
package fizzbuzz

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Run("条件に一致しない場合にStringに変換された数字を返すか", func(t *testing.T) {
		num := 1
		ret, _ := Convert(num)
		assert.Equal(t, strconv.Itoa(num), ret)

		num = 112
		ret, _ = Convert(num)
		assert.Equal(t, strconv.Itoa(num), ret)
	})

	t.Run("3の倍数のときに'Fizz'と返すか", func(t *testing.T) {
		num := 3
		ret, _ := Convert(num)
		assert.Equal(t, fizz, ret)

		num = 333
		ret, _ = Convert(num)
		assert.Equal(t, fizz, ret)
	})

	t.Run("5の倍数のときに'Buzz'と返すか", func(t *testing.T) {
		num := 5
		ret, _ := Convert(num)
		assert.Equal(t, buzz, ret)

		num = 55
		ret, _ = Convert(num)
		assert.Equal(t, buzz, ret)
	})

	t.Run("3と5の倍数のときに'Fizz Buzz'と返すか", func(t *testing.T) {
		expect := fizz + " " + buzz

		num := 15
		ret, _ := Convert(num)
		assert.Equal(t, expect, ret)

		num = 1515
		ret, _ = Convert(num)
		assert.Equal(t, expect, ret)
	})
}
