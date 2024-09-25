package appconfig

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"go.uber.org/zap"
)

var (
	LocalEnv = "local"
)

type AppConfig interface {
	GetEnv() string
	GetAWSConfig() aws.Config
	GetLogger() *zap.Logger
}

func GetEnvironment() string {
	env, present := os.LookupEnv("ENV")
	if present {
		return env
	}
	return LocalEnv
}

func GetEnvVarStringOrDefault(envVar string, defaultValue string) string {
	env, present := os.LookupEnv(envVar)
	if present {
		return env
	}
	return defaultValue
}
