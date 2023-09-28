# Sample code to create VPC

## Input

```yaml
# config/demo.yaml
"vpcId": "vpc-0ece9a0b3eda519e9"
"accountId": "756007437776"
"region": "ap-southeast-1"
"awsProfile": "default"
"environment": "demo"
"cidr": "10.23.0.0/16"
```

## Ouput

This CDK will create a main stack and nested stacks (NetworkingStack this time)

Why I don't put everything into a single stack:
* Hard to maintain code base since it can be very long
* Made changes to single stack can affect other resources that I don't want to modify

The reason I develop it in Go instead of Typescript while I still can do it is I work with Kubernetes and Helm so Go is more familiar to me and I want to unite the language I use

![Created stacks](.images/stacks.png)
![Created subnets](.images/subnets.png)
![Created vpc](.images/vpc.png)
## Useful commands

 * `cdk deploy --context environment=demo`      deploy this stack to your default AWS account/region
 * `cdk diff --context environment=demo`        compare deployed stack with current state
 * `cdk synth --context environment=demo`       emits the synthesized CloudFormation template
