package info

import (
	"fmt"
	"github.com/antonio-kim-1994/go-ncp/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
)

func GetRelationCodeList(c *gin.Context) {
	url := os.Getenv("COST_URL")
	accessKey := os.Getenv("NCP_ACC")
	request := util.RequestInfo{
		Method: "GET",
		Path:   "/billing/v1/cost/getCostRelationCodeList",
		Query:  fmt.Sprintf("?responseFormatType=json"),
	}

	h := util.GenerateHmac(request)

	req, err := http.NewRequest(request.Method, fmt.Sprintf("%s%s%s", url, request.Path, request.Query), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		log.Err(err)
		return
	}

	req.Header.Add("x-ncp-apigw-timestamp", h.Timestamp)
	req.Header.Add("x-ncp-iam-access-key", accessKey)
	req.Header.Add("x-ncp-apigw-signature-v2", h.Signature)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response"})
		return
	}

	c.Data(http.StatusOK, "application/json", body)

	//var response GetContractSummaryListResponse
	//err = xml.Unmarshal(body, &response)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response"})
	//	fmt.Println("Error unmarshalling XML:", err)
	//	return
	//}
	//
	//c.JSON(http.StatusOK, response)
	//fmt.Println(response)
	return
}
