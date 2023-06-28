package main

import "fmt"

func main() {
	cenyProdaneho := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0

	for _, val := range cenyProdaneho {
		sum += val
	}
	sum /= float64(len(cenyProdaneho))

	fmt.Printf("Prumerna cena byla %v\n", sum)
}
