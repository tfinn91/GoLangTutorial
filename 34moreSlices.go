package main

import "fmt"



func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}





func main() {                       //Slices are created with the make function. It works by allocating a zeroed 
                                    //array and returning a slice that refers to that array: len(a)=5
    a := make([]int, 5)  
    printSlice("a", a)


    b := make([]int, 0, 5)    //set length to 0 and capacity to 5
    printSlice("b", b)

    c := [...]string{"This", "is", "a", "string"}
    fmt.Printf("c: %s and this string has a length of: %v\n\n", c, len(c))

    d := [...]string{"Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf"}
    fmt.Println(d)


    fmt.Printf("d[2:] == %s //Everything from index 2 and on\n", d[2:])
    fmt.Printf("d[:2] ==%s  //Everything up to index 2\n", d[:2])
    fmt.Printf("d[3:6] ==%s //Everything from index 3 up to index 6\n", d[3:6])



    e := d[3:]

    fmt.Printf("\ne := d[3:] == %s\n", e)
    fmt.Printf("d[0] == %s\n", d[0])
    fmt.Printf("e[0] == %s\n", e[0])




    f := e[:3] 
    fmt.Printf("\nf := e[:2] == %s\n", f)

    fmt.Printf("\n&d == %v\n", &d[3])
    fmt.Printf("&f == %v\n", &f[0])
    fmt.Println("This shows us that they share the same memory. Let's mess around with some of those variables a little...\n")


    d[5] = "NEWVAL"
    fmt.Printf("d[5] = 'NEWVAL' == %s\n", d)
    fmt.Printf("f now == %s\n", f)



    fmt.Printf("\n\nSo far we have the following: \n\td == %s\n\te == %s\n\tf == %s\n", d, e, f)


    e[0] = "Zebra"

    fmt.Printf("\ne[0] = 'Zebra' gives us the following: \n\n\td == %s\n\te == %s\n\tf == %s\n", d, e, f)

    fmt.Printf("Cool to see how changing affects the same mem location but HOLY COW if you aren't careful that could be dangerous!!!!!\n")

}


