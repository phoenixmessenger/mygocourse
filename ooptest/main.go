package main

import "log"

type Laptop interface {
	Pon() error
	Poff()
	KeyChek(k rune) error
	MouseClick(c *MC) error
	Display() []byte
}

type MC struct {
	LB         bool
	RB         bool
	MB         bool
	ScrollUp   bool
	ScrollDown bool
}

type EnovoS3 struct {
}

func (e *EnovoS3) Pon() error { //в скобках - принадлженость к объекту, а Pon() есть метод, привязанный
	log.Println("Pon EnovoS3")
	return nil
}
func (e *EnovoS3) Poff() {
}

func (e *EnovoS3) KeyChek(k rune) error {
	log.Println(k)
	return nil
}
func (e *EnovoS3) MouseClick(c *MC) error {
	return nil
}
func (e *EnovoS3) Display() []byte {
	return nil
}

func main() {
	e := &EnovoS3{} //фигурные скобки - инстанцирование объекта
	Work(e)
}

func Work(l Laptop) {
	l.Pon()
	l.Poff()
	l.MouseClick(nil)
	l.Display()
	l.KeyChek('А')
}
