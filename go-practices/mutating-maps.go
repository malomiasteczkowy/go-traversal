package main

import "fmt"


func main(){
	dzienniczek:=make(map[int]string)
	var wartosc string
	var ok bool
	const MaxId int = 10

	dzienniczek[1]="Leonard"
	dzienniczek[2]="Chester"
	dzienniczek[3]="Tom"
	dzienniczek[4]="Ray"
	dzienniczek[5]="Pam"
	dzienniczek[6]="Beath"
	dzienniczek[7]="Eddie"
	dzienniczek[8]="Kate"
	dzienniczek[9]="Agnes"
	dzienniczek[10]="Leslie"								

	fmt.Println(dzienniczek)

	fmt.Printf("[%d] %v\n", 9, dzienniczek[9])
	delete(dzienniczek, 9)

	fmt.Printf("[%d] %v\n", 9, dzienniczek[9])

	ok=true
	for i:=1; ok && i<MaxId; i++ {
		wartosc, ok = dzienniczek[i]
		if ok {
			fmt.Printf("Numer %d obecny (%s)\n", i, wartosc)
		}	
	}

}