package model

type ContractSummaryAPI struct {
	GetContractSummaryListResponse ContractSummaryListResponse `json:"getContractSummaryListResponse"`
}

type ContractSummaryListResponse struct {
	TotalRaw            int               `json:"totalRows"`
	WriteDate           string            `json:"writeDate"`
	ContractSummaryList []ContractSummary `json:"contractSummaryList"`
	RequestID           string            `json:"requestId"`
	ReturnCode          string            `json:"returnCode"`
	ReturnMessage       string            `json:"returnMessage"`
}

type ContractSummary struct {
	MemberNo      string       `json:"memberNo"`
	RegionCode    string       `json:"regionCode"`
	ContractCount int          `json:"contractCount"`
	ContractType  ContractType `json:"contractType"`
}

type ContractType struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}

type ProductDemandAPI struct {
	GetProductDemandCostListResponse ProductDemandCostListResponse `json:"getProductDemandCostListResponse"`
}

type ProductDemandCostListResponse struct {
	WriteDate             string                  `json:"writeDate"`
	ProductDemandCostList []ProductDemandCostList `json:"productDemandCostList"`
}

type ProductDemandCostList struct {
	ProductDemandType ProductDemandType `json:"productDemandType"`
	UseAmount         int               `json:"useAmount"`
	DemandAmount      int               `json:"demandAmount"`
}

type ProductDemandType struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
