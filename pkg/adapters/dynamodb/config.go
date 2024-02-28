package dynamodb

type Config struct {
	AwsRegion   string `json:"awsRegion"`
	TablePrefix string `json:"tablePrefix"`
}
