package main 

import "fmt"

func main() {
		

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("a == %d\n", a)

	newA := make([]int, len(a), (cap(a)*2))

	for i := range a {
		newA[i] = a[i]
		
	}

}