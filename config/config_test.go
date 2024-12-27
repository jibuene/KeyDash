package config_test

import (
	"fmt"
	"keydash/config"
	"os"
	"path/filepath"
	"testing"
)

func removeTestFile(filePath string) {
	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}
}
func TestConfig_AddKeyVault(t *testing.T) {
	t.Run("Add, Remove", func(t *testing.T) {
		testFilePath := filepath.Join(config.KEYVAULTSFILEPATH, "testkeyvaults.txt")
		c := config.InitConfig(testFilePath)

		for i := 0; i < 3; i++ {
			c.AddKeyVault("keyvault" + fmt.Sprint(i))
		}
		c.RemoveKeyVault("keyvault1")

		// Check if the keyvaults are added
		if len(c.KeyVaults) != 2 {
			t.Errorf("Expected 2 keyvaults, got %d", len(c.KeyVaults))
		}

		// Check if the keyvaults are added correctly
		if c.KeyVaults[0] != "keyvault0" {
			t.Errorf("Expected keyvault0, got %s", c.KeyVaults[0])
		}
		if c.KeyVaults[1] != "keyvault2" {
			t.Errorf("Expected keyvault2, got %s", c.KeyVaults[1])
		}

		// Cleanup: Remove the test file
		removeTestFile(testFilePath)
	})
}
