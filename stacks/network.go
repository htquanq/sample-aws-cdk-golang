package stacks

import (
	"example-app-go/helper"
	"example-app-go/interfaces"
	"fmt"

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
	// Create vpc with public subnets and private_isolated subnets
	vpc := awsec2.NewVpc(stack, aws.String(AppConfig.Vpc.Name),
		&awsec2.VpcProps{
			IpAddresses:           awsec2.IpAddresses_Cidr(aws.String(AppConfig.Vpc.Cidr)),
			CreateInternetGateway: aws.Bool(helper.CreateInternetGateway(AppConfig.Vpc.InternetGateway)),
			EnableDnsHostnames:    aws.Bool(true),
			EnableDnsSupport:      aws.Bool(true),
			VpcName:               aws.String(AppConfig.Vpc.Name),
			MaxAzs:                aws.Float64(3),
			SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
				{
					CidrMask:            aws.Float64(AppConfig.Vpc.CidrMask),
					SubnetType:          awsec2.SubnetType_PUBLIC,
					Name:                aws.String(fmt.Sprintf("%s-%s", AppConfig.Vpc.Name, "public")),
					MapPublicIpOnLaunch: aws.Bool(true),
				},
				{
					CidrMask:   aws.Float64(AppConfig.Vpc.CidrMask),
					SubnetType: awsec2.SubnetType_PRIVATE_ISOLATED,
					Name:       aws.String(fmt.Sprintf("%s-%s", AppConfig.Vpc.Name, "private")),
				},
			},
		},
	)

	return stack, vpc
}
