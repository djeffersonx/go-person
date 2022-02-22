package config

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var awsRegion = os.Getenv("aws.region")
var awsEndpoint = os.Getenv("aws.endpoint")

type SSMConfigProvider struct{}

func (appConfigProvider SSMConfigProvider) LoadConfiguration() {

	validateRequiredConfig("awsRegion", awsRegion)
	validateRequiredConfig("awsEndpoint", awsEndpoint)

	awsSession := getSession()
	ssmClient := ssm.New(awsSession)
	applicationParameters := getApplicationParametersKeys()

	getParametersResponse, err := ssmClient.GetParametersWithContext(context.Background(), &ssm.GetParametersInput{
		Names: aws.StringSlice(applicationParameters),
	})

	if err != nil {
		panic(err)
	}

	for _, param := range getParametersResponse.Parameters {
		if err := os.Setenv(*param.Name, *param.Value); err != nil {
			panic(err)
		}
	}

}

func getSession() *session.Session {

	ses, err := session.NewSessionWithOptions(session.Options{
		Profile: ApplicationProfile,
		Config: aws.Config{
			Region:   &awsRegion,
			Endpoint: &awsEndpoint,
		},
	})
	if err != nil {
		panic(err)
	}
	return ses
}

func getApplicationParametersKeys() []string {
	appConfigs := ApplicationConfigs
	var absoluteConfigKey []string
	for _, configKey := range appConfigs {
		absoluteConfigKey = append(
			absoluteConfigKey, createAbsoluteKey(configKey),
		)
	}
	return absoluteConfigKey
}

func createAbsoluteKey(configKey string) string {
	return fmt.Sprintf("%s/%s/%s", ApplicationName, ApplicationProfile, configKey)
}
