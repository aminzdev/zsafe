package lib

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/zstd"
)

// CompressDirectory compresses a given directory into a .tar.zst file
func CompressDirectory(src string, dst string) error {
	outFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Initialize Zstandard writer
	zstdWriter, err := zstd.NewWriter(outFile, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
	if err != nil {
		return err
	}
	defer zstdWriter.Close()

	// Initialize tar writer
	tarWriter := tar.NewWriter(zstdWriter)
	defer tarWriter.Close()

	// Add files to the tar archive
	err = addFilesToTar(tarWriter, src, "")
	if err != nil {
		return err
	}

	return nil
}

// DecompressDirectory decompresses a .tar.zst file into the specified directory
func DecompressDirectory(src string, dst string) error {
	// Open the .tar.zst file
	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Initialize Zstandard reader
	zstdReader, err := zstd.NewReader(inFile)
	if err != nil {
		return err
	}
	defer zstdReader.Close()

	// Initialize tar reader
	tarReader := tar.NewReader(zstdReader)

	// Extract files from the tar archive
	for {
		// Read the next header
		header, err := tarReader.Next()
		if err == io.EOF {
			// End of archive
			break
		}
		if err != nil {
			return err
		}

		// Construct the full file path
		extractedPath := filepath.Join(dst, header.Name)

		// If it's a directory, create it
		if header.Typeflag == tar.TypeDir {
			err = os.MkdirAll(extractedPath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		// Create the file
		outFile, err := os.Create(extractedPath)
		if err != nil {
			return err
		}

		// Copy the content from the tar archive to the new file
		_, err = io.Copy(outFile, tarReader)
		if err != nil {
			return err
		}

		outFile.Close()
	}

	return nil
}

// addFilesToTar recursively adds files to the tar archive
func addFilesToTar(tw *tar.Writer, srcPath string, basePath string) error {
	return filepath.Walk(srcPath, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create the relative path inside the archive
		relPath, err := filepath.Rel(srcPath, file)
		if err != nil {
			return err
		}

		// Construct tar header
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		header.Name = filepath.Join(basePath, relPath) // Set correct path in tar archive

		// Write header to tar archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// If it's a directory, skip writing file data
		if info.IsDir() {
			return nil
		}

		// Open the file to copy its contents
		fileReader, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fileReader.Close()

		// Copy file data to tar writer
		_, err = io.Copy(tw, fileReader)
		if err != nil {
			return err
		}

		//fmt.Println("Added:", relPath)
		return nil
	})
}
