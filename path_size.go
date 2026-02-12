package code

import (
	"fmt"
	"os"
	"strings"
)

type Options struct {
	Recursive bool
	Human     bool
	All       bool
}

func GetPathSize(path string, options Options) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	var size int64

	// Файл
	if !info.IsDir() {
		if !options.All && strings.HasPrefix(info.Name(), ".") {
			return 0, nil
		}
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		name := entry.Name()

		if !options.All && strings.HasPrefix(name, ".") {
			continue
		}

		fullPath := path + string(os.PathSeparator) + name

		if entry.IsDir() {
			if options.Recursive {
				subSize, err := GetPathSize(fullPath, options)
				if err != nil {
					return 0, err
				}
				size += subSize
			}
			continue
		}

		fInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}
		size += fInfo.Size()
	}

	return size, nil
}

func FormatSize(size int64, options Options) string {
	if size < 0 {
		size = 0
	}

	if !options.Human || size < 1024 {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	s := float64(size)
	i := 0

	for s >= 1024 && i < len(units)-1 {
		s /= 1024
		i++
	}

	return fmt.Sprintf("%.1f%s", s, units[i])
}
