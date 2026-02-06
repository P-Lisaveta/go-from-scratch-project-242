package size

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize вычисляет суммарный размер файла или директории в байтах.
func GetPathSize(path string, all bool, recursive bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	// если это файл
	if !info.IsDir() {
		if !all && strings.HasPrefix(info.Name(), ".") {
			return 0, nil
		}
		return info.Size(), nil
	}

	// если это директория
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64

	for _, entry := range entries {
		name := entry.Name()

		// пропускаем скрытые файлы/директории, если --all не указан
		if !all && strings.HasPrefix(name, ".") {
			continue
		}

		fullPath := filepath.Join(path, name)

		if entry.IsDir() {
			if recursive {
				subSize, err := GetPathSize(fullPath, all, recursive)
				if err != nil {
					return 0, err
				}
				size += subSize
			}
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return 0, err
		}
		size += info.Size()
	}

	return size, nil
}

// FormatSize форматирует размер в байтах.
func FormatSize(size int64, human bool) string {
	if size < 0 {
		size = 0
	}

	if !human || size < 1024 {
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
