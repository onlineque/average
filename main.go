package average

import "fmt"

func GetAverage(prices ...float64) float64 {
	var sum float64 = 0.0
	for _, val := range prices {
		sum += val
	}

	sum /= float64(len(prices))
	return sum
}
