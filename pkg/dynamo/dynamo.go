package dynamo

import (
	"softball-manager/common/pkg/awsconfig"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func CreateLocalClient() *dynamodb.Client {
	cfg, err := awsconfig.GetLocalAWSConfig()
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://host.docker.internal:8000")
	})
}
