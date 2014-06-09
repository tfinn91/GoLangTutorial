package main

import "fmt"

var i, j = 1, 2
var c, python, java = true, false, "no!"  //initializer is present so no type needs to be specified.
var integer = 1
var floatnum = 3.414542343
var this_is_a_string = "hello world"


func main() {
    fmt.Println(i, j, c, python, java)
    fmt.Println(integer, floatnum, this_is_a_string)
    fmt.Printf("This is another string containing the num: %g\n", floatnum)
    fmt.Printf("\tPrintf statements are the exact same..: %s\n", java)
}