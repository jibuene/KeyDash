package main

import (
	"bytes"
	"keydash/config"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_handleKeyVaultFlag(t *testing.T) {
	t.Run("Help", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer log.SetOutput(os.Stderr)

		handleKeyVaultFlag("help", &config.Config{}, []string{})

		if strings.Contains("Available keyvault commands:", buf.String()) {
			t.Errorf("expected: %s, got %s", "Available keyvault commands:", buf.String())
		}
	})
}
