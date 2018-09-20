package main

import "fmt"

func main(){
	fmt.Println("Calling function f()")
	f()
	fmt.Println("Program finished normally")
}

func f() {
	fmt.Println("Function f() calling function g()")
	defer func (){
		fmt.Println("Function f() calling defered function")
		if r:=recover(); r!=nil{
			fmt.Println("Function f() recovered from panic: ", r)
		}
	}()
	g(0)
	fmt.Println("Function f() finished normally")

}

func g(i int){
	if i>3{
		fmt.Println("Function g() calling panic")
		panic(fmt.Sprintf("argument %v caused panic", i))
		return
	}
	fmt.Println("Function g() called with arg ", i)
	defer fmt.Println("Function g() release defered statement with arg ", i)
	g(i+1)
}