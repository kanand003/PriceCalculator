package main

import (
	cmdmanager "example.com/price-calculator/cmdManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.1, 0.2, 0.3}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%f.json", taxRate*100))
		cmd := cmdmanager.NewCommandManager()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		priceJob.Calculate()
	}

}
