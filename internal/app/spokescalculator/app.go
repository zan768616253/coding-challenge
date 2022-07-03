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
	var report calculator.ReportData
	err = json.Unmarshal([]byte(data), &report)
	if err != nil {
		return
	}

}
