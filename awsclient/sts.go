package awsclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type StsClient struct {
	client *sts.Client
}

func NewStsClient(cfg aws.Config) *StsClient {
	client := sts.NewFromConfig(cfg)
	return &StsClient{
		client: client,
	}
}

func (sc *StsClient) SetCredentialProvider(ctx context.Context, cfg *aws.Config, account, role string) {
	an := arn.ARN{
		Partition: "aws",
		Service:   "iam",
		Region:    "",
		AccountID: account,
		Resource:  role,
	}
	provider := stscreds.NewAssumeRoleProvider(sc.client, an.String())

	cfg.Credentials = aws.NewCredentialsCache(provider)
}
