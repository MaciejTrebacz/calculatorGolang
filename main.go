package main

import (
	"fmt"
	"main/fileManager"
	"main/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.2}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan bool, len(taxRates))
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*1000))
		var newJob = prices.NewTaxIncludedPriceJob(fm, taxRate)
		go newJob.Process(doneChans[index], errorChans[index])

	}

	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
