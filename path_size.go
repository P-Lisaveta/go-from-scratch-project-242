package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "0B", err
	}

	var size int64
	if !info.IsDir() {
		if !all && strings.HasPrefix(info.Name(), ".") {
			return "0B", nil
		}
		// Читаем ТОЛЬКО содержимое файла, игнорируем метаданные
		f, err := os.Open(path)
		if err != nil {
			return "0B", err
		}
		defer f.Close()
		stat, err := f.Stat()
		if err != nil {
			return "0B", err
		}
		size = stat.Size() // Логический размер
		return fmt.Sprintf("%dB", size), nil
	}

	// Директории - суммируем размеры вложенных файлов
	entries, err := os.ReadDir(path)
	if err != nil {
		return "0B", err
	}

	for _, entry := range entries {
		name := entry.Name()
		if !all && strings.HasPrefix(name, ".") {
			continue
		}

		entryPath := filepath.Join(path, name)
		entrySize, err := GetPathSize(entryPath, recursive, human, all)
		if err != nil {
			return "0B", err
		}
		sizeStr, _ := strconv.ParseInt(entrySize, 10, 64)
		size += sizeStr
	}
	return fmt.Sprintf("%dB", size), nil
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
