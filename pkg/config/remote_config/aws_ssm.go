package pkg

import (
	"context"
	"errors"
	"fmt"
	"go-person/pkg/config"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var applicationProfile = os.Getenv("application.profile")
var awsRegion = os.Getenv("aws.region")
var awsEndpoint = os.Getenv("aws.endpoint")

func init() {

	validateRequiredConfig("applicationProfile", applicationProfile)
	validateRequiredConfig("awsRegion", awsRegion)
	validateRequiredConfig("awsEndpoint", awsEndpoint)

	awsSession := getSession()
	ssmClient := ssm.New(awsSession)
	applicationParameters := getApplicationParameters()

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

func validateRequiredConfig(configName string, configValue string) {
	if configValue == "" {
		panic(errors.New(fmt.Sprintf("Configuraiton: %s is required", configName)))
	}
}

func getSession() *session.Session {

	ses, err := session.NewSessionWithOptions(session.Options{
		Profile: applicationProfile,
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

func getApplicationParameters() []string {
	appConfigs := config.Configuration
	var allConfigs []string
	for configKey := range appConfigs {
		allConfigs = append(
			allConfigs, createConfigKey(configKey),
		)
	}
	return allConfigs
}

func createConfigKey(k string) string {
	return fmt.Sprintf("%s/%s/%s", config.Configuration, applicationProfile, k)
}
