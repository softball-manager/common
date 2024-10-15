package dynamo

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/softball-manager/common/pkg/appconfig"
)

var (
	PlayerTableNamePrefix = "player-table"
	PlayerIDPrefix        = "Player#"

	TeamTableNamePrefix = "team-table"
	TeamIDPrefix        = "Team#"

	GameTableNamePrefix = "game-table"
	GameIDPrefix        = "Game#"
)

func CreateClient(cfg appconfig.AppConfig) *dynamodb.Client {
	logger := cfg.GetLogger()
	if cfg.GetEnv() == appconfig.LocalEnv {
		logger.Info("creating local dynamo db client")
		return createLocalClient(cfg.GetAWSConfig())
	}
	logger.Info("creating dynamo db client")
	return dynamodb.NewFromConfig(cfg.GetAWSConfig())
}

func createLocalClient(cfg aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://host.docker.internal:8000")
	})
}
