package main

import "fmt"

func fibonacci() func() int {
	count := 0
	pre1 := 0
	pre2 := 1
	return func() int {
		answer := 0
		if count == 0 {
			answer = 0
		} else if count == 1 {
			answer = 1
		} else {
			answer = pre1 + pre2
			pre1 = pre2
			pre2 = answer
		}
		count += 1
		return answer
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}