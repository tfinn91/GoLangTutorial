package main 

import "fmt"

func checkLarger(x, y int) int {

	if x > y {
		return x
	} else if y > x {
		return y
	} else {
		return 0
	}
	
	
}


func main() {

var one, two int = 4, 9

if checkLarger(one, two) != 0 {
	fmt.Printf("The larger number is: %v\n", checkLarger(4, 9))
	
} else {
	fmt.Println("The numbers are the same!")
}


	


	
}