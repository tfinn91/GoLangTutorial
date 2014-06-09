package main 

import "fmt"


func main() {
	
	var counter int                      //initializes to 0

	for counter < 100 {					//this is a while loop: NO () but YES {}

		counter++
		fmt.Printf("counter is now: %v\n", counter)
	
	}



	//for {
	//	fmt.Printf("hello")           //essentially a while(1)
	//}



}