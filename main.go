package main

import "fmt"

func main() {
	a := 10
	b := 20
	c, d := func(a, b int) (int, int) {
		res := a + b      // first value
		resII := a * b    // second value
		return res, resII // return two value
	}(a, b)
	///123456

	fmt.Println("multiplication:", d)
	fmt.Println("sum:", c)
}
