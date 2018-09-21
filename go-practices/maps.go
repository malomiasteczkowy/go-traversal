package main

import "fmt"

type Vertex struct{
	Lat float64
	Long float64
}

var m map[string]Vertex

func main(){

	// make returns map of the given type
	// the map is at least 2-elements size
	// more elements will automaticaly invoke map resizing
	m = make(map[string]Vertex, 2)

	m["Mount Everest"]=Vertex{27.5917, 86.5531}
	m["Warsaw"]=Vertex{52.14, 21.1}

	fmt.Println(m["Mount Everest"])
	fmt.Println(m["Warsaw"])


	// third element will cause map resizing
	m["Katmandu"]=Vertex{27.42, 85.20}
	fmt.Println("---")

	fmt.Println(m["Mount Everest"])
	fmt.Println(m["Warsaw"])
	fmt.Println(m["Katmandu"])


}