package main

import "fmt"

func main(){

	// origin empty-slice
	var s []int
	printSlice(s)

	s=append(s, 1)
	printSlice(s)

	s=append(s, 2)
	printSlice(s)	

	s=append(s, 3)
	printSlice(s)		

	s=append(s, 4, 5, 6, 7, 8)
	printSlice(s)	
}

func printSlice(s []int){
	fmt.Printf("cap=%v, len=%v, [%v]\n", cap(s), len(s), s)
}