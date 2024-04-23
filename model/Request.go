package model 

// Allowance
type Allowance struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

// RequestModel
type RequestModel struct {
	TotalIncome float64    `json:"totalIncome"`
	WHT         float64    `json:"wht"`
	Allowances  []Allowance `json:"allowances"`
}