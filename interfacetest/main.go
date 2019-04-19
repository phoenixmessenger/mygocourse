package main

import "fmt"

func myPrint(i interface{}) {

	vint, ok := i.(int)
	if ok {
		fmt.Printf("%d\n", vint)
		return
	}

	vfloat64, ok := i.(float64)
	if ok {
		fmt.Printf("%f\n", vfloat64)
		return
	}

	vstr, ok := i.(string)
	if ok {
		fmt.Printf("%s\n", vstr)
		return
	}
}
