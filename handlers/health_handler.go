package handlers

import (
	"context"
	"net/http"
	"github.com/danrodsg/health-check-go/checker"
	"github.com/gin-gonic/gin"
	"time"
)


type HealthResponse struct {
	Status      string                      `json:"status"`
	Timestamp   time.Time                   `json:"timestamp"`
	Components map[string]string `json:"components"`
}

type HealthHandler struct {
	Checkers []checker.DependencyChecker
}

func NewHealthHandler(checkers []checker.DependencyChecker) *HealthHandler {
	return &HealthHandler{Checkers: checkers}
}


func (h *HealthHandler) HealthCheck(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	response := HealthResponse{
		Status: "UP",
		Timestamp: time.Now(),
		Components: make(map[string]string),
	}
	httpStatus := http.StatusOK

	// Executa todas as checagens em paralelo
	for _, chk := range h.Checkers {
		err := chk.Check(ctx)
		status := "UP"
		if err != nil {
			status = "DOWN"
			httpStatus = http.StatusServiceUnavailable 
		}
		response.Components[chk.Name()] = status
	}

	if httpStatus != http.StatusOK {
		response.Status = "DOWN"
	}

	c.JSON(httpStatus, response)
}