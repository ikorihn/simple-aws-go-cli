package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/script/awsclient"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/bitfield/script"
	"github.com/urfave/cli/v2"
)

type Runner struct {
	cfg aws.Config
}

func (r *Runner) Ecs() *cli.Command {
	var cluster, service, color string

	return &cli.Command{
		Name:  "ecs",
		Usage: "ecs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "cluster",
				Aliases:     []string{"c"},
				Destination: &cluster,
			},
			&cli.StringFlag{
				Name:        "service",
				Aliases:     []string{"s"},
				Destination: &service,
			},
			&cli.StringFlag{
				Name:        "color",
				Aliases:     []string{"c"},
				Destination: &color,
			},
		},
		Subcommands: []*cli.Command{
			{
				Name: "create-cluster",
				Action: func(cc *cli.Context) error {
					fmt.Println("start")
					req, err := http.NewRequest(http.MethodGet, "https://httpbin.org/get?name=John", nil)
					if err != nil {
						return err
					}

					res, err := script.Do(req).JQ(".args").Stdout()
					if err != nil {
						return err
					}

					fmt.Println(res)

					ecsClient := awsclient.NewEcsClient(r.cfg)
					ecsClient.GetColor("foo")

					return nil

				},
			},
		},
	}
}

func main() {
	var account string
	var region string
	var iamRole string

	var runner *Runner

	app := &cli.App{
		Name:  "script",
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "account",
				Usage:       "AWS account id",
				Destination: &account,
			},
			&cli.StringFlag{
				Name:        "region",
				Value:       "ap-northeast-1",
				Usage:       "AWS region",
				Destination: &region,
			},
			&cli.StringFlag{
				Name:        "iam-role",
				Usage:       "AWS IAM Role",
				Destination: &iamRole,
			},
		},
		Before: func(cc *cli.Context) error {
			ctx := context.Background()
			var err error
			cfg, err := awsclient.NewConfig(ctx, region)
			if err != nil {
				return err
			}

			stsc := awsclient.NewStsClient(cfg)
			stsc.SetCredentialProvider(ctx, &cfg, account, iamRole)

			runner = &Runner{
				cfg: cfg,
			}

			return nil
		},

		Commands: []*cli.Command{
			runner.Ecs(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
