package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createCognito(ctx *pulumi.Context, name string) (*cognito.UserPool, error) {
	userPool, err := cognito.NewUserPool(ctx, name+"-cognito-user-pool", &cognito.UserPoolArgs{
		Name: pulumi.String(name + "-cognito-user-pool"),
		PasswordPolicy: &cognito.UserPoolPasswordPolicyArgs{
			MinimumLength: pulumi.Int(8),
		},
	})
	if err != nil {
		return nil, err
	}
	return userPool, nil
}
