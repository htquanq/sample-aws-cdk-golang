package helper

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

func ImageMapper(image string) awsec2.GenericLinuxImage {
	if strings.ToUpper(image) == "AMAZON_LINUX_2" {
		image := awsec2.AmazonLinuxImage(awsec2.NewAmazonLinuxImage(&awsec2.AmazonLinuxImageProps{
			Generation: awsec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
		}))
		return image
	} else if strings.ToUpper(image) == "UBUNTU" {
		image := awsec2.MachineImage_FromSsmParameter(jsii.String("/aws/service/canonical/ubuntu/server/focal/stable/current/amd64/hvm/ebs-gp2/ami-id"), &awsec2.SsmParameterImageOptions{
			Os: awsec2.OperatingSystemType_LINUX,
		})
		return image
	}
	return nil
}
