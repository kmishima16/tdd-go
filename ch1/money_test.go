package main

import "testing"

func TestMultiplicationtA(t *testing.T) {
	t.Run("$5 * 2 = 10", func(t *testing.T) {
		dollar := Dollar{5}
		product := dollar.times(2)
		expect := 10
		actual := product.amount
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
		product = dollar.times(3)
		expect = 15
		actual = product.amount
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
	})
	t.Run("equals()の実装", func(t *testing.T) {
		expect := true
		actual := Dollar{5}.equals(Dollar{5})
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
	})
}
