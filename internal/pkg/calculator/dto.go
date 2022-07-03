package calculator

type Report struct {
	Category       string   `json:"object_category"`
	ConnectionId   string   `json:"connection_id"`
	User           string   `json:"user"`
	CreationDate   string   `json:"object_creation_date"`
	Records        []Record `json:"data"`
	Currency       string   `json:"currency"`
	OriginType     string   `json:"object_origin_type"`
	OriginCategory string   `json:"object_origin_category"`
	Type           string   `json:"object_type"`
	Class          string   `json:"object_class"`
	BalanceDate    string   `json:"balance_date"`
}

type Record struct {
	AccountCode       string  `json:"account_code"`
	AccountCurrency   string  `json:"account_currency"`
	AccountIdentifier string  `json:"account_identifier"`
	AccountStatus     string  `json:"account_status"`
	AccountType       string  `json:"account_type"`
	AccountTypeBank   string  `json:"account_type_bank"`
	ValueType         string  `json:"value_type"`
	AccountCategory   string  `json:"account_category"`
	AccountName       string  `json:"account_name"`
	SystemAccount     string  `json:"system_account"`
	TotalValue        float64 `json:"total_value"`
}
