package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

func CheckAuthKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("auth-key")
		signingKey := os.Getenv("VALIDATION_KEY")

		if h != signingKey {
			log.Warn().Msg(fmt.Sprintf("[Requested Header : %s] invalid header request.", h))
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "invalid header",
				"message": "유효하지 않은 접근입니다.",
			})
			c.Abort()
			return
		} else {
			log.Info().Msg(fmt.Sprintf("[Requested Header: %s] valid header access.", h))
			c.Next()
		}
	}
}

func CheckDateQueryValidation(c *gin.Context) (string, string, error) {
	var (
		//date              = time.Now()
		startDate, sExist = c.GetQuery("startDate")
		endDate, eExist   = c.GetQuery("endDate")
		err               error
	)

	// startDate, endDate Query가 존재하지 않을 경우 오늘 날짜를 기준으로 startDate와 endDate 지정
	if !sExist || !eExist {
		err = fmt.Errorf("'startDate' or 'endDate' Query parameter is required\n")
		return "", "", err
	} else {
		// startDate, endDate 포맷 검증('YYYY-MM-DD'가 아닌 경우 응답 거부)
		_, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			err = fmt.Errorf("startDate Value '%s' is Invalid Query Format. Query Must be submitted in 'YYYY-MM-DD' format", startDate)
			return "", "", err
		}

		_, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			err = fmt.Errorf("endDate Value '%s' is Invalid Query Format. Query Must be submitted in 'YYYY-MM-DD' format", endDate)
			return "", "", err
		}
	}

	return startDate, endDate, nil
}
