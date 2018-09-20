package main

import "fmt"

func main(){

	auta := []string{"Citroen", "Peugeot", "Volvo", "Fiat", "Mercedes"}

	for i,v := range auta{
		fmt.Printf("[%d]\t %v\n", i, v)
	}

	fmt.Println(auta)
}