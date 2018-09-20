package main

import "fmt"

func main(){
	var x float64=2
	var root float64
	root=square_root(x)
	fmt.Println("Square root of ", x, " is ", root)
	fmt.Println("Rounded value: ", round(root, 10))
}

func square_root(x float64) float64{
	const MaxNumberOfIterations = 100000
	var z float64 = 1.0
	var delta float64 = 1.0
	var i int

	for i<=MaxNumberOfIterations{
		i++
		fmt.Printf("krok %v: z= %v\n", i, z)
		delta=(z*z-x)/(2*z)
		if round(delta, 8)==0{
			break
		}
		z-=delta
	}
	fmt.Printf("Square root calculated in %v steps\n", i)
	return z
}

func round(x float64, precision int) float64{
	var n uint64=1
	var iRounded uint64
	var fRounded float64

	for i:=0; i<precision; i++{
		n*=10;
	}

	iRounded=uint64(x*float64(n))
	fRounded=float64(iRounded)/float64(n)

	return fRounded
}