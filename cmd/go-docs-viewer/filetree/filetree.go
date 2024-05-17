package filetree

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FileEntry struct {
	Name     string
	Path     string
	IsDir    bool
	Contents []FileEntry
}

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func ScanDirs(path string) []FileEntry {
	out := []FileEntry{}
	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if IsDir(file) {
			out = append(out, FileEntry{
				Name:     filepath.Base(file),
				Path:     file,
				IsDir:    true,
				Contents: ScanDirs(file + "/*"),
			})
			continue
		}

		if strings.ToLower(file[len(file)-3:]) == ".md" {
			out = append(out, FileEntry{
				Name:     filepath.Base(file),
				Path:     file,
				IsDir:    false,
				Contents: []FileEntry{},
			})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].IsDir != out[j].IsDir {
			return true
		}
		return strings.Compare(out[i].Name, out[j].Name) < 0
	})
	return out
}

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(b), nil

}
