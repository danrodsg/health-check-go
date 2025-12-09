package checker

import (
	"context"
	"errors"
	"time"
)

// DependencyChecker define a interface para qualquer componente que queremos checar.
type DependencyChecker interface {
	Check(ctx context.Context) error
	Name() string
}

// DatabaseChecker simula a checagem de um banco de dados
type DatabaseChecker struct{}

func NewDatabaseChecker() *DatabaseChecker {
	return &DatabaseChecker{}
}

func (d *DatabaseChecker) Name() string {
	return "database"
}

// Check simula a conexão com o DB.
func (d *DatabaseChecker) Check(ctx context.Context) error {
	// Simula a tentativa de conexão que pode falhar ou ter timeout
	select {
	case <-time.After(50 * time.Millisecond): // Tempo de conexão OK
		return nil
	case <-ctx.Done():
		return ctx.Err() // Timeout do Context atingido
	}
}

// ExternalServiceChecker simula a checagem de um serviço externo
type ExternalServiceChecker struct{}

func NewExternalServiceChecker() *ExternalServiceChecker {
	return &ExternalServiceChecker{}
}

func (e *ExternalServiceChecker) Name() string {
	return "external_api"
}

// Check simula a chamada a um serviço externo.
func (e *ExternalServiceChecker) Check(ctx context.Context) error {
	// Simula uma falha ocasional ou um timeout
	if time.Now().Second()%10 == 0 { // Falha a cada 10 segundos para demonstração
		return errors.New("external service is temporarily unavailable")
	}
	return nil
}