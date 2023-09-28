package main

import (
	"log"
	"os"
	"reflect"

	"github.com/aws/aws-cdk-go/awscdk/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	. "example-app-go/helper"
	"example-app-go/interfaces"
	"example-app-go/stacks"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ExampleAppGoStackProps struct {
	awscdk.StackProps
}

func NewExampleAppGoStack(scope constructs.Construct, id string, props *ExampleAppGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("ExampleAppGoQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	envName := app.Node().GetContext(jsii.String("environment"))
	yamlFile, err := os.ReadFile("./config/" + reflect.ValueOf(envName).String() + ".yaml")
	if err != nil {
		log.Fatal(err)
	}

	stack := NewExampleAppGoStack(app, "htquanqDemoStack", &ExampleAppGoStackProps{
		awscdk.StackProps{
			Env: env(Config(yamlFile)),
		},
	})
	// Create network stack
	_, networkStack := stacks.NetworkStack(stack, "htquanqDemoNetworkingStack", Config(yamlFile), nil)

	stacks.AutoScalingGroupNestedStack(stack, "htquanqDemoAutoScalingGroupStack", Config(yamlFile), &stacks.AutoScalingGroupsNestedStackProps{
		Vpc: *&networkStack,
	})
	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env(AppConfig interfaces.AppConfig) *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(AppConfig.AWSAccountID),
		Region:  jsii.String(AppConfig.AWSProfileRegion),
	}
}
