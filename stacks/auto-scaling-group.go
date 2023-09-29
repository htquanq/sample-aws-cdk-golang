package stacks

import (
	"example-app-go/helper"
	"example-app-go/interfaces"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

type AutoScalingGroupsNestedStackProps struct {
	awscdk.NestedStackProps
	Vpc awsec2.IVpc
}

func AutoScalingGroupNestedStack(scope constructs.Construct, id string, AppConfig interfaces.AppConfig, props *AutoScalingGroupsNestedStackProps) (awscdk.NestedStack, awsautoscaling.AutoScalingGroup) {
	var autoScalingGroupProps awscdk.NestedStackProps
	if props != nil {
		autoScalingGroupProps = props.NestedStackProps
	}

	stack := awscdk.NewNestedStack(scope, &id, &autoScalingGroupProps)
	for i := 0; i < len(AppConfig.AutoScalingGroupObjects); i++ {
		instance := fmt.Sprintf("%s.%s", *&AppConfig.AutoScalingGroupObjects[i].InstanceType, *&AppConfig.AutoScalingGroupObjects[i].InstanceSize)
		isPublicSubnet := false
		if helper.SubnetType(AppConfig.AutoScalingGroupObjects[i].SubnetType) == awsec2.SubnetType_PUBLIC {
			isPublicSubnet = true
		}

		awsautoscaling.NewAutoScalingGroup(stack, aws.String(*&AppConfig.AutoScalingGroupObjects[i].Name),
			&awsautoscaling.AutoScalingGroupProps{
				AutoScalingGroupName:     &AppConfig.AutoScalingGroupObjects[i].Name,
				AllowAllOutbound:         &AppConfig.AutoScalingGroupObjects[i].Outbound,
				VpcSubnets:               &awsec2.SubnetSelection{SubnetType: helper.SubnetType(AppConfig.AutoScalingGroupObjects[i].SubnetType)},
				AssociatePublicIpAddress: aws.Bool(isPublicSubnet),
				DesiredCapacity:          &AppConfig.AutoScalingGroupObjects[i].DesiredSize,
				MinCapacity:              &AppConfig.AutoScalingGroupObjects[i].MinSize,
				MaxCapacity:              &AppConfig.AutoScalingGroupObjects[i].MaxSize,
				MachineImage:             helper.ImageMapper(*&AppConfig.AutoScalingGroupObjects[i].Image),
				InstanceType:             awsec2.NewInstanceType(&instance),
				HealthCheck: awsautoscaling.HealthCheck_Ec2(
					&awsautoscaling.Ec2HealthCheckOptions{
						Grace: awscdk.Duration_Minutes(aws.Float64(15)),
					},
				),
				Vpc: *&props.Vpc,
			},
		)
	}

	return stack, nil
}
