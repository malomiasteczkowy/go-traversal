package main

import "fmt"

type Rectangle struct {
	width int
	height int
}

func area1(rectangleObj Rectangle) int{
	return rectangleObj.width*rectangleObj.height;
}

func area2(pRectangle *Rectangle) int{
	return pRectangle.width*pRectangle.height;
}

func main(){
	//rectangleObj := rectangle{10,20}
	var rectangle Rectangle
	rectangle.width=10
	rectangle.height=30

	rectangleArea := area1(rectangle)
	fmt.Println("Rectangle area:", rectangleArea)

	rectangleArea=0
	
	rectangleArea = area2(&rectangle)
	fmt.Println("Rectangle area:", rectangleArea)

	fmt.Printf("What type is it? [%T]\n", &rectangle)
}