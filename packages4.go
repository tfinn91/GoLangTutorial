package main

import (
    "fmt"
    "math/rand"
    "time"
)



func main() {

	rand.Seed(time.Now().Unix())

	someRand := rand.Intn(10)
    fmt.Println("My favorite number is", someRand)
}