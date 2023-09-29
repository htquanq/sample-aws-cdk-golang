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
	SubnetType        string  `yaml:"subnetType"`
}

type VpcConfig struct {
	Name            string  `yaml:"name"`
	Cidr            string  `yaml:"cidr"`
	CidrMask        float64 `yaml:"cidrMask"`
	InternetGateway bool    `yaml:"internetGateway"`
}

type AppConfig struct {
	AWSAccountID            string             `yaml:"accountId"`
	AWSProfileName          string             `yaml:"awsProfile"`
	AWSProfileRegion        string             `yaml:"region"`
	Environment             string             `yaml:"environment"`
	AutoScalingGroupObjects []AutoScalingGroup `yaml:"scalingGroup"`
	Vpc                     VpcConfig          `yaml:"vpc"`
}
