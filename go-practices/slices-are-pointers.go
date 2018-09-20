package main

import "fmt"

type Car struct{
	name string
	status string
}

func main(){
	var t =[3]Car{
			{"Citroen", "ok"},
			{"Volvo", "ok"},
			{"Peugeot", "ok"},
		}

	//pointer to 2nd array element
	var pointers []Car = t[0:3]

	// this will work the same:
	// pointers:=t[:]

	fmt.Println("Cars: ", t)
	fmt.Println("---- Volvo has got broken ----")
	pointers[1].status="broken"

	fmt.Println("Cars: ", t)

	fmt.Println("---- Let's repair volvo ----")
	var volvo[]Car=t[1:2]
	volvo[0].status="repaired"
	fmt.Println("Cars: ", t)

	fmt.Println("volvo slice", volvo)
	fmt.Println("volvo slice len", len(volvo))
	fmt.Println("volvo slice capacity", cap(volvo))
	fmt.Println("---- Let's extend volvo slice ----")
	
	volvo=volvo[:2]
	fmt.Println("volvo extended slice", volvo)
	fmt.Println("volvo extended slice len", len(volvo))
	fmt.Println("volvo extended slice capacity", cap(volvo))


}