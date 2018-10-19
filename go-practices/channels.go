package main

import "fmt"

func sum(tab []int, c chan int){
	
	var s int
	s=0
	
	for _,n := range tab{
		s = s + n
	} 
	c <- s
}

func main(){
	var pp = []int{0, 1, 2, 3, 4, 5, 6, 7}

	c := make(chan int)

	go sum(pp[:len(pp)/2], c)
	go sum(pp[len(pp)/2:], c)
	x := <- c
	y := <- c

	fmt.Println(x, y, x+y)
}
