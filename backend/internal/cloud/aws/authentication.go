package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewClient() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO())
}
