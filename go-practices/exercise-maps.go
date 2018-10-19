package main

//import "fmt"
import "strings"
import "golang.org/x/tour/wc"

const TotalWords int = 1000

func WordCount(s string) map[string]int{

	words:=strings.Fields(s)
	wordMap:=make(map[string]int)

	for _,v := range words {
		n:=wordMap[v]
		wordMap[v]=n+1
	}

	return wordMap

}

func main(){

	wc.Test(WordCount)
/*
	s:= `
		W górze tyle gwiazd
		W dole tyle miast.
		Gwiazdy miastom dają znać,
		Że dzieci muszą spać!

		Ach śpij, kochanie.
		Jeśli gwiazdkę z nieba chcesz - dostaniesz.
		Czego pragniesz, daj mi znać.
		Ja Ci wszystko mogę dać.
		Więc dlaczego nie chcesz spać?


		Ach śpij, bo właśnie
		Księżyc ziewa i za chwilę zaśnie.
		A gdy rano przyjdzie świt,
		Księżycowi będzie wstyd,
		Że on zasnął, a nie Ty.

		Aaaa, aaaaa
		Były sobie kotki dwa.
		Aaaa, aaaa
		Szaro-bure, szaro-bure obydwa

		Ach śpij, bo nocą,
		Kiedy gwiazdy się na niebie złocą,
		Wszystkie dzieci, nawet złe,
		Pogrążone są we śnie,
		A Ty jeden (jedna) tylko nie.
	`

	dummy:=WordCount(s)

	for key,value:=range dummy{
		fmt.Printf("[%d] %s\n", value, key)
	}
	*/
}