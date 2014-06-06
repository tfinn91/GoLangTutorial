package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{5, 9}
    fmt.Println(Vertex{5, 9})

    v.X = 99
    v.Y = 350000

    fmt.Printf("The struct now reads: %v, %v\n", v.X, v.Y)
    fmt.Println(Vertex{v.X, v.Y})
}