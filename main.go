package main

import (
	"flag"
	"fmt"
	"keydash/config"
	"keydash/secretclient"
	"log"
)

func main() {
	keyVaultConfig := config.InitConfig(config.KEYVAULTSFILEFQDN)

	keyVaultFlag := flag.String("keyvault", "notset", "Keyvault specific commands.")
	secretNameFlag := flag.String("secret", "", "The name of the secret to retrieve.")
	flag.Parse()
	extraArgs := flag.Args()

	if (len(extraArgs) > 0 && extraArgs[0] == "help") || (len(extraArgs) == 0 && *keyVaultFlag == "notset") {
		log.Print("Usage: keydash [--keyvault <command>] [--secret <secret-name>]")
		log.Print("Use `--keyvault help` to see keyvault commands.")
		log.Print("Example usage:")
		log.Print("    keydash --keyvault add mykeyvault // Adds a keyvault to the config file.")
		log.Print("    keydash --secret mysecret         // Retrieves the secret named 'mysecret'")
		log.Print("    keydash secret                    // Retrieves the secret named 'secret'")
		return
	}

	if *keyVaultFlag != "notset" {
		handleKeyVaultFlag(*keyVaultFlag, &keyVaultConfig, extraArgs)
		return
	}

	if len(keyVaultConfig.KeyVaults) == 0 {
		log.Fatal("No keyvaults found. Use `--keyvault help` to see options.")
	}

	secretToFind := ""

	if len(extraArgs) > 0 && extraArgs[0] != "" {
		secretToFind = extraArgs[0]
	}
	if *secretNameFlag != "" {
		secretToFind = *secretNameFlag
	}

	if secretToFind == "" {
		log.Fatal("Secret name is required. Use --secret <secret-name>.")
	}

	secret := ""
	for _, keyVault := range keyVaultConfig.KeyVaults {
		client := secretclient.ConnectToSecretClient(keyVault)
		secret = secretclient.FindSecret(client, secretToFind)
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
func handleKeyVaultFlag(keyVaultFlag string, keyVaultConfig *config.Config, extraArgs []string) {
	switch keyVaultFlag {
	case "help":
		log.Print("Available keyvault commands: 'help', 'add', 'list', 'remove'")
	case "add":
		if len(extraArgs) == 0 {
			log.Fatal("Usage of --keyvault add: --keyvault add <keyvault-name>")
		}
		keyVaultConfig.AddKeyVault(extraArgs[0])
	case "remove":
		if len(extraArgs) == 0 {
			log.Fatal("Usage of --keyvault remove: --keyvault remove <keyvault-name>")
		}
		keyVaultConfig.RemoveKeyVault(extraArgs[0])
	case "list":
		log.Printf("Listing all keyvaults found:")
		for _, keyVault := range keyVaultConfig.KeyVaults {
			log.Printf("    - %s", keyVault)
		}
	case "notset":
	default:
		log.Fatalf(`Invalid keyvault command: %s`, keyVaultFlag)
	}
}
