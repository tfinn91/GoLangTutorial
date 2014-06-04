package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {

	

	a := "hello"
	b := "world"


	x, y := swap(a, b)
    fmt.Println(x, y)
}