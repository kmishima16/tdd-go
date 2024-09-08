package main

type Dollar struct {
	amount int
}

func (d Dollar) times(n int) Dollar {
	return Dollar{d.amount * n}
}
