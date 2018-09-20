package main

import "fmt"

func main(){

	var t[10] string

	t[1]="lipa"
	t[2]="akacja"

	fmt.Println(t[1])
	fmt.Println(t[2])
	fmt.Println(t)

	var trasa = []string {"Warszawa", "Przemysl", "Lwow", "Kolomyja"}
	//lub
	//trasa:=[]string {"Warszawa", "Przemysl", "Lwow", "Kolomyja"}

	fmt.Println(trasa[3])
	fmt.Println(trasa)

}