package main

import (
	"fmt"
	"github.com/antonio-kim-1994/go-ncp/internal/repository"
	"github.com/antonio-kim-1994/go-ncp/routes"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

//	@title			Ncloud Cost monitoring API
//	@version		1.0
//	@description	Naver Cloud 비용 모니터링 API
//	@termsOfService	http://swagger.io/terms/

// @host		localhost:8080
// @BasePath	/ncp
func main() {
	// Load Env
	err := godotenv.Load("config/ncp.env")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load Environment.")
	}

	r := routes.InitRouter()
	err = repository.InitDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to Connect DB.")
		os.Exit(1)
	} else {
		log.Info().Msg("Success to Connect DB.")
	}

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Info().Msg(fmt.Sprintf("Server starting on %s", s.Addr))

	err = s.ListenAndServe()
	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("Server failed to start: %v", err))
		os.Exit(1)
	}
}
