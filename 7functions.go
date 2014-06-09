package main

import "fmt"

func add(x int, y int) int {  //type comes after the var name and then it RETURNS an int
    return x + y
}

func multiply(x int, y int ) int {
	return ((x + add(5, 2)) * (y + add(3,1))) //2+7 * 2+4 
}

func main() {
    fmt.Println(add(42, 13))
    fmt.Println(multiply(2, 2)) //functions similar to Cish
}