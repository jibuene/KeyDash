package secretclient

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func FindSecret(secretClient *azsecrets.Client, secretName string) (string, string) {
	foundSecretID := ""
	foundSecret := ""
	pager := secretClient.NewListSecretsPager(nil)
	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, secret := range page.Value {
			if strings.HasPrefix(secret.ID.Name(), secretName) {
				foundSecretID = secret.ID.Name()
				foundSecret = GetSecret(secretClient, secret.ID.Name())
				break
			}
		}
		if foundSecretID != "" {
			break
		}
	}
	return foundSecretID, foundSecret
}

func ConnectToSecretClient(keyVaultName string) *azsecrets.Client {
	vaultURI := fmt.Sprintf("https://%s.vault.azure.net/", keyVaultName)

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	// Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(vaultURI, cred, nil)

	if err != nil {
		log.Fatalf("failed to create a Key Vault client: %v", err)
	}
	return client
}

func GetSecret(secretClient *azsecrets.Client, secretName string) string {
	secret, err := secretClient.GetSecret(context.Background(), secretName, "", nil)
	if err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}

	return *secret.Value
}
