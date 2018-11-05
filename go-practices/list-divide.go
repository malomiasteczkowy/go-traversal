package main

import "fmt"

/* 	list-divide 
	dostaje liste 'n'-elementowa
	dziele ja na 'm' podlist
	do ostatniej podlisty dopisuje elementy, ktore stanowia reszte z dzielenia n / m
*/

const SIZE int = 3521				/* n */
const NUM_OF_TREADS int = 10		/* m */


func main(){
	var list [SIZE]int

	// pointers
	pList := list[:]
	
	populateList(pList)
	subLists:=divideList(pList, NUM_OF_TREADS)
	fmt.Println(subLists)

}

/* lista 'n'-elementowa zawiera kolejne liczby naturalne */
func populateList(pList []int){
	for i,_ := range pList {
		pList[i]=i	
	} 
}

func divideList(pList []int, n int) [][]int{

	var x int
	listSize := len(pList)
	subListSize := listSize/n
	subLists:=make([][]int, n)

	rest := listSize-subListSize*n

	startPoint:=0

	for x=range subLists{
		subLists[x]=make([]int, subListSize)
		copy(subLists[x][:], pList[startPoint:startPoint+subListSize])
		startPoint=startPoint+subListSize
	}

	// ostatnia sub-lista jest dluzsza o reszte z dzielenia
	for r:=0; r<rest; r++{
		subLists[x] = append(subLists[x], pList[startPoint+r])
	}

	return subLists
}