package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

var KEYVAULTSFILE = "keyvaults.txt"

func main() {
	allKeyVaults := getAllKeyVaults()

	keyVaultFlag := flag.String("keyvault", "notset", "Keyvault specific commands.")
	secretNameFlag := flag.String("secret", "", "The name of the secret to retrieve.")
	flag.Parse()

	if *keyVaultFlag != "notset" {
		handleKeyVaultFlag(*keyVaultFlag, allKeyVaults, flag.Args())
		return
	}

	if len(allKeyVaults) == 0 {
		log.Fatal("No keyvaults found. Use --keyvault help to see options.")
	}

	if *secretNameFlag == "" {
		log.Fatal("Secret name is required. Use --secret <secret-name>.")
	}

	secret := ""
	for _, keyVault := range allKeyVaults {
		client := connectToSecretClient(keyVault)
		secret = findSecret(client, *secretNameFlag)
		if secret != "" {
			break
		}
	}

	if secret == "" {
		log.Fatalf("Secret %s not found in any keyvaults.", *secretNameFlag)
	}

	fmt.Printf("Secret: %s\n", secret)
}

// handleKeyVaultFlag handles the keyvault flag passed to the program.
// It can add, remove, list or show help for keyvaults.
func handleKeyVaultFlag(keyVaultFlag string, allKeyVaults []string, extraArgs []string) {
	switch keyVaultFlag {
	case "help":
		log.Print("Available keyvault commands: 'help', 'add', 'list', 'remove'")
	case "add":
		if len(extraArgs) == 0 {
			log.Fatal("Usage of --keyvault add: --keyvault add <keyvault-name>")
		}
		addKeyVault(extraArgs[0])
	case "remove":
		if len(extraArgs) == 0 {
			log.Fatal("Usage of --keyvault remove: --keyvault remove <keyvault-name>")
		}
		removeKeyVault(extraArgs[0])
	case "list":
		log.Printf("Listing all keyvaults found:")
		for _, keyVault := range allKeyVaults {
			log.Printf("    - %s", keyVault)
		}
	case "notset":
	default:
		log.Fatalf(`Invalid keyvault command: %s`, keyVaultFlag)
	}
}

// getAllKeyVaults reads the keyvaults file and returns a slice of keyvault names.
func getAllKeyVaults() []string {
	file := openKeyVaultFile()
	defer file.Close()
	keyVaults := []string{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		keyVaults = append(keyVaults, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file %s: %v", KEYVAULTSFILE, err)
	}

	return keyVaults
}

// addKeyVault appends a keyvault name to the keyvaults file.
func addKeyVault(keyVaultName string) {
	file := openKeyVaultFile()
	defer file.Close()
	_, err := file.WriteString(keyVaultName + " ")
	if err != nil {
		log.Fatal(err)
	}
}

// removeKeyVault removes a keyvault name from the keyvaults file.
func removeKeyVault(keyVaultName string) {
	file := openKeyVaultFile()
	defer file.Close()

	tempFile, err := os.Create("tempfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer tempFile.Close()
	writer := bufio.NewWriter(tempFile)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if scanner.Text() == keyVaultName {
			continue
		}
		_, err = writer.WriteString(scanner.Text() + " ")
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	writer.Flush()

	err = os.Rename("tempfile.txt", KEYVAULTSFILE)
	if err != nil {
		log.Fatal(err)
	}
}

// openKeyVaultFile opens the keyvaults file in append mode.
// If the file does not exist, it will be created.
func openKeyVaultFile() *os.File {
	file, err := os.OpenFile(KEYVAULTSFILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func findSecret(secretClient *azsecrets.Client, secretName string) string {
	foundSecret := ""
	pager := secretClient.NewListSecretsPager(nil)
	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, secret := range page.Value {
			if strings.HasPrefix(secret.ID.Name(), secretName) {
				foundSecret = getSecret(secretClient, secret.ID.Name())
				break
			}
		}
		if foundSecret != "" {
			break
		}
	}
	return foundSecret
}

func connectToSecretClient(keyVaultName string) *azsecrets.Client {
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

func getSecret(secretClient *azsecrets.Client, secretName string) string {
	secret, err := secretClient.GetSecret(context.Background(), secretName, "", nil)
	if err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}

	return *secret.Value
}
