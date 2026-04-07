package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		queue, err := sqs.NewQueue(ctx, "isn20261", nil)
		if err != nil {
			return err
		}

		ctx.Export("queue_name", queue.Name)
		ctx.Export("queue_url", queue.Url)
		ctx.Export("queue_arn", queue.Arn)

		return nil
	})
}
