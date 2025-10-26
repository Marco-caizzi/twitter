package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/Marco-caizzi/twitter/awsgo"
	"github.com/Marco-caizzi/twitter/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dataSecret models.Secret

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(awsgo.Ctx, input)
	if err != nil {
		fmt.Println("Error retrieving secret:", err)
		return dataSecret, err
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		return dataSecret, fmt.Errorf("secret %s has no string value", secretName)
	}

	err = json.Unmarshal([]byte(secretString), &dataSecret)
	if err != nil {
		return dataSecret, err
	}
	fmt.Println(" > Successfully retrieved secret:", secretName)
	return dataSecret, nil
}
