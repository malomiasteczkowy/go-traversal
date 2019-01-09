package main

import (
	"fmt"
	"math/rand"
)

const ACTION_BUFFER_SIZE = 1000000

var g_poz int

func push(value int, buffer []int, poz *int) int{
	if *poz >= ACTION_BUFFER_SIZE-1 {
		return -1
	}
	if *poz < 0 {
		return -1
	}	

	buffer[*poz]=value
	fmt.Println("pushed ", value)
	*poz++

	return 0
}

func pull(buffer []int, poz *int) (int, int){
	if *poz > ACTION_BUFFER_SIZE {
		return -1, -1
	}
	if *poz <= 0 {
		return -1, -1
	}	

	*poz--
	return buffer[*poz], 0
}

func main(){
	var buffer [ACTION_BUFFER_SIZE]int
	poz := 0

	for {
		ret:=push(rand.Intn(100), buffer[:], &poz)
		if ret!=0{
			break
		}
	}

	for {
		v,ret:=pull(buffer[:], &poz)
		if ret==-1 {
			break
		}
		fmt.Println("pulled ", v)
	}
	
}