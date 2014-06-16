package main 

import (
	"fmt"
	"math/rand"
    "time"
)


func main() {
	

	rand.Seed(time.Now().Unix())

	var randNums[100] int

//filling the array
	for i := 0; i < len(randNums); i++ {
		randNums[i] = rand.Intn(100)
	}

	fmt.Printf("The array of rand nums is: \n")

	for i := 0; i < len(randNums); i++ {
		fmt.Printf("%d: %d\n", i, randNums[i])
	}


	fmt.Println("Now let's sort them!")
	var tempLow int
	var tempHi int

//sorting the array


	for i := 0; i < len(randNums)-1; i++ {
		for j := i+1; j < len(randNums); j++ {
			if randNums[i] > randNums[j] {
				tempHi = randNums[i]
				tempLow = randNums[j]

				randNums[i] = tempLow
				randNums[j] = tempHi
			}
		}

	}


//print sorted array	
	fmt.Printf("The array of sorted nums is: \n")

	for i := 0; i < len(randNums); i++ {
		fmt.Printf("%d: %d\n", i, randNums[i])
	}

}