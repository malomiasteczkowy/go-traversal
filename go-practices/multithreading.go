package main

import "fmt"

func ogle(x *int, y *int, ready chan bool) int{

	*y += *x

	ready <- true

	return 0
}

func main(){
	a := 1
	b := 2

	ready := make(chan bool)

	go ogle(&a, &b, ready)

	<-ready
	
	fmt.Println(a, b)
	
}