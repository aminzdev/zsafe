package cmd

import (
	"github.com/spf13/cobra"
)

// RestoreCmd defines the "restore" subcommand that restores a backup to a specified directory.
var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a backup from a backup file to a destination directory",
	Long: `This command restores a backup from the specified backup file and extracts it to the specified destination directory.
It decrypts the backup (if encrypted), decompresses it, and recreates the original directory structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Restoring backup...")
	},
}

func init() {
	RootCmd.AddCommand(RestoreCmd)
}
