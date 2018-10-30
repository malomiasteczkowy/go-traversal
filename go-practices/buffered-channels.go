package main

import "fmt"

func main(){
	// pojemnosc 2, pomiesci dwie wartosci
	// odczyt z kanalu przywraca przetwarzanie już po pobraniu pierwszej wartosci
	// buffered channel moze byc wiec wykorzystywac do obslugi wyscigu watkow

	ch := make (chan int, 2) // pojemnosc 2, pomiesci dwie wartosci

	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}