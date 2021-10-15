package authentication_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/efarrer/page2pod/authentication"
	"github.com/efarrer/page2pod/authentication/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAuthenticate_ReturnsAuthErrorIfCantGetValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	secretValueGetter := mocks.NewMockSecretValueGetter(ctrl)
	secretValueGetter.EXPECT().GetSecretValue(gomock.Any()).Return(nil, errors.New("Some error"))

	err := authentication.Authenticate(secretValueGetter, "", "")
	require.Error(t, err)
}

func TestAuthenticate_ReturnsAuthErrorIfSecretNotAString(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	secretValueGetter := mocks.NewMockSecretValueGetter(ctrl)
	secretValueGetter.EXPECT().GetSecretValue(gomock.Any()).Return(&secretsmanager.GetSecretValueOutput{}, nil)

	err := authentication.Authenticate(secretValueGetter, "", "")
	require.Error(t, err)
}

func TestAuthenticate_ReturnsAuthErrorIfCantUnMarshalResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	username := "username"
	password := "some secret"
	bytes, err := json.Marshal("")
	require.NoError(t, err)

	secretValueGetter := mocks.NewMockSecretValueGetter(ctrl)
	secretValueGetter.EXPECT().GetSecretValue(gomock.Any()).Return(&secretsmanager.GetSecretValueOutput{
		SecretString: aws.String(string(bytes)),
	}, nil)

	err = authentication.Authenticate(secretValueGetter, username, password)
	require.Error(t, err)
}

func TestAuthenticate_ReturnsAuthErrorIfBadPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	username := "username"
	password := "some secret"
	badPassword := "bad password"
	bytes, err := json.Marshal(map[string]string{username: password})
	require.NoError(t, err)

	secretValueGetter := mocks.NewMockSecretValueGetter(ctrl)
	secretValueGetter.EXPECT().GetSecretValue(gomock.Any()).Return(&secretsmanager.GetSecretValueOutput{
		SecretString: aws.String(string(bytes)),
	}, nil)

	err = authentication.Authenticate(secretValueGetter, username, badPassword)
	require.Error(t, err)
}

func TestAuthenticate_ReturnsAuthErrorIfBadUsernameAndEmptyPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	username := "username"
	password := "some secret"
	badUsername := "bad username"
	emptyPassword := ""
	bytes, err := json.Marshal(map[string]string{username: password})
	require.NoError(t, err)

	secretValueGetter := mocks.NewMockSecretValueGetter(ctrl)
	secretValueGetter.EXPECT().GetSecretValue(gomock.Any()).Return(&secretsmanager.GetSecretValueOutput{
		SecretString: aws.String(string(bytes)),
	}, nil)

	err = authentication.Authenticate(secretValueGetter, badUsername, emptyPassword)
	require.Error(t, err)
}

func TestAuthenticate_ReturnsNilIfAuthenticated(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	username := "username"
	password := "some secret"
	bytes, err := json.Marshal(map[string]string{username: password})
	require.NoError(t, err)

	secretValueGetter := mocks.NewMockSecretValueGetter(ctrl)
	secretValueGetter.EXPECT().GetSecretValue(gomock.Any()).Return(&secretsmanager.GetSecretValueOutput{
		SecretString: aws.String(string(bytes)),
	}, nil)

	err = authentication.Authenticate(secretValueGetter, username, password)
	require.NoError(t, err)
}
