package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"os"
)

// EncryptFile encrypts a file using AES-256 in CFB mode
func EncryptFile(inputPath, outputPath, key string) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Create AES cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	// Generate IV and write it to the output file
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}
	if _, err := outFile.Write(iv); err != nil {
		return err
	}

	// Create a cipher stream writer
	stream := cipher.NewCFBEncrypter(block, iv)
	writer := &cipher.StreamWriter{S: stream, W: outFile}

	// Copy data from input to encrypted output
	_, err = io.Copy(writer, inFile)
	if err != nil {
		return err
	}

	return nil
}

// DecryptFile decrypts a file using AES-256 in CFB mode
func DecryptFile(inputPath, outputPath, key string) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Read IV from the input file
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(inFile, iv); err != nil {
		return err
	}

	// Create AES cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	// Create a cipher stream reader for decryption
	stream := cipher.NewCFBDecrypter(block, iv)
	reader := &cipher.StreamReader{S: stream, R: inFile}

	// Copy decrypted data from the input to the output file
	_, err = io.Copy(outFile, reader)
	if err != nil {
		return err
	}

	return nil
}

// DeriveKey generates a 32-byte key from a password using PBKDF2
func DeriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, 100_000, 32, sha256.New)
}
