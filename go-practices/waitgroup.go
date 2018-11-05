package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

/* WaitGroup is a kind of a semaphore in its behaviour */

func main(){

	var wg sync.WaitGroup
	respond := make(chan string, 5)

	// set semaphore to 5
	wg.Add(5)

	fmt.Println("starting tasks")
	// run 5 asynchronous tasks; each task will substract 1 from semaphore
	go task(1, &wg, respond)
	go task(2, &wg, respond)
	go task(3, &wg, respond)
	go task(4, &wg, respond)
	go task(5, &wg, respond)				

	// waits until semaphore is zeroed
	wg.Wait()

	// no more values will be placed int channel
	// but only sender should close the channel !!!
	// ... but 'range' operator reads from channel untils the channel IS CLOSED
	// so 'close' is needed here anyway
	close(respond)
	for queryResponse := range respond {
		fmt.Printf("got response:\t %s\n", queryResponse)
	}

}

func task(n int, wg *sync.WaitGroup, respond chan string){
	defer wg.Done()	// substracts 1 from the semaphore wg

	time.Sleep(time.Duration(rand.Intn(10))*time.Second)
	respond <- fmt.Sprintf("task %d is done", n)
}