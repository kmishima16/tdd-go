package main

type Dollar struct {
	amount int
}

func (d Dollar) times(n int) Dollar {
	return Dollar{d.amount * n}
}

func (d Dollar) equals(obj Dollar) bool {
	return d.amount == obj.amount
}

type Franc struct {
	amount int
}

func (d Franc) times(n int) Franc {
	return Franc{d.amount * n}
}

func (d Franc) equals(obj Franc) bool {
	return d.amount == obj.amount
}
