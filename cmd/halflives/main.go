package main

import (
	"fmt"
	"math"
)

//"gitlab.com/waveletlet/pilltime"

func main() {
	// half-life
	//y = a * 2^-(t/h)
	a := 450.0
	h := float64(84)
	trigger := 0
	for t := float64(0); t < h*120; t += 5 {
		y := a * math.Exp2(-t/h)
		switch {
		case y < 0.00001 && trigger < 4:
			fmt.Printf("< 0.00001 at %v:\t\t%v half lives\t\t%v left\n", t, t/h, y)
			trigger += 1
			break
		case y < 0.0001 && trigger < 3:
			fmt.Printf("< 0.0001 at %v:\t\t%v half lives\t\t%v left\n", t, t/h, y)
			trigger += 1
		case y < 0.001 && trigger < 2:
			fmt.Printf("< 0.001 at %v:\t\t%v half lives\t\t%v left\n", t, t/h, y)
			trigger += 1
		case y < 0.01 && trigger < 1:
			fmt.Printf("< 0.01 at %v:\t\t%v half lives\t\t%v left\n", t, t/h, y)
			trigger += 1
		}

		// exponential decay
		//y = a × e^kt
		//		k := math.Log(a/2/a) / h
		//		y2 := a * math.Exp(k*t)
		//
		//		fmt.Printf("Expoential:\t%v\t%v\n", y2, k)
	}
}
