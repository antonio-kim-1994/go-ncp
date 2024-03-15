package service

import (
	"encoding/json"
	"github.com/antonio-kim-1994/go-ncp/internal/model"
	"github.com/antonio-kim-1994/go-ncp/internal/repository"
	"github.com/rs/zerolog/log"
)

func MergeContractData(startDate, endDate string) ([]model.ContractData, error) {
	var (
		err        error
		mergedData []model.ContractData
	)

	todayUsage, err := repository.GetContractData(startDate)
	if err != nil {
		log.Err(err)
		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return nil, err
	}

	yesterdayUsage, err := repository.GetContractData(endDate)
	if err != nil {
		log.Err(err)
		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return nil, err
	}

	for _, i := range todayUsage {
		for _, j := range yesterdayUsage {
			if i.Code == j.Code {
				i.PastCount = j.CurrentCount
				i.ContractDifference = i.CurrentCount.Count - j.CurrentCount.Count
				mergedData = append(mergedData, i)
				break
			}
		}
	}

	//log.Info().Msg(fmt.Sprintf("\n\nMerged Data : %+v", mergedData))

	return mergedData, nil
}

func MergeProductDemandData(startDate, endDate string) ([]model.ProductDemandData, error) {
	var (
		err        error
		mergedData []model.ProductDemandData
	)

	todayUsage, err := repository.GetProductDemandUsageData(startDate)
	if err != nil {
		log.Err(err)
		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return nil, err
	}

	yesterdayUsage, err := repository.GetProductDemandUsageData(endDate)
	if err != nil {
		log.Err(err)
		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return nil, err
	}

	for _, i := range todayUsage {
		for _, j := range yesterdayUsage {
			if i.Code == j.Code {
				i.PastAmount = j.CurrentAmount
				i.AmountDifference = i.CurrentAmount.Amount - j.CurrentAmount.Amount
				mergedData = append(mergedData, i)
				break
			}
		}
	}

	//log.Info().Msg(fmt.Sprintf("\n\nMerged Data : %+v", mergedData))

	return mergedData, nil
}

func CompareContractData(startDate, endDate string) ([]byte, error) {
	comparedData, err := MergeContractData(startDate, endDate)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	jsonData, err := json.MarshalIndent(comparedData, "", "    ")
	if err != nil {
		log.Err(err)
		return nil, err
	}

	return jsonData, nil
}

func CompareProductDemandData(startDate, endDate string) ([]byte, error) {
	comparedData, err := MergeProductDemandData(startDate, endDate)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	jsonData, err := json.MarshalIndent(comparedData, "", "    ")
	if err != nil {
		log.Err(err)
		return nil, err
	}

	return jsonData, nil
}
