package main

import (
	"example.com/price-calculator/prices"
)


func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.7, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

}


