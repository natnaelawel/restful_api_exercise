package entity

type Data struct {
	Success bool `json:"success"`
	Base string `json:"base",omitempty`
	Date string `json:"date",omitempty`
	TimeStamp int `json:"timestamp"`
	Rates Rates `json:"rates",omitempty`
}

type Rates struct {
	AED float64 `json:"AED",omitempty`
	CHF float64 `json:"CHF",omitempty`
	CNY float64 `json:"CNY",omitempty`
	GBP float64 `json:"GBP",omitempty`
	JPY float64 `json:"JPY",omitempty`
	USD float64 `json:"USD",omitempty`
	ETB float64 `json:"ETB",omitempty`
	EUR float64 `json:"EUR",omitempty`
}
