package size

import (
	"fmt"
	"os"
)

func GetPathSize(path string) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	var size int64
	if !info.IsDir() {
		size = info.Size()
	} else {
		entries, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			fInfo, err := entry.Info()
			if err != nil {
				return "", err
			}
			size += fInfo.Size()
		}
	}
	return fmt.Sprintf("%d\t%s", size, path), nil
}
