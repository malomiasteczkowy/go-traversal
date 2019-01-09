package main

import(
	"fmt"
	"strconv"
	"strings"
)

func main(){
	var f float64
	var err error
	
	str:=" 0.00079100000"

	f,err = strconv.ParseFloat(strings.Trim(str, " "), 64)
	if err != nil{
		fmt.Println("Fucked up!", err)
		return
	}
	fmt.Println(str, f)
}