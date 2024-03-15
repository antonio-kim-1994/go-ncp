package main

import (
	"fmt"
	"github.com/antonio-kim-1994/go-ncp/routes/api/v1/usage"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestGetContractSummaryList(t *testing.T) {
	err := godotenv.Load("ncp.env")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	url := "/get-contract-summary-list"
	query := "?contractMonth=202402"

	r := SetupRouter()
	r.GET(url, usage.GetContractSummaryList)
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", url, query), nil)
	req.Header.Add("auth-key", "YWlqaW5ldAo=")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	//apiName := "Get Contract Summary List"
	//mockResponse := fmt.Sprintf(`{"message":"Success to call %s api."}`, apiName)
	//responseData, _ := io.ReadAll(w.Body)
	//assert.Equal(t, mockResponse, string(responseData))
}

func TestGetContractUsageList(t *testing.T) {
	err := godotenv.Load("ncp.env")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	url := "/get-contract-usage-list"
	query := "?startMonth=202402&endMonth"

	r := SetupRouter()
	r.GET(url, usage.GetContractSummaryList)
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", url, query), nil)
	req.Header.Add("auth-key", "YWlqaW5ldAo=")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProductDemandCostList(t *testing.T) {
	err := godotenv.Load("ncp.env")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	url := "/get-product-demand-cost-list"
	query := "?startMonth=202402&endMonth"

	r := SetupRouter()
	r.GET(url, usage.GetContractSummaryList)
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", url, query), nil)
	req.Header.Add("auth-key", "YWlqaW5ldAo=")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDemandCostList(t *testing.T) {
	err := godotenv.Load("ncp.env")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	url := "/get-demand-cost-list"
	query := "?startMonth=202402&endMonth"

	r := SetupRouter()
	r.GET(url, usage.GetContractSummaryList)
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", url, query), nil)
	req.Header.Add("auth-key", "YWlqaW5ldAo=")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRelationCodeList(t *testing.T) {
	err := godotenv.Load("ncp.env")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	url := "/get-relation-code-list"
	query := ""

	r := SetupRouter()
	r.GET(url, usage.GetContractSummaryList)
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", url, query), nil)
	req.Header.Add("auth-key", "YWlqaW5ldAo=")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetContractUsageListByDaily(t *testing.T) {
	err := godotenv.Load("ncp.env")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	now := time.Now()
	firstDayOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	url := "/get-contract-usage-daily"
	query := fmt.Sprintf("?useStartDay=%s&useEndDay=%s", now.Format("20060102"), firstDayOfMonth.Format("20060102"))

	r := SetupRouter()
	r.GET(url, usage.GetContractSummaryList)
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", url, query), nil)
	req.Header.Add("auth-key", "YWlqaW5ldAo=")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
