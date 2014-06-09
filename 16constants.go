package main

import "fmt"

const Pi = 3.14                              //automatically gets type

func main() {
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true						//again automatically gets type. damn cool
    fmt.Println("Go rules?", Truth)
}