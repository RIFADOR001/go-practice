package main

import (
	"fmt"

	"example.com/price-calculator/prices"
	"example.com/price-calculator/filemanager"
	// "example.com/price-calculator/cmdmanager"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.10, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		// cmdm := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)

		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

}


