package main

import (
	"github.com/danrodsg/health-check-go/checker"
	"github.com/danrodsg/health-check-go/handlers"
	"github.com/danrodsg/health-check-go/metrics"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func main() {
	// 1. Inicializar os Checkers (Dependências)
	dbChecker := checker.NewDatabaseChecker()
	externalChecker := checker.NewExternalServiceChecker()

	// 2. Agrupar os Checkers
	checkers := []checker.DependencyChecker{dbChecker, externalChecker}
	
	// 3. Inicializar o Handler
	healthHandler := handlers.NewHealthHandler(checkers)

	// 4. Configurar e rodar o Servidor
	router := gin.Default()

	// Rota para o Health Check
	router.GET("/health", func(c *gin.Context) {
		start := time.Now()
		// O handler faz o trabalho
		healthHandler.HealthCheck(c)
		// Registra a latência APÓS a execução do handler
		metrics.RecordLatency(start) 
	})

	// Rota para Métricas Prometheus (Exposta pelo client_golang)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Run(":8080")
}