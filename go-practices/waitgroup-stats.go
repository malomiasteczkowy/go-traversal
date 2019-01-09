package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
	"sync/atomic"
)

/* WaitGroup is a kind of a semaphore in its behaviour */

/* should use sync/atomic */

const MAXRETCODE int = 1000
const THREADCOUNT int = 100

type Stat struct{
	in	uint64
	retcodes [MAXRETCODE]uint64
	other uint64
}

/* thread unsafe */
func printStat(s *Stat){
	total:=s.in
	fmt.Printf("Total: %d\n", total)
	for i,v := range s.retcodes{
		if v>0{
			fmt.Printf("[%d]: %2d\n", i, v)	
		}
	}
	fmt.Printf("Other: %d\n", s.other)
}

/* thread safe - using sync/atomic */
func updateStat(s *Stat, inRetCode int){
	atomic.AddUint64(&(s.in), 1)
	if inRetCode < MAXRETCODE {
		atomic.AddUint64(&(s.retcodes[inRetCode]), 1)
	} else {
		atomic.AddUint64(&(s.other), 1)
	}
}

func main(){

	var wg sync.WaitGroup
	respond := make(chan string, THREADCOUNT)

	/* Stat init */
	stat := new(Stat)
	
	fmt.Printf("Starting %d tasks\n", THREADCOUNT)

	// set semaphore to x
	wg.Add(THREADCOUNT)
	
	// run x asynchronous tasks; each task will substract 1 from semaphore
	for i:=0;i<THREADCOUNT;i++{
		go task(i, &wg, respond, stat)	
	}

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

	printStat(stat)

}

func task(n int, wg *sync.WaitGroup, respond chan string, stat *Stat){
	defer wg.Done()	// substracts 1 from the semaphore wg

	time.Sleep(time.Duration(rand.Intn(10))*time.Second)
	updateStat(stat, rand.Intn(10))

	respond <- fmt.Sprintf("task %d is done", n)
}