package cmd

import (
	"github.com/spf13/cobra"
)

var (
	output   string
	password string
)

// RootCmd is the main command that sets up the basic CLI structure and command handling.
var RootCmd = &cobra.Command{
	Use:   "zsafe",
	Short: "A simple backup manager CLI for creating and restoring backups",
	Long: `ZSafe allows you to create and restore backups of files or directories.
You can compress and encrypt backups for security and restore them later.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
