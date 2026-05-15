package main

import (
	"fmt"
	"math"
)

func main() {
	b := 2.5
	b3 := b * b * b

	f := func(x float64) float64 {
		sum := b3 + x*x*x
		return (1 + math.Pow(math.Sin(sum), 2)) / math.Cbrt(sum)
	}

	fmt.Println("Задача А")
	fmt.Println("   x   |     y")
	fmt.Println("-------|----------")

	xA := []float64{}
	for x := 1.28; x <= 3.28+0.001; x += 0.4 {
		xA = append(xA, x)
	}

	for i := 0; i < len(xA); i++ {
		yA := f(xA[i])
		fmt.Printf("%6.2f | %8.6f\n", xA[i], yA)
	}

	fmt.Println("\nЗадача Б")
	fmt.Println("   x    |     y")
	fmt.Println("--------|----------")

	xB := []float64{1.1, 2.4, 3.6, 1.7, 3.9}

	for i := 0; i < len(xB); i++ {
		yB := f(xB[i])
		fmt.Printf("%5.1f   | %8.6f\n", xB[i], yB)
	}
}
