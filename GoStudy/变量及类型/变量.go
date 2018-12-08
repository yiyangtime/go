package main

import (
	"fmt"
)

func test() (int, string) {
	return 1, "aaaa"
}

func main() {
	var x int
	var y float32 = 1.6
	var s string = "111"
	f := 123
	fmt.Println("x=", x, "y=", y, "s=", s)
	fmt.Printf("f=%d\n", f)

	var (
		a int
		b float32
	)
	fmt.Println(a, b)

	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100
	fmt.Println(i, data[0])

	_, s = test()
	println(s)

	z := 0
	_ = z

}
