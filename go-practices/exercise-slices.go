package main

import "golang.org/x/tour/pic"

func Pic(dx int, dy int) [][]uint8 {
	picture:=make([][]uint8, dx)

	for x:=range picture{
		picture[x]=make([]uint8, dy)

		for y:=range picture[x]{
			picture[x][y]=uint8((x+y)/2)
		}
	}

	return picture
}


func main(){
	pic.Show(Pic)
}