package main

import "fmt"

func main() {
	i := 42
	j := 3.142
	k := 0.867 + 0.5i
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("j is of type %T\n", j)
	fmt.Printf("k is of type %T\n", k)
}