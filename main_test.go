package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_handleKeyVaultFlag(t *testing.T) {
	// We need to do this. Hack untill we refactor file operations to a class
	err := os.MkdirAll(KEYVAULTSFILEPATH, os.ModePerm)

	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}
	t.Run("Help", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer log.SetOutput(os.Stderr)

		handleKeyVaultFlag("help", []string{}, []string{})

		if strings.Contains("Available keyvault commands:", buf.String()) {
			t.Errorf("expected: %s, got %s", "Available keyvault commands:", buf.String())
		}
	})
	t.Run("Add,Remove,List", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		for i := 0; i < 3; i++ {
			handleKeyVaultFlag("add", []string{}, []string{"keyvault" + fmt.Sprint(i)})
		}
		handleKeyVaultFlag("remove", []string{}, []string{"keyvault1"})

		allKeyVaults := getAllKeyVaults()
		handleKeyVaultFlag("list", allKeyVaults, []string{})

		logOutput := buf.String()
		if !strings.Contains(logOutput, "keyvault0") || !strings.Contains(logOutput, "keyvault2") {
			t.Errorf("expected: %s, got %s", "keyvault0, keyvault2", logOutput)
		}

		// Cleanup
		handleKeyVaultFlag("remove", []string{}, []string{"keyvault0"})
		handleKeyVaultFlag("remove", []string{}, []string{"keyvault2"})

	})
}
