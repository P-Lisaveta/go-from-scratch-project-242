package code

import (
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "strings"
)

// FormatSize конвертирует размер в человекочитаемый формат
func FormatSize(size int64, human bool) string {
    if !human {
        return strconv.FormatInt(size, 10) + "B"
    }

    const unit = 1024
    if size < unit {
        return fmt.Sprintf("%dB", size)
    }

    units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
    sizeFloat := float64(size)
    unitIndex := 0

    for sizeFloat >= unit && unitIndex < len(units)-1 {
        sizeFloat /= unit
        unitIndex++
    }

    return fmt.Sprintf("%.1f%s", sizeFloat, units[unitIndex])
}

// GetPathSize вычисляет размер файла или директории
func GetPathSize(path string, recursive, includeHidden bool) (int64, error) {
    info, err := os.Stat(path)
    if err != nil {
        return 0, err
    }

    if !info.IsDir() {
        return info.Size(), nil
    }

    entries, err := os.ReadDir(path)
    if err != nil {
        return 0, err
    }

    var totalSize int64
    for _, entry := range entries {
        if !includeHidden && strings.HasPrefix(entry.Name(), ".") {
            continue
        }

        fullPath := filepath.Join(path, entry.Name())
        
        if entry.IsDir() {
            if recursive {
                size, err := GetPathSize(fullPath, recursive, includeHidden)
                if err != nil {
                    continue
                }
                totalSize += size
            }
            continue
        }

        fileInfo, err := entry.Info()
        if err != nil {
            continue
        }
        totalSize += fileInfo.Size()
    }

    return totalSize, nil
}