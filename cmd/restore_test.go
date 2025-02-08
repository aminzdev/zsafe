package cmd

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRestoreCmd checks the functionality of the restore command.
func TestRestoreCmd(t *testing.T) {
	// Test case setup
	sourceDir := "./testdata"
	restoreDir := "./testrestore"
	backupFile := "backup.zsf"
	encryptedFile := backupFile + ".enc"
	decryptedFile := encryptedFile + ".dec"
	testPassword := "testpassword"

	defer os.Remove(backupFile)
	defer os.Remove(encryptedFile)
	defer os.Remove(decryptedFile)

	// Ensure the test directory exists for the test
	err := os.MkdirAll(sourceDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(sourceDir)
	defer os.RemoveAll(restoreDir)

	// Create a dummy file to test
	testFile := sourceDir + "/testfile.txt"
	err = os.WriteFile(testFile, []byte("Test data for backup"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	RootCmd.SetArgs([]string{"backup", sourceDir, "--password", testPassword})

	// Run the command
	err = RootCmd.Execute()
	if err != nil {
		t.Fatalf("Error running command: %v", err)
	}

	tests := []struct {
		name           string
		args           []string
		password       string
		expectedOutput string
		expectError    bool
	}{
		{
			name:           "successful restore without encryption",
			args:           []string{"restore", backupFile, restoreDir},
			password:       "",
			expectedOutput: fmt.Sprintf("Restoring backup...\nBackup restored successfully to %s\n", restoreDir),
			expectError:    false,
		},
		{
			name:           "error when no password for encrypted file",
			args:           []string{"restore", encryptedFile, restoreDir},
			password:       "",
			expectedOutput: "the backup file is encrypted. Please provide a password to decrypt the file",
			expectError:    true,
		},
		{
			name:           "successful restore with password",
			args:           []string{"restore", encryptedFile, restoreDir},
			password:       testPassword,
			expectedOutput: fmt.Sprintf("Restoring backup...\nBackup File decrypted successfully from %s\nBackup restored successfully to %s\n", encryptedFile, restoreDir),
			expectError:    false,
		},
		{
			name:           "error when insufficient arguments",
			args:           []string{"restore", backupFile},
			password:       "",
			expectedOutput: "Error: accepts 2 arg(s), received 1\n",
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture command output
			var buf bytes.Buffer
			RootCmd.SetOut(&buf)
			RootCmd.SetErr(&buf)

			// Execute command with arguments
			args := append(tt.args, "--password", tt.password)
			RootCmd.SetArgs(args)
			err := RootCmd.Execute()

			// Check error expectation
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Check command output
			assert.Contains(t, buf.String(), tt.expectedOutput)
		})
	}
}
