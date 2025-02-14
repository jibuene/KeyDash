package config

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	FilePath  string
	KeyVaults []string
}

var KEYVAULTSFILE = "keyvaults.txt"

var KEYVAULTSFILEPATH = getConfigFilePath()

// Full path to the keyvaults file
var KEYVAULTSFILEFQDN = filepath.Join(KEYVAULTSFILEPATH, KEYVAULTSFILE)

// openKeyVaultFile opens the keyvaults file in append mode.
// If the file does not exist, it will be created.
func openKeyVaultFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func InitConfig(path string) Config {
	// Create config directory if it does not exist
	err := os.MkdirAll(KEYVAULTSFILEPATH, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	file := openKeyVaultFile(path)
	defer file.Close()
	keyVaults := []string{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		keyVaults = append(keyVaults, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file %s: %v", path, err)
	}

	return Config{FilePath: path, KeyVaults: keyVaults}
}

// addKeyVault appends a keyvault name to the keyvaults file.
func (c *Config) AddKeyVault(keyVaultName string) {
	file := openKeyVaultFile(c.FilePath)
	defer file.Close()
	_, err := file.WriteString(keyVaultName + " ")
	if err != nil {
		log.Fatal(err)
	}
	c.KeyVaults = append(c.KeyVaults, keyVaultName)
}

// removeKeyVault removes a keyvault name from the keyvaults file.
func (c *Config) RemoveKeyVault(keyVaultName string) {
	file := openKeyVaultFile(c.FilePath)
	newKeyVaults := []string{}
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
		newKeyVaults = append(newKeyVaults, scanner.Text())
		_, err = writer.WriteString(scanner.Text() + " ")
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	writer.Flush()

	err = os.Rename("tempfile.txt", c.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	c.KeyVaults = newKeyVaults
}

// getConfigFilePath returns the path of the config file.
// This file is used to store the keyvault names.
func getConfigFilePath() string {
	var homeDir string

	switch runtime.GOOS {
	case "windows":
		homeDir = os.Getenv("USERPROFILE")
		if homeDir == "" {
			log.Fatal("Could not determine home directory on Windows")
		}
	case "linux":
		homeDir = os.Getenv("HOME")
		if homeDir == "" {
			log.Fatal("Could not determine home directory on Linux")
		}
	case "darwin":
		homeDir = os.Getenv("HOME")
		if homeDir == "" {
			log.Fatal("Could not determine home directory on MacOS")
		}
	default:
		log.Fatalf("Unsupported OS: %s", runtime.GOOS)
	}

	configFilePath := filepath.Join(homeDir, ".KeyDash")
	return configFilePath
}
