package main

import(
	"fmt"
)

func polindrine(a int){
	num := a
	var remainder int
   	res := 0
   	
   	for a > 0 {
      
      remainder = a % 10
      
      res = (res * 10) + remainder
      
      a = a / 10
   	}
	
	if num == res {
        
        fmt.Println("Palindrome")
   } else {
        
        fmt.Println("Not a Palindrome")
   }
}

func main() {
	var a int
	
	fmt.Println("Enter number: ")
	fmt.Scan(&a)
	
	polindrine(a)
}