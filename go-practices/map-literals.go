package main

import "fmt"

type Person struct {
	Age int
	Sex string
}

var m = map[string]Person{
	"Ann": 
		Person{25, "female",},
	"Tom":
		Person{40, "male",},
	"Old Jim":
		Person{98, "male",},
}

func main(){
	fmt.Println(m)
	m["Kate"]=Person{35, "female"}
	fmt.Println(m)
}