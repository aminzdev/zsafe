package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBackupCmd(t *testing.T) {
	// Test case setup
	sourceDir := "./testdata" // Test directory with some test files
	backupFile := "backup.zsf"
	encryptedFile := backupFile + ".enc"
	password := "testpassword"

	defer os.Remove(backupFile)
	defer os.Remove(encryptedFile)

	// Ensure the test directory exists for the test
	err := os.MkdirAll(sourceDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(sourceDir)

	// Create a dummy file to test
	testFile := sourceDir + "/testfile.txt"
	err = os.WriteFile(testFile, []byte("Test data for backup"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	// Create a buffer to capture the command output
	var buf bytes.Buffer
	RootCmd.SetOut(&buf)

	// Simulate running the backup command
	RootCmd.SetArgs([]string{"backup", sourceDir, "--password", password})

	// Run the command
	err = RootCmd.Execute()
	if err != nil {
		t.Fatalf("Error running command: %v", err)
	}

	// Check that the output contains success messages
	output := buf.String()
	assert.Contains(t, output, "Backup File created successfully")
	assert.Contains(t, output, "Backup File encrypted successfully")

	// Check that the backup and encrypted files exist
	_, err = os.Stat(backupFile)
	assert.NoError(t, err, "Backup file should exist")

	// Check that the encrypted file exists if a password was provided
	_, err = os.Stat(encryptedFile)
	assert.NoError(t, err, "Encrypted backup file should exist")

	// Clean up the generated files
	os.Remove(backupFile)
	os.Remove(encryptedFile)
}
