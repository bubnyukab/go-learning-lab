package day1

import (
	"fmt"
	"math"
)

func Day1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	
	// reverse loop 
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
	
	// reverse string
	s := "yourmama"
	reversed_s := ""
	for i := len(s)-1; i >= 0; i-- {
		reversed_s += string(s[i]) 	
	}
	fmt.Println(s, reversed_s)


	// for Sqrt func 
	num := 32.0
	fmt.Println(Sqrt(num))
	fmt.Println(math.Sqrt(num))
}

func Sqrt(x float64) float64 {
	z := 1.0

	//for i:= 0; i < 10; i++ {		
	//	z -= (z*z - x) / (2*z)
	//	
	//	fmt.Println(z)
	//}
	
	for {
		next := z - (z*z-x)/(2*z)
	
		if next-z < 0.0000000001 && next-z > -0.0000000001 {
		break
		}

		z = next
	}

	
	return z
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for i := range pic {
		pic[i] = make([]uint8, dy)
		for j := range pic[i] {
            pic[i][j] = uint8((i^j)*2)
        }
	}
	return pic 
}
