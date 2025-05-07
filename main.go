package main

import (
	cmdmanager "example.com/price-calculator/cmdManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.1, 0.2, 0.3}
	doneChans := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%f.json", taxRate*100))
		cmd := cmdmanager.NewCommandManager()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		go priceJob.Calculate(doneChans[i])
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
