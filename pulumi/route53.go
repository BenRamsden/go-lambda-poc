package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateARecordArgs struct {
	hostedZoneId pulumi.StringOutput
	dist         *cloudfront.Distribution
}

func createARecord(ctx *pulumi.Context, name string, args *CreateARecordArgs) (*route53.Record, error) {
	record, err := route53.NewRecord(ctx, name+"-a-record", &route53.RecordArgs{
		Name:   pulumi.String("poc.sandbox.jugo.io"),
		ZoneId: args.hostedZoneId,
		Type:   pulumi.String("A"),
		Aliases: route53.RecordAliasArray{&route53.RecordAliasArgs{
			Name:                 args.dist.DomainName,
			EvaluateTargetHealth: pulumi.Bool(true),
			ZoneId:               args.dist.HostedZoneId,
		}},
	})
	if err != nil {
		return nil, err
	}
	return record, nil
}
