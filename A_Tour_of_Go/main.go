package main

import (
	"fmt"
)

// UpperCamelCase
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

var c, python, java bool = true, true, false

func main() {
	fmt.Println("3 + 4 =", add(3, 4))
	fmt.Println(swap("world", "hello"))
	var i, j int = 1, 2
	k := 3
	fmt.Println(i, j, k, c, python, java)

}
