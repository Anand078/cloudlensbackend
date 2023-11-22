package models

// Poc details
type Poc struct {
	ID         int64  `json:"id,omitempty"`
	Account    string `json:"account,omitempty"`
	Pocname    string `json:"pocname,omitempty"`
	Technology string `json:"technology,omitempty"`
	Objective  string `json:"objective,omitempty"`
	Owner      string `json:"owner,omitempty"`
	Teamname   string `json:"teamname,omitempty"`
	Status     string `json:"status,omitempty"`
	Remarks    string `json:"remarks,omitempty"`
	Link       string `json:"link,omitempty"`
	AssignedTo string `json:"assignedto,omitempty"`
}

// Feed details
type Feed struct {
	ID   int64  `json:"id,omitempty"`
	Feed string `json:"feed,omitempty"`
}

// Poc tech count detail
type TechCount struct {
	Technology string `json:"technology"`
	Count      string `json:"count"`
}

// Poc tech count detail
type PieChartCount struct {
	Account string `json:"account"`
	Count   string `json:"count"`
}
