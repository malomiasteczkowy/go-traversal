package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	x float64
	y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) Scale(f float64){
	v.x = v.x * f
	v.y = v.y * f
}

func main(){
	v := Vertex{3, 4}

	fmt.Printf("original Abs: %v\n", v.Abs())
	(&v).Scale(0.5)
	fmt.Printf("scaled 1:2 Abs: %v\n", v.Abs())

}