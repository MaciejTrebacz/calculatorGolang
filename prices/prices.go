package prices

import (
	"fmt"
	"main/conversion"
	"main/fileManager"
)

type TaxIncludedPriceJob struct {
	IOManager         fileManager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(fm fileManager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   fm,
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

	job.TaxIncludedPrices = result

	job.IOManager.WriteJSON(job.TaxIncludedPrices)

}

func (job *TaxIncludedPriceJob) LoadData() {

	// converting to float64 our string values
	lines, _ := job.IOManager.ReadLines()
	prices, _ := conversion.StringsToFloat(lines)

	job.InputPrices = prices
}
