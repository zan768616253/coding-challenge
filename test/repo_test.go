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

	Describe("CalculateRawData successful", func() {
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
		})

		Context(" has invalid parameter", func() {
			It(" should return err msg when userId does not exist", func() {
				err := cal.CalculateRawData()
				Expect(err).To(BeNil())
				rawData := cal.RawData
				Expect(rawData.Revenue).To(Equal(float64(32431)))
				Expect(rawData.Expenses).To(Equal(2220.68))
				Expect(rawData.Profit).To(Equal(float64(0)))
			})
		})
	})
})
