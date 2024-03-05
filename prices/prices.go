package prices

import (
	"fmt"
	"main/conversion"
	"main/fileManager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

// important are pointers to make sure we working on same class :D like singleton

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}

func (job *TaxIncludedPriceJob) LoadData() {

	// converting to float64 our string values
	lines, _ := fileManager.ReadLines("prices.txt")
	prices, _ := conversion.StringsToFloat(lines)

	job.InputPrices = prices
}
