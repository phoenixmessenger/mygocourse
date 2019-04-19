package main

import "math"

type KvUravn struct {
	a, b, c float64
}

type KvResult struct { //объявляем тип KvResult
	x1, x2 float64
	isTwo  bool
	Error  bool
}

func Solution(s KvUravn) KvResult {
	var result KvResult //объявляем переменную result типа KvUravn
	var D float64
	D = math.Pow(s.b, 2) - 4*s.a*s.c
	dsqrt := math.Sqrt(D)
	if D > 0 {
		result.isTwo = true
		result.x1 = (-1 * s.b) + (dsqrt / 2 * s.a)
		result.x2 = (-1 * s.b) - (dsqrt / 2 * s.b)
	} else if D == 0 {
		result.isTwo = false
		result.x1 = (-1 * s.b) / 2 * s.a

	} else {
		result.Error = true
	}
	return result
}
