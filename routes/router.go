package routes

import (
	"fmt"
	"github.com/antonio-kim-1994/go-ncp/internal/service"
	"github.com/antonio-kim-1994/go-ncp/pkg/middleware"
	"github.com/antonio-kim-1994/go-ncp/routes/api/v1/info"
	"github.com/antonio-kim-1994/go-ncp/routes/api/v1/usage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Router
	ncp := r.Group("/api/v1")
	ncp.Use(middleware.CheckAuthKey())

	// usage route
	u := ncp.Group("usage")
	{
		// GET
		u.GET("contract-summary-list", usage.GetContractSummaryList)
		u.GET("product-demand-cost-list", usage.GetProductDemandCostList)
		u.GET("contract-usage-list", usage.GetContractUsageList)
		u.GET("demand-cost-list", usage.GetDemandCostList)
		u.GET("contract-usage-daily", usage.GetContractUsageListByDaily)

		// Compare
		u.POST("compare", compareData)

		// POST
		u.POST("contract-summary-list", usage.SaveContractSummaryList)
		u.POST("product-demand-cost-list", usage.SaveProductDemandCostList)
	}

	p := ncp.Group("info")
	{
		p.GET("product-price-list", info.GetProductPriceList)
		p.GET("relation-code-list", info.GetRelationCodeList)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func compareData(c *gin.Context) {
	var err error
	startDate, endDate, err := middleware.CheckDateQueryValidation(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, x := c.GetQuery("type")
	if !x {
		err = fmt.Errorf("'type' query parameter is required")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	switch t {
	case "contract":
		result, err := service.CompareContractData(startDate, endDate)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Data(http.StatusOK, "application/json", result)
		return
	case "product":
		result, err := service.CompareProductDemandData(startDate, endDate)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Data(http.StatusOK, "application/json", result)
		return
	default:
		err = fmt.Errorf("invalid 'type' query parameter submitted. Please check your parameter value")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
