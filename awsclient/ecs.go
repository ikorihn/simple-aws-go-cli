package awsclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type EcsManager interface {
	CreateCluster(ctx context.Context, params *ecs.CreateClusterInput, optFns ...func(*ecs.Options)) (*ecs.CreateClusterOutput, error)
}

type EcsClient struct {
	client EcsManager
}

func NewEcsClient(cfg aws.Config) *EcsClient {
	client := ecs.NewFromConfig(cfg)
	return &EcsClient{
		client: client,
	}
}

func (ec *EcsClient) GetColor(service string) string {
	return service + "-green"
}
func (ec *EcsClient) CreteCluster(ctx context.Context, cluster, service string) string {
	ec.client.CreateCluster(ctx, &ecs.CreateClusterInput{
		ClusterName: aws.String(cluster),
	})
	return service + "-green"
}
