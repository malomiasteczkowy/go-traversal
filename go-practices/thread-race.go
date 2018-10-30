package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){

	rand.Seed(time.Now().UTC().UnixNano())

	query := "This is my query"

	// buffered channel
	ready := make(chan string, 2)

	go googleIt(query, ready)
	go bingIt(query, ready)

	response:=  <-ready

	fmt.Println(response)
}

func googleIt(query string, ready chan string){
	// sleep
	time.Sleep(time.Duration(rand.Intn(10))*time.Second)
	ready <- "Google response"
}

func bingIt(query string, ready chan string){
	// sleep
	time.Sleep(time.Duration(rand.Intn(10))*time.Second)
	ready <- "Bing response"
}