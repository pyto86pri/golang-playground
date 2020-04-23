package main

import (
	"fmt"
	"math"
	"runtime"
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

func nsum(n int) (sum int) {
	if n < 0 {
		return
	}
	for i := 0; i <= n; i++ {
		sum += i
	}
	return
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func myOs() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}

func deferHello() {
	defer fmt.Println("World!")

	fmt.Println("Hello")
}

var c, python, java bool = true, true, false

func main() {
	fmt.Println("3 + 4 =", add(3, 4))
	fmt.Println(swap("world", "hello"))
	var i, j int = 1, 2
	k := 3
	fmt.Println(i, j, k, c, python, java)
	fmt.Println(nsum(10))
	fmt.Println(nsum(-2))
	fmt.Println(pow(2, 10, 2000))
	myOs()
	deferHello()
}
