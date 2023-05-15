package awsclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ecsClient interface {
	CreateCluster(ctx context.Context, params *ecs.CreateClusterInput, optFns ...func(*ecs.Options)) (*ecs.CreateClusterOutput, error)
}

type Ecs struct {
	client ecsClient
}

func New(cfg aws.Config) *Ecs {
	client := ecs.NewFromConfig(cfg)
	return &Ecs{
		client: client,
	}
}

func (ec *Ecs) CreteCluster(ctx context.Context, cluster, service string) string {
	ec.client.CreateCluster(ctx, &ecs.CreateClusterInput{
		ClusterName: aws.String(cluster),
	})
	return service + "-green"
}
