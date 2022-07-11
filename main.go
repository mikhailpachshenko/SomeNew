package main

import "fmt"

func main() {
	a := 10
	b := 20
	c, d := func(a, b int) (int, int) {
		res := a + b
		resII := a * b
		return res, resII
	}(a, b)

	fmt.Println("multiplication:", d)
	fmt.Println("sum:", c)
}
