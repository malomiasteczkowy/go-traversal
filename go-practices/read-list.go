package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
)

const NUM_LINES = 5

func main(){
	var err error
	var file *os.File
	var lines []string
	var n int
	
	fileName := "orca-loader.in"
	file, err = os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err!=nil{
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for {
		n, lines, err = ReadLines(scanner, NUM_LINES)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Printf("lines: %v\n", lines)
		if n<NUM_LINES{
			break
		}
	}
}

func ReadLines(scanner *bufio.Scanner, n int) (int, []string, error){

	lines := make([]string, n)
	i := 0;
	for scanner.Scan(){
		lines[i] = scanner.Text()
		i++
		if i==n{
			break
		}
	}
	err:=scanner.Err()

	return i, lines, err
}