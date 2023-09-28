package interfaces

type AutoScalingGroup struct {
	Name              string  `yaml:"name"`
	MaxSize           float64 `yaml:"maxSize"`
	MinSize           float64 `yaml:"minSize"`
	DesiredSize       float64 `yaml:"desiredSize"`
	Outbound          bool    `yaml:"outbound"`
	AssociatePublicIP bool    `yaml:"associatePublicIP"`
	InstanceType      string  `yaml:"instanceType"`
	InstanceSize      string  `yaml:"instanceSize"`
	Image             string  `yaml:"image"`
}

type AppConfig struct {
	AWSAccountID            string             `yaml:"accountId"`
	AWSProfileName          string             `yaml:"awsProfile"`
	AWSProfileRegion        string             `yaml:"region"`
	VpcId                   string             `yaml:"vpcId"`
	Environment             string             `yaml:"environment"`
	Cidr                    string             `yaml:"cidr"`
	AutoScalingGroupObjects []AutoScalingGroup `yaml:"scalingGroup"`
}
