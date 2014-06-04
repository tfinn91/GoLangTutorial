package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi) //Pi is the name exported by "math". We can refer to Pi with a capital letter. math.pi will not work where as math.Pi will work
    fmt.Println(math.Sqrt(4)) //works the same way

    
}