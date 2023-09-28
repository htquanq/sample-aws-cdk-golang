package helper

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func SubnetType(subnetType string) awsec2.SubnetType {
	if strings.ToLower(subnetType) == "public" {
		return awsec2.SubnetType_PUBLIC
	} else if strings.ToLower(subnetType) == "private_nat" {
		return awsec2.SubnetType_PRIVATE_WITH_NAT
	} else if strings.ToLower(subnetType) == "private_egress" {
		return awsec2.SubnetType_PRIVATE_WITH_EGRESS
	} else if strings.ToLower(subnetType) == "private_isolated" {
		return awsec2.SubnetType_PRIVATE_ISOLATED
	}
	return awsec2.SubnetType_PUBLIC
}

func CreateInternetGateway(CreateInternetGateway bool) bool {
	if !CreateInternetGateway {
		return false
	}
	return true
}

func MapIPPublicOnLaunch(MapIPPublicOnLaunch bool) bool {
	if !MapIPPublicOnLaunch {
		return false
	}
	return true
}
