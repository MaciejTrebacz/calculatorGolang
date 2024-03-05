package main

import (
	"main/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.2}

	for _, taxRate := range taxRates {
		var newJob = prices.NewTaxIncludedPriceJob(taxRate)
		newJob.Process()

	}
}
