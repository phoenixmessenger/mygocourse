package main

import "log"

type Math struct {
	A float64
	B float64
}

func (m *Math) Add() float64 {
	c := m.A + m.B
	return c
}
func (p *Math) Product() float64 {
	d := p.A * p.B
	return d
}
func (m *Math) Minus() float64 {
	c := m.A - m.B
	return c
}

func (m *Math) Divide() float64 {
	c := m.A / m.B
	return c
}
func main() {
	var m = new(Math)
	m.A = 64.8
	m.B = 0
	r := m.Add() //r := Add(m)
	h := m.Product()
	e := m.Minus()
	f := m.Divide()
	log.Println(r)
	log.Println(h)
	log.Println(e)
	log.Println(f)
}
