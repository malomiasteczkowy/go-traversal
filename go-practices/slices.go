package main

import "fmt"

func main(){
	var primes = [5]int{2, 3, 5, 7, 11}
	var slice []int

	slice=primes[1:4]

	fmt.Println("array (fixed width): ", primes)
	fmt.Println("slice (dynamic width): ", slice)
}