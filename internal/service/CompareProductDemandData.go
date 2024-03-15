package service

//func mergeProductDemandData(today, yesterday string) ([]model.ProductDemandData, error) {
//	var (
//		err error
//		mergedData []model.ProductDemandData
//	)
//
//	todayUsage, err := repository.GetProductDemandUsageData(today)
//	if err != nil {
//		log.Err(err)
//		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
//		return nil, err
//	}
//
//	yesterdayUsage, err := repository.GetProductDemandUsageData(yesterday)
//	if err != nil {
//		log.Err(err)
//		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
//		return nil, err
//	}
//
//	for _, i := range todayUsage {
//		for _, j := range yesterdayUsage {
//			if i.Code == j.Code {
//				i.PastAmount = j.CurrentAmount
//				i.AmountDifference = i.CurrentAmount.Amount - j.CurrentAmount.Amount
//				mergedData = append(mergedData, i)
//				break
//			}
//		}
//	}
//
//	//log.Info().Msg(fmt.Sprintf("\n\nMerged Data : %+v", mergedData))
//
//	return mergedData, nil
//}

//func CompareProductDemandData(startDate, endDate string) ([]byte, error) {
//	comparedData, err := mergeProductDemandData(startDate, endDate)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
//		return
//	}
//
//	jsonData, err := json.MarshalIndent(comparedData, "", "    ")
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
//		return
//	}
//
//	c.Data(http.StatusOK, "application/json", jsonData)
//	return
//}
