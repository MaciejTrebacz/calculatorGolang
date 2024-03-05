package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloat(lines []string) ([]float64, error) {
	var floats []float64
	for _, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price to float failed")
			fmt.Println(err)
			return nil, err
		}
		floats = append(floats, floatPrice)

	}
	return floats, nil
}
