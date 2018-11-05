package main

import (
	"fmt"
	"time"
)

func main(){

	now:=time.Now()
	then:=time.Date(2008, time.September, 13, 17, 0, 0, 0, time.FixedZone("pl", 3600))

	diff:=now.Sub(then)

	fmt.Printf("%v\n", diff)

}