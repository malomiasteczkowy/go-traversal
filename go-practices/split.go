package main

import "fmt"
import "strings"

func main(){
  s:= "Ala;Mateusz;Karol;Leszek;"
  z:=strings.Split(s, ";");
  for i,v := range z {
    fmt.Printf("%d\t%s\n", i, v)
  }
  fmt.Printf("[%d]\n", len(z))
}
