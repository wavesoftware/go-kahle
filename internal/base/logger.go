package base

import (
	"github.com/pkg/errors"
	"github.com/wavesoftware/go-ensure"
	"go.uber.org/zap"
)

// Environment is a type that defines runtime environments
type Environment string

const (
	// ProductionEnvironment represents a production environment
	ProductionEnvironment  Environment = "production"

	// DevelopmentEnvironment represents a development environment
	DevelopmentEnvironment Environment = "development"
)

// CurrentEnvironment hold a current environment for the server
var CurrentEnvironment = calculateEnvironment()
var logger *zap.SugaredLogger

// Logger returns a global configured logger
func Logger() *zap.SugaredLogger {
	if logger == nil {
		logger = createLogger()
	}
	return logger
}

func createLogger() *zap.SugaredLogger {
	var log *zap.Logger
	var err error
	if CurrentEnvironment == ProductionEnvironment {
		log, err = zap.NewProduction()
	} else {
		log, err = zap.NewDevelopment()
	}
	ensure.NoError(err)
	return log.Sugar()
}

func calculateEnvironment() Environment {
	val := Env("ENVIRONMENT", string(ProductionEnvironment))
	if val == string(ProductionEnvironment) {
		return ProductionEnvironment
	}
	if val == string(DevelopmentEnvironment) {
		return DevelopmentEnvironment
	}
	panic(errors.Errorf("invalid environment: %s", val))
}
