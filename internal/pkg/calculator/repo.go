package calculator

import "math"

const (
	Current                   = "current"
	Bank                      = "bank"
	CurrentAccountsReceivable = "current_accounts_receivable"
	CurrentAccountsPayable    = "current_accounts_payable"
)

var typesAssets []string = []string{Current, Bank, CurrentAccountsReceivable}
var typesLiabilities []string = []string{Current, CurrentAccountsPayable}

type Repository = *Calculator

// New Create Calculator
func New(report *ReportData) *Calculator {
	return &Calculator{Report: report}
}

// CalculateRawData Calculator raw data including revenue, expenses, profit, assets and liability.
func (c *Calculator) CalculateRawData() error {
	rawData := &RawData{}

	records := c.Report.Records
	for _, record := range records {
		value := record.TotalValue * 100
		valueType := record.ValueType
		category := record.AccountCategory
		acType := record.AccountType

		if category == "revenue" { // Total revenue
			rawData.Revenue += value
		} else if category == "expense" { // Total expenses
			rawData.Expenses += value
		} else if category == "assets" && valueType == "debit" && stringSliceContain(acType, typesAssets) { // Total assets
			rawData.Assets += value
		} else if category == "assets" && valueType == "credit" && stringSliceContain(acType, typesAssets) {
			rawData.Assets -= value
		} else if category == "liability" && valueType == "credit" && stringSliceContain(acType, typesLiabilities) { // Total liability
			rawData.Liability += value
		} else if category == "liability" && valueType == "dedit" && stringSliceContain(acType, typesLiabilities) {
			rawData.Liability -= value
		}

		// Gross profit
		if acType == "sales" && valueType == "debit" {
			rawData.Profit += value
		}
	}
	rawData.Revenue = rawData.Revenue / 100
	rawData.Profit = rawData.Profit / 100
	rawData.Expenses = rawData.Expenses / 100
	rawData.Assets = rawData.Assets / 100
	rawData.Liability = rawData.Liability / 100
	c.RawData = rawData
	return nil
}

// CalculateAdvancedData Calculator advanced data including profit and working capital ratio.
func (c *Calculator) CalculateAdvancedData() error {
	advancedData := &AdvancedData{}

	rawData := c.RawData
	if rawData.Revenue == 0 {
		advancedData.GrossProfitMargin, advancedData.NetProfitMargin = 0, 0
	} else {
		advancedData.GrossProfitMargin = math.Round(rawData.Profit / rawData.Revenue * 100)
		advancedData.NetProfitMargin = math.Round((rawData.Revenue - rawData.Expenses) / rawData.Revenue * 100)
	}

	if rawData.Liability == 0 {
		advancedData.WorkingCapitalRatio = 0
	} else {
		advancedData.WorkingCapitalRatio = math.Round(rawData.Assets / rawData.Liability * 100)
	}

	c.AdvancedData = advancedData
	return nil
}
