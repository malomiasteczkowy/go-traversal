package main

import (
	"fmt"
	"time"
	"math/rand"
)

const TASK_COUNT int = 555

func main(){
	
	rand.Seed(time.Now().UTC().UnixNano())

	ready := make(chan string, TASK_COUNT)

	fmt.Printf("starting %d tasks\n\n", TASK_COUNT)
	// start watkow
	for n:=1; n<=TASK_COUNT; n++{
		go task(ready, n)
	}

	// przechwytujemy tylko komunikaty trzech pierwszych watkow
	for i:=1; i<=3; i++{
		fmt.Printf("%d place:\t %s\n", i, <-ready)
	}  	

}

func task(ready chan string, n int){
	time.Sleep(time.Duration(rand.Intn(30))*time.Second)
	ready <- fmt.Sprintf("task %d", n)
}