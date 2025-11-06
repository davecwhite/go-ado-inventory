package main

import (
	"context"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

func getADOToken(tenant string, clientId string, clientSecret string) string {

	credential, err := confidential.NewCredFromSecret(clientSecret)
	if err != nil {
		panic(err)
	}

	confidentialClient, err := confidential.New("https://login.microsoftonline.com/"+tenant, clientId, credential)

	if err != nil {
		panic(err)
	}

	adoScopes := []string{"499b84ac-1321-427f-aa17-267ca6975798/.default"}
	adoResult, err := confidentialClient.AcquireTokenByCredential(context.TODO(), adoScopes)
	if err != nil {
		panic(err)
	}

	return adoResult.AccessToken
}
