package main

import (
	"github.com/pulumi/pulumi-aws-apigateway/sdk/go/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create IAM role for Lambda
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Principal": {
						"Service": "lambda.amazonaws.com"
					}
				}]
			}`),
			ManagedPolicyArns: pulumi.StringArray{
				iam.ManagedPolicyAWSLambdaBasicExecutionRole,
			},
		})
		if err != nil {
			return err
		}

		// Create Lambda function
		fn, err := lambda.NewFunction(ctx, "fn", &lambda.FunctionArgs{
			Runtime: pulumi.String("go1.x"),
			Handler: pulumi.String("handler"),
			Role:    role.Arn,
			Code:    pulumi.NewFileArchive("./function"),
		})
		if err != nil {
			return err
		}

		// Create API Gateway
		localPath := "www"
		method := apigateway.MethodGET
		api, err := apigateway.NewRestAPI(ctx, "api", &apigateway.RestAPIArgs{
			Routes: []apigateway.RouteArgs{
				{
					Path:      "/",
					LocalPath: &localPath,
				},
				{
					Path:         "/date",
					Method:       &method,
					EventHandler: fn,
				},
			},
		})
		if err != nil {
			return err
		}

		// Export the API URL
		ctx.Export("url", api.Url)

		return nil
	})
}
