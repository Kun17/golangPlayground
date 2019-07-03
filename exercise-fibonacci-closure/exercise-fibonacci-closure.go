package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prepre := 0
	pre := 0
	count := -1
	return func() int{
		count++
		switch count {
			case 0:
				return 0
			case 1: 
				pre = 1
				return 1
			default:
				res := prepre + pre
				prepre = pre
				pre = res
				return res
		}		
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}