package main

import (
	"fmt"
	"main/fileManager"
	"main/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.2}

	for _, taxRate := range taxRates {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*1000))
		var newJob = prices.NewTaxIncludedPriceJob(fm, taxRate)
		newJob.Process()

	}
}
