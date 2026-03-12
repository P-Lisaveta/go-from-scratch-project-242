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

    // Находим подходящую единицу измерения
    units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
    sizeFloat := float64(size)
    unitIndex := 0

    for sizeFloat >= unit && unitIndex < len(units)-1 {
        sizeFloat /= unit
        unitIndex++
    }

    // Форматируем с одной десятичной цифрой
    return fmt.Sprintf("%.1f%s", sizeFloat, units[unitIndex])
}

// GetSize вычисляет размер файла или директории
func GetSize(path string, recursive, includeHidden bool) (int64, error) {
    info, err := os.Stat(path)
    if err != nil {
        return 0, err
    }

    // Если это файл - возвращаем его размер
    if !info.IsDir() {
        return info.Size(), nil
    }

    // Если это директория - читаем содержимое
    entries, err := os.ReadDir(path)
    if err != nil {
        return 0, err
    }

    var totalSize int64
    for _, entry := range entries {
        // Пропускаем скрытые файлы, если не указан флаг --all
        if !includeHidden && strings.HasPrefix(entry.Name(), ".") {
            continue
        }

        fullPath := filepath.Join(path, entry.Name())
        
        if entry.IsDir() {
            if recursive {
                // Рекурсивно обрабатываем поддиректорию
                size, err := GetSize(fullPath, recursive, includeHidden)
                if err != nil {
                    // Продолжаем подсчёт даже если есть ошибки с отдельными файлами
                    continue
                }
                totalSize += size
            }
            continue
        }

        // Получаем размер файла
        fileInfo, err := entry.Info()
        if err != nil {
            continue
        }
        totalSize += fileInfo.Size()
    }

    return totalSize, nil
}