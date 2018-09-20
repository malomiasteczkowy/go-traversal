package main

import "fmt"

func main(){
	a:=make([]int, 5)
	printSlice("a", a)

	b:=make([]int, 0, 5)
	printSlice("b", b)

	c:=b[:cap(b)]
	printSlice("c", c)

	d:=c[2:5]
	printSlice("d", d)
}

func printSlice(s string, slice []int){
	fmt.Printf("%s: cap %v, len %v\n", s, cap(slice), len(slice))
}