package main

import (
	"testing"
)

func TestMoney(t *testing.T) {
	t.Run("$5 * 2 = 10", func(t *testing.T) {
		expect := Dollar{10}
		actual := Dollar{5}.times(2)
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
		expect = Dollar{15}
		actual = Dollar{5}.times(3)
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
	})
	t.Run("equals()の実装", func(t *testing.T) {
		expect := true
		actual := Dollar{5}.equals(Dollar{5})
		if expect != actual {
			t.Errorf("expect %v, actual %v", expect, actual)
		}
		expect = false
		actual = Dollar{5}.equals(Dollar{6})
		if expect != actual {
			t.Errorf("expect %v, actual %v", expect, actual)
		}
	})
	t.Run("フランのかけ算", func(t *testing.T) {
		expect := Franc{10}
		actual := Franc{5}.times(2)
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
		expect = Franc{15}
		actual = Franc{5}.times(3)
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
	})
}
