package repository

import (
	"github.com/antonio-kim-1994/go-ncp/internal/model"
	"github.com/rs/zerolog/log"
)

func GetContractData(date string) ([]model.ContractData, error) {
	var (
		err error
		db  = GetDB()

		query = `
SELECT 
    ncloud_contract_summary_list.contract_type_code, 
    ncloud_contract_summary_list.contract_code_name, 
    ncloud_contract_summary_list.contract_count,
    ncloud_contract_summary_list.write_date
from devops.ncloud_contract_summary_list
where write_date = ?;
`
		contractData []model.ContractData
	)

	rows, err := db.Query(query, date)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d model.ContractData
		if err := rows.Scan(&d.Code, &d.CodeName, &d.CurrentCount.Count, &d.CurrentCount.WriteDate); err != nil {
			log.Err(err)
			return nil, err
		}
		contractData = append(contractData, d)
	}

	if err = rows.Err(); err != nil {
		log.Err(err)
		return nil, err
	}

	return contractData, nil
}

func GetProductDemandUsageData(date string) ([]model.ProductDemandData, error) {
	var (
		err error
		db  = GetDB()

		query = `
SELECT 
    ncloud_product_cost_list.product_demand_type_code, 
    ncloud_product_cost_list.product_demand_code_name, 
    ncloud_product_cost_list.product_demand_amount, 
    ncloud_product_cost_list.write_date
from devops.ncloud_product_cost_list
where write_date = ?;
`
		productDemandData []model.ProductDemandData
	)

	rows, err := db.Query(query, date)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var d model.ProductDemandData
		if err := rows.Scan(&d.Code, &d.CodeName, &d.CurrentAmount.Amount, &d.CurrentAmount.WriteDate); err != nil {
			log.Err(err)
			return nil, err
		}
		productDemandData = append(productDemandData, d)
	}

	if err = rows.Err(); err != nil {
		log.Err(err)
		return nil, err
	}

	return productDemandData, nil
}
