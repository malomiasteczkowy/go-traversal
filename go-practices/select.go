package main

import "fmt"

func fibonacci(c, quit chan int){
	x, y := 0, 1

	for {
		select{
		case c<-x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("bye")
			return	
		}

	}
	close(c)
}

func main(){
	c := make(chan int, 10)
	quit := make(chan int)
	
	go func(){
		for i:=0; i<10; i++{
			fmt.Println(i)
		}
		quit<-0
	}()
	
	fibonacci(c, quit)
}