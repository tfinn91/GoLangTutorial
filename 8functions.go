package main

import "fmt"

func add(x, y, z, a, b, c int) int {  //just another way to declare them all as int
    return x + y
}

func main() {
    fmt.Println(add(42, 13, 2, 3, 6, 9))
}