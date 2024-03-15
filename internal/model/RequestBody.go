package model

type CompareRequestData struct {
	StartDate string `json:"startDate"` // yyyy-mm-dd
	EndDate   string `json:"endDate"`   // yyyy-mm-dd
	Type      string `json:"type"`      // contract, product
}
