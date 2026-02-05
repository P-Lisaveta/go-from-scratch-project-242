package size

import (
	"fmt"
	"os"
	"strings"
)

// GetPathSize вычисляет суммарный размер файла или (неглубоко) директории в байтах.
func GetPathSize(path string, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	var size int64
	if !info.IsDir() {
		if !all && strings.HasPrefix(info.Name(), ".") {
			return 0, nil
		}
		size = info.Size()
	} else {
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			if !all && strings.HasPrefix(info.Name(), ".") {
				continue
			}
			fInfo, err := entry.Info()
			if err != nil {
				return 0, err
			}
			size += fInfo.Size()
		}
	}
	return size, nil
}

// FormatSize форматирует размер в байтах.
// Если human == false, возвращает строку вида "123B".
// Если human == true, конвертирует в человекочитаемый формат (B, KB, MB, GB, TB, PB, EB).
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

	// для KB и выше используем один знак после запятой
	return fmt.Sprintf("%.1f%s", s, units[i])
}
