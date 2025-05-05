package main

import "fmt"

func main() {
	prices := []float64{3, 4, 5}
	taxRates := []float64{0.1, 0.2, 0.3}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println(result)
}
