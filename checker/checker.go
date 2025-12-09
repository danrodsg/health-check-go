package checker

import (
	"context"
	"errors"
	"time"
)


type DependencyChecker interface {
	Check(ctx context.Context) error
	Name() string
}


type DatabaseChecker struct{}

func NewDatabaseChecker() *DatabaseChecker {
	return &DatabaseChecker{}
}

func (d *DatabaseChecker) Name() string {
	return "database"
}

func (d *DatabaseChecker) Check(ctx context.Context) error {
	select {
	case <-time.After(50 * time.Millisecond): 
		return nil
	case <-ctx.Done():
		return ctx.Err() 
	}
}

type ExternalServiceChecker struct{}

func NewExternalServiceChecker() *ExternalServiceChecker {
	return &ExternalServiceChecker{}
}

func (e *ExternalServiceChecker) Name() string {
	return "external_api"
}

func (e *ExternalServiceChecker) Check(ctx context.Context) error {

	if time.Now().Second()%10 == 0 { 
		return errors.New("external service is temporarily unavailable")
	}
	return nil
}