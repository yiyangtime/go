package main

func test() (int, int) {
	return 1, 2
}

func add(x, y int) int {
	return x + y
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}

func main() {
	x, _ := test()
	println(x)

	println(add(test()))
	println(sum(test()))
}
