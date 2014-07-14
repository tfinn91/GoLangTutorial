package main

import (
    "github.com/etsy/statsd/examples/go"
    "fmt"
    )

func main() {

    // Create a new StatsD connection
    host := "localhost"
    port := 8125
    client := statsd.New(host, port)

    


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