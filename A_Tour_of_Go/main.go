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

func nsum(n int, c chan<- int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	c <- sum
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

// Vertex
type Vertex struct {
	X, Y float64
}

// Abs
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var i interface{} = "hello"

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	switch v := i.(type) {
	case string:
		fmt.Println(v, "string!!")
	case float64:
		fmt.Println(v, "float!!")
	default:
		fmt.Println(v, "other!!")
	}

	c := make(chan int)
	go nsum(100, c)
	res := <-c
	fmt.Println(res)
}
