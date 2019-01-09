package main

import "fmt"

func main(){

	empty:=""
	nonempty:="hello"

	fmt.Printf("empty's length: %d\n", len(empty))
	fmt.Printf("nonempty's length: %d\n", len(nonempty))
}