package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func backupData(archiveName string) {
	outFile, err := os.Create(archiveName)
	if err != nil {
		log.Fatalf("Failed to create archive: %v", err)
	}
	defer outFile.Close()

	gzWriter := gzip.NewWriter(outFile)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	root := "."
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the archive file itself
		if info.Name() == archiveName {
			return nil
		}

		hdr, err := tar.FileInfoHeader(info, path)
		if err != nil {
			return err
		}

		// Set correct name in archive (relative path)
		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if relPath == "." {
			return nil // skip root dir entry
		}
		hdr.Name = relPath

		// Try to preserve ownership (if running as root)
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			hdr.Uid = int(stat.Uid)
			hdr.Gid = int(stat.Gid)
		}

		if err := tarWriter.WriteHeader(hdr); err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			if _, err := io.Copy(tarWriter, file); err != nil {
				return err
			}
			os.Printf(".")
		}

		os.Println()
		return nil
	})
}
