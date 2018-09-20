package main

import "fmt"
import "math"
import "bytes"

/* A tour of Go skonczylem na Zero values 12/17 */

func main(){
	var a,b int = 10,90
	var MaxInt uint64
	var p int = 1
	var c uint64 = 10

	MaxInt=1<<64-1

	x:=110 	// uzywanie ':=' pozwala pominac slowo kluczowe 'var'
			// dziala tylko wewnatrz funkcji

	fmt.Printf("Pierwiastek kwadratowy z 7: %g\n",  math.Sqrt(7))
	fmt.Printf("Liczba Pi: %g\n", math.Pi)
	fmt.Printf("%d+%d=%d\n", a, b, suma(a,b))
	fmt.Println(calkowita_i_reszta(10.5))
	fmt.Printf("x=%d\n", x)
	fmt.Printf("Type %T, value %d\n", MaxInt, MaxInt)

	fmt.Printf("p=%d\n", p)
	fmt.Printf("p<<1= %d\n", p<<1)
	fmt.Printf("p<<2= %d\n", p<<2)
	fmt.Printf("p<<3= %d\n", p<<3)
	fmt.Printf("p<<10= %d\n", p<<10)
	fmt.Printf("binarna reprezenracja liczby %d= %s\n", c, int2bin(c))
	fmt.Printf("binarna reprezenracja liczby %d= %s\n", MaxInt, int2bin(MaxInt))
	fmt.Printf("bledne modulo 2 z %g=  %d\n", float64(MaxInt), uint64(math.Mod(float64(MaxInt),2)))
	fmt.Printf("moje modulo 2 z %d=  %d\n", MaxInt, modulo2(MaxInt))
}

func int2bin(x uint64) string{
	var r uint64
	var a uint64=x
	var s []uint64

	buffer:=bytes.NewBufferString("") 

	for a>0{
		r=modulo2(a)
		s=append(s, uint64(r))
		a=(a-r)/2
	}
	i:=len(s)
	for i>0{
		fmt.Fprintf(buffer, "%d ", s[i-1])
		i--
	}
	return buffer.String()
}

func modulo2(x uint64) uint64{
	var y uint64
	y=x>>1
	return x-(y<<1)
}

func calkowita_i_reszta(f float64) (fc float64, fr float64){

	fc=math.Trunc(f)
	fr=f-fc
	return
}

func suma(a int, b int) int{
	return a+b;
}
