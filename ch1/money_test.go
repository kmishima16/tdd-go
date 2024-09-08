package main

import "testing"

func TestMultiplication(t *testing.T) {
	dollar := Dollar{5}
	dollar.times(2)
	expect := 10
	actual := dollar.amount
	if expect != actual {
		t.Errorf("expect %d, actual %d", expect, actual)
	}
}
