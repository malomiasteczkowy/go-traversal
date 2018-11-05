package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

/* WaitGroup is a kind of a semaphore in its behaviour */

/* should use sync/atomic */

type RetStat struct{
	retcode int
	count int
}

type Stat struct{
	in	int
	retcodes []*RetStat
}

func printStat(s *Stat){
	fmt.Printf("Total: %d\n", s.in)
	for _,rs := range s.retcodes{
		fmt.Printf("[%d]:%d\n", rs.retcode, rs.count)
	}
}

func updateStat(s *Stat, chs chan *Stat, inRetcode int){
	s.in++
	
	total:= <- chs.in
	chs.in <- total+1

	for _,rs := range s.retcodes{
		if rs.retcode == inRetcode{
			rs.count++
			return
		}
	}

	s.retcodes = append(s.retcodes, new(RetStat))
	last := len(s.retcodes)
	s.retcodes[last-1].retcode = inRetcode
	s.retcodes[last-1].count++
}

func main(){

	var wg sync.WaitGroup
	respond := make(chan string, 5)

	/* Stat init */
	stat := new(Stat)
	chstat := make(chan *Stat, 5)
	
	// set semaphore to 5
	wg.Add(5)

	fmt.Println("starting tasks")
	// run 5 asynchronous tasks; each task will substract 1 from semaphore
	go task(1, &wg, respond, stat, chstat)
	go task(2, &wg, respond, stat, chstat)
	go task(3, &wg, respond, stat, chstat)
	go task(4, &wg, respond, stat, chstat)
	go task(5, &wg, respond, stat, chstat)		

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

func task(n int, wg *sync.WaitGroup, respond chan string, stat *Stat, chstat chan *Stat){
	defer wg.Done()	// substracts 1 from the semaphore wg

	time.Sleep(time.Duration(rand.Intn(10))*time.Second)
	updateStat(stat, chstat, rand.Intn(10))

	respond <- fmt.Sprintf("task %d is done", n)
}