package main

import (
	"fmt"
	"time"
	"math/rand"
)

func googleIt(ready chan string, query *string){
	time.Sleep(time.Duration(rand.Intn(10))*time.Second)
	ready <- "Google response"
}

func main(){
	// 	case <-time.After(5 * time.Second):
	//	rand.Seed(time.Now().UTC().UnixNano())
	//	time.Sleep(time.Duration(rand.Intn(10))*time.Second)
	rand.Seed(time.Now().UTC().UnixNano())

	query := "this is my query"

	ready := make(chan string, 1)

	go googleIt(ready, &query)

	select{
	case response:= <-ready: {
		fmt.Println(response)
	}
	case <-time.After(5 * time.Second):{
		fmt.Println("Timeout")
	}

	}
}