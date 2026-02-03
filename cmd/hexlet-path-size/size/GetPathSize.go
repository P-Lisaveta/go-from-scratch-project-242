package PathSize

import (
	"fmt"
	"os"
)

func GetPathSize(path string, _, _, _ bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	var size int64
	if !info.IsDir() {
		size = info.Size()
	} else {
		entries, _ := os.ReadDir(path)
		for _, entry := range entries {
			if !entry.IsDir() {
				fInfo, _ := entry.Info()
				size += fInfo.Size()
			}
		}
	}
	return fmt.Sprintf("%d\t%s", size, path), nil
}
