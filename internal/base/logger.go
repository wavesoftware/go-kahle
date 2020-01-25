package base

import (
	"github.com/pkg/errors"
	"github.com/wavesoftware/go-ensure"
	"go.uber.org/zap"
)

type Environment string

const (
	ProductionEnvironment  Environment = "production"
	DevelopmentEnvironment Environment = "development"
)

var CurrentEnvironment = calculateEnvironment()
var logger *zap.SugaredLogger

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
