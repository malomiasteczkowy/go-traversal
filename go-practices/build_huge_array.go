package main

import "fmt"

const HUGESIZE = 110000000

func buildHugeArray(size uint64) []string{
	var hugeArray [HUGESIZE]string
	var wrk string

	for i:=0; i<HUGESIZE; i++{
		wrk=fmt.Sprintf("import;co_id;%d", i)
		hugeArray[i]=wrk
	}

	return hugeArray[:]
}

func main(){
	hugeArray := buildHugeArray(HUGESIZE)

	fmt.Println(hugeArray[123])

	fmt.Println(len(hugeArray))
}