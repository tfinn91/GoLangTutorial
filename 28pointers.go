package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    q := &v 							//set q to whatever v points to --> Vertex{}
    q.X = 40							//q.X is really just p.X
    v.Y = 91
    fmt.Println(*q)						//print whatever q points to
}