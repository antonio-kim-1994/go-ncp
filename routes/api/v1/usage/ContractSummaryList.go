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

// FetchContractSummaryListAPI
//
//	@Summary	Get Ncloud contract summary list
//	@Tags		contractSummary
//	@ID			get-contract-summary-list
//	@Accept		json
//	@Produce	json
//	@Param		auth-key		header		string	true	"Auth Key"
//	@Param		contractMonth	query		string	true	"Contract Month"
//	@Success	200				{string}	string	"ok"
//	@Failure	400				{string}	string	"Failed to call API"
//	@Router		/get-contract-summary-list [get]
func FetchContractSummaryListAPI(c *gin.Context) (*model.ContractSummaryAPI, error) {
	var (
		url           = os.Getenv("COST_URL")
		accessKey     = os.Getenv("NCP_ACC")
		contractMonth = c.Query("contractMonth")
		request       = util.RequestInfo{
			Method: "GET",
			Path:   "/billing/v1/cost/getContractSummaryList",
			Query:  fmt.Sprintf("?contractMonth=%s&responseFormatType=json", contractMonth),
		}
		h = util.GenerateHmac(request)

		result *model.ContractSummaryAPI
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

	err = util.ParseJSON(body, &result)
	if err != nil {
		log.Err(err).Msg("Fail to parse response data.")
		return nil, err
	}

	return result, nil
}

func GetContractSummaryList(c *gin.Context) {
	contractData, err := FetchContractSummaryListAPI(c)
	if err != nil {
		log.Err(err).Msg("Failed to fetch [ContractSummaryList] API.")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch [ContractSummaryList] API."})
		return
	}

	contractData.GetContractSummaryListResponse.WriteDate = time.Now().Format("2006-01-02")

	parsedJson, err := util.MarshalSturct(contractData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.Data(http.StatusOK, "application/json", parsedJson)
}

func SaveContractSummaryList(c *gin.Context) {
	contractData, err := FetchContractSummaryListAPI(c)
	if err != nil {
		log.Err(err).Msg("Failed to fetch [GetContractSummaryList] API.")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch [GetContractSummaryList] API."})
		return
	}

	err = repository.WriteContractSummaryList(contractData)
	if err != nil {
		log.Err(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully saved Contract Summary List Data."})
	return
}
