package usage

import (
	"fmt"
	"github.com/antonio-kim-1994/go-ncp/internal/model"
	"github.com/antonio-kim-1994/go-ncp/internal/repository"
	"github.com/antonio-kim-1994/go-ncp/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
	"time"
)

func FetchProductDemandCostList(c *gin.Context) (*model.ProductDemandAPI, error) {
	var (
		url        = os.Getenv("COST_URL")
		accessKey  = os.Getenv("NCP_ACC")
		startMonth = c.Query("startMonth")
		endMonth   = c.Query("endMonth")

		request = util.RequestInfo{
			Method: "GET",
			Path:   "/billing/v1/cost/getProductDemandCostList",
			Query:  fmt.Sprintf("?startMonth=%s&endMonth=%s&responseFormatType=json", startMonth, endMonth),
		}

		h = util.GenerateHmac(request)

		result *model.ProductDemandAPI
	)

	req, err := http.NewRequest(request.Method, fmt.Sprintf("%s%s%s", url, request.Path, request.Query), nil)
	if err != nil {
		log.Err(err).Msg("Failed Generate request.")
		return nil, err
	}

	req.Header.Add("x-ncp-apigw-timestamp", h.Timestamp)
	req.Header.Add("x-ncp-iam-access-key", accessKey)
	req.Header.Add("x-ncp-apigw-signature-v2", h.Signature)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("Error sending request")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("Error reading response")
		return nil, err
	}

	//c.Data(http.StatusOK, "application/json", body)

	err = util.ParseJSON(body, &result)
	if err != nil {
		log.Err(err).Msg("Fail to parse response data.")
		return nil, err
	}

	return result, nil
}

func GetProductDemandCostList(c *gin.Context) {
	productDemandCostList, err := FetchProductDemandCostList(c)
	if err != nil {
		log.Err(err).Msg("Failed to fetch [ContractSummaryList] API.")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch [ContractSummaryList] API."})
		return
	}

	productDemandCostList.GetProductDemandCostListResponse.WriteDate = time.Now().Format("2006-01-02")

	parsedJson, err := util.MarshalSturct(productDemandCostList)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.Data(http.StatusOK, "application/json", parsedJson)
}

func SaveProductDemandCostList(c *gin.Context) {
	productDemandCostList, err := FetchProductDemandCostList(c)
	if err != nil {
		log.Err(err).Msg("Failed to fetch [GetProductDemandCostList] API.")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch [GetProductDemandCostList] API."})
		return
	}

	err = repository.WriteProductDemandCostList(productDemandCostList)
	if err != nil {
		log.Err(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully saved Product Demand Cost List Data."})
		return
	}
}
