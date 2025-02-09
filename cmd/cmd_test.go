package cmd

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// TestRootCmd tests the root command
func TestRootCmd(t *testing.T) {
	// Create a buffer to capture the output of the root command
	var buf strings.Builder
	RootCmd.SetOut(&buf)

	// Explicitly set the arguments to be empty so that no subcommand is passed
	RootCmd.SetArgs([]string{})

	// Run the root command without specifying any subcommands
	err := RootCmd.Execute()

	// Assert that there is no error in executing the command
	assert.NoError(t, err)

	// Assert that the output contains the expected text
	assert.Contains(t, buf.String(), "Usage:")
	assert.Contains(t, buf.String(), "Available Commands:")
	assert.Contains(t, buf.String(), "Flags:")
}

// TestInvalidCommand tests the invalid command behavior (e.g., an unrecognized command)
func TestInvalidCommand(t *testing.T) {
	// Run an invalid command
	RootCmd.SetArgs([]string{"invalid"})
	err := RootCmd.Execute()

	// Assert that an error occurred
	assert.Error(t, err)

	// Assert that the output contains an error message
	assert.Contains(t, err.Error(), "unknown command")
}
