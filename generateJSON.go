package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
)

func generateJSONFunction(keyVaultName string, defaultPath string) {

	var retrivedList Secrets

	vaultURL := fmt.Sprintf("https://%s.vault.azure.net/", keyVaultName)

	// Create a new DefaultAzureCredential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to obtain a credential: %v", err)
	}

	// Create a new client
	client, err := azsecrets.NewClient(vaultURL, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// List secrets
	pager := client.NewListSecretPropertiesPager(nil)

	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			log.Fatalf("Failed to get next page of secrets: %v", err)
		}

		for _, secret := range page.Value {
			newSecret := Secret{KeyVaultSecretName: secret.ID.Name(), VaultSecretPath: defaultPath, VaultSecretName: secret.ID.Name(), VaultSecretKey: "secret", Copy: false}
			retrivedList.Secrets = append(retrivedList.Secrets, newSecret)
		}
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(retrivedList, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal secrets to JSON: %v", err)
	}

	// Write to file
	file, err := os.Create("secrets.json")
	if err != nil {
		log.Fatalf("Failed to create JSON file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to write JSON data to file: %v", err)
	}

	fmt.Println("List of secrets in Keyvault has been written to secrets.json")
}
