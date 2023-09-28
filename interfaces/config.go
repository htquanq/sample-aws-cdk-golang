package interfaces

type AppConfig struct {
	AWSAccountID     string `yaml:"accountId"`
	AWSProfileName   string `yaml:"awsProfile"`
	AWSProfileRegion string `yaml:"region"`
	VpcId            string `yaml:"vpcId"`
	Environment      string `yaml:"environment"`
	Cidr             string `yaml:"cidr"`
}
