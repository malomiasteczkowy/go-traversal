package main

import (
	"fmt"
)

type DaneTechniczne struct{
	PojSilnika int
	Moc int
	TypSilnika string
}

type Samochod struct{
	Marka string
	Model string
	Dane DaneTechniczne
}

func New(marka string, model string) (*Samochod){
	return &Samochod{Marka: marka, Model: model}
}

func main(){
	s:= New("Volvo","V40")
	fmt.Println(s)

	s.Dane.Moc = 90
	s.Dane.PojSilnika = 1900
	s.Dane.TypSilnika = "F8Q4"
	fmt.Println(s)
}