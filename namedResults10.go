package main

import "fmt"

func doSomething(myNum int) (x, y int) {
    x = myNum / 3
    y = myNum - x
    return
}

func main() {
    fmt.Println(doSomething(100))
}