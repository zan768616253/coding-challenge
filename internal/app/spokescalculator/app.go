package spokescalculator

import (
	"9Spokes-code-challenge/internal/pkg/calculator"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func Main() {
	// Open file
	log.Println("9Spokes Calculator Start Processing.")
	jsonFile, err := os.Open("data.json")
	if err != nil {
		return
	}
	// Centralize the error handling here
	defer func() {
		if err != nil {
			log.Println("9Spokes Calculator err:", err)
		} else {
			log.Println("9Spokes Calculator Finish Processing.")
		}
		// Close the jsonFile reader
		err = jsonFile.Close()
		if err != nil {
			return
		}
	}()

	// Read file
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	// Convert binary data to calculator.Report
	var report *calculator.ReportData
	err = json.Unmarshal([]byte(data), &report)
	if err != nil {
		return
	}
	cal := calculator.New(report)
	err = cal.CalculateRawData()
	if err != nil {
		return
	}
	err = cal.CalculateAdvancedData()
	if err != nil {
		return
	}
	log.Printf("Revenue: $%.2f\n", cal.RawData.Revenue)
	log.Printf("Expenses: $%.2f\n", cal.RawData.Expenses)
	log.Printf("Profit: $%.2f\n", cal.RawData.Profit)
	log.Printf("Gross Profit Margin: %.0f%%\n", cal.AdvancedData.GrossProfitMargin)
	log.Printf("Net Profit Margin: %.0f%%\n", cal.AdvancedData.NetProfitMargin)
	log.Printf("Working Capital Ratio: %.0f%%\n", cal.AdvancedData.WorkingCapitalRatio)
}
