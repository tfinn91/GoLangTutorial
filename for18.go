package main

import "fmt"

func main() {



    sum := 0
    for i := 0; i < 10; i++ {


		fmt.Printf("Sum = %v + %v", sum, i)  

        sum += i

        fmt.Printf("=====%v\n", sum)
    }
    fmt.Printf("The final sum is: %v\n", sum)


    for i := 0; i < 10; i++ {
    	
    	fmt.Println(i)
    }

}