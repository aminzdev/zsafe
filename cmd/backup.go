package cmd

import (
	"github.com/spf13/cobra"
)

// BackupCmd defines the "backup" subcommand that creates a backup.
var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Create a backup from a source directory to a destination",
	Long: `This command creates a backup of files or directories from the specified source.
You can optionally specify a compression level, parallel processing, and encryption key for security.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Creating backup...")
	},
}

func init() {
	RootCmd.AddCommand(BackupCmd)
}
