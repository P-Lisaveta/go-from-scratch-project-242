package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	var size int64
	if !info.IsDir() {
		if !all && strings.HasPrefix(info.Name(), ".") {
			return 0, nil
		}
		// Читаем ТОЛЬКО содержимое файла, игнорируем метаданные
		f, err := os.Open(path)
		if err != nil {
			return 0, err
		}
		defer f.Close()
		stat, err := f.Stat()
		if err != nil {
			return 0, err
		}
		size = stat.Size() // Логический размер
		return size, nil
	}

	// Директории - суммируем размеры вложенных файлов
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		name := entry.Name()
		if !all && strings.HasPrefix(name, ".") {
			continue
		}

		entryPath := filepath.Join(path, name)
		entrySize, err := GetPathSize(entryPath, recursive, human, all)
		if err != nil {
			return 0, err
		}
		size += entrySize
	}
	return size, nil
}

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
