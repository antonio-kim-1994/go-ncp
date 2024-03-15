package model

type ContractData struct {
	Code               string        `json:"code"`
	CodeName           string        `json:"codeName"`
	ContractDifference int           `json:"contractDifference"`
	CurrentCount       ContractCount `json:"currentCount"`
	PastCount          ContractCount `json:"pastCount"`
}

type ContractCount struct {
	Count     int    `json:"count"`
	WriteDate string `json:"writeDate"`
}

type ProductDemandData struct {
	Code             string       `json:"code"`
	CodeName         string       `json:"codeName"`
	AmountDifference int          `json:"amountDifference"`
	CurrentAmount    DemandAmount `json:"currentAmount"`
	PastAmount       DemandAmount `json:"pastAmount"`
}

type DemandAmount struct {
	Amount    int    `json:"amount"`
	WriteDate string `json:"writeDate"`
}
