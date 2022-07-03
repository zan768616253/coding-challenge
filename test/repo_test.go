package test

import (
	"9Spokes-code-challenge/internal/pkg/calculator"
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"testing"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator")
}

var _ = Describe("Calculator Test", func() {
	var cal *calculator.Calculator

	Describe("CalculateRawData", func() {
		// test CalculateRawData
		BeforeEach(func() {
			jsonFile, err := os.Open("data_test.json")
			defer jsonFile.Close()
			Expect(err).To(BeNil())

			text, err := ioutil.ReadAll(jsonFile)
			Expect(err).To(BeNil())

			var report *calculator.ReportData
			err = json.Unmarshal([]byte(text), &report)
			Expect(err).To(BeNil())

			cal = calculator.New(report)
			err = cal.CalculateRawData()
			Expect(err).To(BeNil())
			err = cal.CalculateAdvancedData()
			Expect(err).To(BeNil())
		})

		Context(" calculate successfully 1", func() {
			It(" should get all the raw data", func() {
				rawData := cal.RawData
				Expect(rawData.Revenue).To(Equal(float64(32431)))
				Expect(rawData.Expenses).To(Equal(2220.68))
				Expect(rawData.Profit).To(Equal(float64(0)))
			})
		})

		Context(" calculate successfully 2", func() {
			It(" should get all the advanced data", func() {
				advancedData := cal.AdvancedData
				Expect(advancedData.GrossProfitMargin).To(Equal(float64(0)))
				Expect(advancedData.NetProfitMargin).To(Equal(float64(93)))
				Expect(advancedData.GrossProfitMargin).To(Equal(float64(0)))
			})
		})
	})
})
