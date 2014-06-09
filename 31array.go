package main

import "fmt"

func main() {
    var a[2] string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[1], a[0])					//print string in reverse order...just showing how to print


	for i := 0; i < 2; i++ {
	fmt.Printf(a[i])						//print array with spaces between indices 
	fmt.Printf(" ")
	}


	fmt.Printf("\n")

	fmt.Println(a)							//prints string with square brackets
}