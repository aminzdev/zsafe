package cmd

import (
	"errors"
	"fmt"
	"github.com/aminzdev/zsafe/lib"
	"github.com/spf13/cobra"
	"path/filepath"
)

// RestoreCmd defines the "restore" subcommand that restores a backup to a specified directory.
var RestoreCmd = &cobra.Command{
	Use:   "restore [path] [destination]",
	Short: "Restore a backup from a backup file to a destination directory",
	Long: `This command restores a backup from the specified backup file and extracts it to the specified destination directory.
It decrypts the backup (if encrypted), decompresses it, and recreates the original directory structure.`,
	Args: cobra.ExactArgs(2), // Two arguments are required: backup file and destination path
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("Restoring backup...")

		// The backup file path and the extraction destination path
		backupPath := args[0]
		destinationPath := args[1]
		decryptPath := backupPath + ".dec"

		// Check if the backup file has the .enc extension and password is not provided
		if password == "" && filepath.Ext(backupPath) == ".enc" {
			return errors.New("the backup file is encrypted. Please provide a password to decrypt the file")
		}

		// If a password is provided, decrypt the file first
		if password != "" {
			// Decrypt the file
			derivedKey := lib.DeriveKey(password, []byte{})
			if len(derivedKey) != 32 {
				return errors.New("internal Error: Could not decrypt the backup file")
			}

			err := lib.DecryptFile(backupPath, decryptPath, string(derivedKey))
			if err != nil {
				return errors.New(fmt.Sprintf("Error decrypting file: %v\n", err))

			}
			cmd.Printf("Backup File decrypted successfully from %s\n", backupPath)
			backupPath = decryptPath
		}

		// After decryption, decompress the file
		err := lib.DecompressDirectory(backupPath, destinationPath)
		if err != nil {
			return errors.New(fmt.Sprintf("Error decompressing backup file: %v\n", err))
		}
		cmd.Printf("Backup restored successfully to %s\n", destinationPath)

		return nil
	},
}

func init() {
	RestoreCmd.Flags().StringVarP(&password, "password", "p", "", "Password to decrypt the backup file")
	RootCmd.AddCommand(RestoreCmd)
}
