package main

type Laptop interface {
	Pon() error
	Poff()
	KeyChek(k rune) error
	MouseClick(c *MC) error
}

type MC struct {
	LB         bool
	RB         bool
	MB         bool
	ScrollUp   bool
	ScrollDown bool
}
