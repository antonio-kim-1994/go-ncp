package repository

import (
	"fmt"
	"github.com/antonio-kim-1994/go-ncp/internal/model"
	"github.com/rs/zerolog/log"
	"time"
)

func WriteContractSummaryList(d *model.ContractSummaryAPI) error {
	var err error
	db := GetDB()

	tx, err := db.Begin()
	if err != nil {
		log.Err(err).Msg("Failed to create transaction.")
		return err
	}

	defer tx.Rollback()

	// 오늘 날짜 데이터가 존재 할 경우 DB 저장 취소
	dataCheckQuery := `
SELECT EXISTS(SELECT 1 FROM devops.ncloud_contract_summary_list WHERE write_date = ?) AS data_exists;
`
	var exists bool
	err = tx.QueryRow(dataCheckQuery, time.Now().Format("2006-01-02")).Scan(&exists)
	if err != nil {
		tx.Rollback() // 쿼리 실패 시 롤백
		return err
	}

	if exists == true {
		log.Info().Msg("Contract Summary Data Already exists. Cancel the data storage request.\n")
		err = fmt.Errorf("Contract Summary Data Already exists.")
		tx.Rollback() // 데이터 존재 시 롤백
		return err
	} else {
		for _, data := range d.GetContractSummaryListResponse.ContractSummaryList {
			_, err = tx.Exec(
				"INSERT INTO devops.ncloud_contract_summary_list ( contract_type_code, contract_code_name, contract_region, contract_count, write_date ) VALUES (?, ?, ?, ?, ?)",
				data.ContractType.Code,
				data.ContractType.CodeName,
				data.RegionCode,
				data.ContractCount,
				time.Now().Format("2006-01-02"),
			)
			if err != nil {
				log.Err(err).Msg("Failed to save data.")
				tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func WriteProductDemandCostList(d *model.ProductDemandAPI) error {
	var err error
	db := GetDB()

	tx, err := db.Begin()
	if err != nil {
		log.Err(err).Msg("Failed to create transaction.")
		return err
	}

	defer tx.Rollback()

	// 오늘 날짜 데이터가 존재 할 경우 DB 저장 취소
	dataCheckQuery := `
SELECT EXISTS(SELECT 1 FROM devops.ncloud_product_cost_list WHERE write_date = ?) AS data_exists;
`
	var exists bool
	err = tx.QueryRow(dataCheckQuery, time.Now().Format("2006-01-02")).Scan(&exists)
	if err != nil {
		tx.Rollback() // 쿼리 실패 시 롤백
		return err
	}

	if exists == true {
		log.Info().Msg("Product Demand Cost Data Already exists. Cancel the data storage request.\n")
		err = fmt.Errorf("Product Demand Cost List Data Already exists.")
		tx.Rollback() // 데이터 존재 시 롤백
		return err
	} else {
		for _, data := range d.GetProductDemandCostListResponse.ProductDemandCostList {
			_, err = tx.Exec(
				"INSERT INTO devops.ncloud_product_cost_list ( product_demand_type_code, product_demand_code_name, product_use_amount, product_demand_amount, write_date ) VALUES (?, ?, ?, ?, ?)",
				data.ProductDemandType.Code,
				data.ProductDemandType.CodeName,
				data.UseAmount,
				data.DemandAmount,
				time.Now().Format("2006-01-02"),
			)
			if err != nil {
				log.Err(err).Msg("Failed to save data.")
				tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
