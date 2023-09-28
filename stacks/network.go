package stacks

import (
	"example-app-go/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

type NetworkNestedStackProps struct {
	awscdk.NestedStackProps
}

func NetworkStack(scope constructs.Construct, id string, AppConfig interfaces.AppConfig, props *NetworkNestedStackProps) (awscdk.NestedStack, awsec2.IVpc) {
	var networkProps awscdk.NestedStackProps
	if props != nil {
		networkProps = props.NestedStackProps
	}

	stack := awscdk.NewNestedStack(scope, &id, &networkProps)
	vpc := awsec2.NewVpc(stack, aws.String("demo-vpc"),
		&awsec2.VpcProps{
			IpAddresses: awsec2.IpAddresses_Cidr(aws.String(AppConfig.Cidr)),
			SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
				{
					CidrMask:            aws.Float64(28),
					SubnetType:          awsec2.SubnetType_PUBLIC,
					Name:                aws.String("public"),
					MapPublicIpOnLaunch: aws.Bool(true),
				},
			},
		},
	)

	return stack, vpc
}
