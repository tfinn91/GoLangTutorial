package main

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe   bool       = false       						//cool way to list all vars in global scope
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
    newstring string  = "hello world"
    pi 	   float64	  = 3.14159
)

func main() {
    const f = "%T(%v)\n"                      //%T = type, %v = value
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)
    fmt.Printf(f, newstring, newstring)
    fmt.Printf(f, pi, pi)



    fmt.Println("\nLet's switch up that last format: \n")

    const newf = "%v(%T)\n"

    fmt.Printf(newf, pi, pi)
    fmt.Printf(newf, newstring, newstring)
}