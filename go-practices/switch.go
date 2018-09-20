package main

import "fmt"
import "runtime"

func main(){

	var os string 
	var arch string
	os=runtime.GOOS
	arch=runtime.GOARCH

	fmt.Printf("Go runs on ")
	switch(os){
		case "linux":
			fmt.Printf("Linux")
		case "darwin":
			fmt.Printf("OS X")
		default:
			fmt.Printf(os)
	}
	fmt.Println("", arch)
}