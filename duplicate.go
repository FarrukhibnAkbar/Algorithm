package main

import "fmt"

func duplicate(array []int) bool{
	var ok bool

	for i := 0; i < len(array)-1; i++ {
		for j :=i+1; j < len(array); j++ {
		
			if array[i] == array[j]{
				ok = true
			}
		}
	}

	return ok
}

func main() {

	s := []int{1, 3, 1, 5, 4, 2}
	ok := duplicate(s)

	if ok {
		fmt.Println("Duplicate has element in array")
	} else {
		fmt.Println("No duplicate")
	}
}