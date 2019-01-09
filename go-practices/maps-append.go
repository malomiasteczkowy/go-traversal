package main

import "fmt"

var destinations map[string][]string

func main(){

	// make returns map of the given type
	// the map is at least 2-elements size
	// more elements will automaticaly invoke map resizing
	destinations = make(map[string][]string, 1)

	destinations["Orange"]=[]string{"501", "502"}
	destinations["Plus"]=[]string{"605", "609"}

	fmt.Println(destinations["Orange"])
	fmt.Println(destinations["Plus"])
	fmt.Println(destinations["T-Mobile"])

	prefixes := destinations["Orange"]
	prefixes = append(prefixes, "503")
	destinations["Orange"]=prefixes

	prefixes = destinations["Plus"]
	prefixes = append(prefixes, "696")
	destinations["Plus"]=prefixes

	prefixes = destinations["T-Mobile"]
	prefixes = append(prefixes, "602")
	destinations["T-Mobile"]=prefixes	

	prefixes = destinations["Play"]
	prefixes = append(prefixes, "790")
	destinations["Play"]=prefixes	

	fmt.Println(destinations["Orange"])
	fmt.Println(destinations["Plus"])
	fmt.Println(destinations["T-Mobile"])
	fmt.Println(destinations["Play"])

	for i,v := range destinations{
		fmt.Println(i, v)
	}

	
}