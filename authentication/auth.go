package authentication

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var authError = errors.New("Unable to authenticate")

type SecretValueGetter interface {
	GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
}

func Authenticate(svc SecretValueGetter, username, password string) error {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("credentials"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		fmt.Printf("Unable to get secret value %v\n", err)
		return authError
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	if result.SecretString == nil {
		fmt.Println("Expecting a secret string, but got binary data")
		return authError
	}

	mapping := map[string]string{}
	err = json.Unmarshal([]byte(*result.SecretString), &mapping)
	if err != nil {
		fmt.Println("Unable to unmarshal data")
		return authError
	}

	// Your code goes here.
	if expectedPass, ok := mapping[username]; !ok || expectedPass != password {
		return authError
	}

	fmt.Printf("Authenticated " + username)
	return nil
}
