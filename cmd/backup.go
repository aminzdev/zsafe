package cmd

import (
	"errors"
	"fmt"
	"github.com/aminzdev/zsafe/lib"
	"github.com/spf13/cobra"
)

// BackupCmd defines the "backup" subcommand that creates a backup.
var BackupCmd = &cobra.Command{
	Use:   "backup [path]",
	Short: "Create a backup from a source directory",
	Long: `This command creates a backup of the specified directory.
You can optionally specify a compression level, and password for security.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("Creating backup...")

		if output == "" {
			output = "backup"
		}

		sourcePath := args[0]
		backupPath := output + ".zsf"
		encryptedPath := backupPath + ".enc"

		// Step 1: Compress the directory
		err := lib.CompressDirectory(sourcePath, backupPath)
		if err != nil {
			return err
		}
		cmd.Printf("Backup File created successfully at %s\n", backupPath)

		// Step 2: Encrypt the compressed file if password is provided
		if password != "" {
			// Passing an empty slice for salt, so the derived key is only based on the password
			derivedKey := lib.DeriveKey(password, []byte{}) // No salt
			if len(derivedKey) != 32 {
				return errors.New("internal Error: Could not encrypt the compressed file")
			}

			err = lib.EncryptFile(backupPath, encryptedPath, string(derivedKey))
			if err != nil {
				return errors.New(fmt.Sprintf("Error encrypting file: %v", err))
			}
			cmd.Printf("Backup File encrypted successfully at %s\n", encryptedPath)
		}
		return nil
	},
}

func init() {
	BackupCmd.Flags().StringVarP(&output, "output", "o", "", "Name for the backup file")
	BackupCmd.Flags().StringVarP(&password, "password", "p", "", "Password to encrypt the backup file")
	RootCmd.AddCommand(BackupCmd)
}
