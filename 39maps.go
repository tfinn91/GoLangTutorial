package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{ 40.68433, -74.39967,}
    m["newKey"] = Vertex{ 45, 54.34235,}


	for key, value := range m {
    	fmt.Println("Key:", key, ", Value:", value )
	}

	fmt.Printf("m has a length of: %d\n\n", len(m))

	delete(m, "Bell Labs")

	fmt.Println("After deleting \"Bell Labs\" the new map has:")

	for key, value := range m {
    	fmt.Println("Key:", key, ", Value:", value)
	}

	fmt.Printf("m NOW has a length of: %d\n", len(m))

	v, ok := m["Bell Labs"]
    fmt.Println("Checking Bell Labs value:", v, "Present?", ok)   //this checks to see if the value of "Bell Labs" is present
    

	}