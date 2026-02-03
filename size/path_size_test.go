package size

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathSizeFile(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "test1")
	result, err := GetPathSize(filePath)
	if err != nil {
		t.Fatalf("GetPathSize error: %v", err)
	}
	require.Equal(t, "1\t../testdata/test1", result)
}

func TestPathSizeDir(t *testing.T) {
	dirPath := filepath.Join("..", "testdata", "testdirectory")
	result, err := GetPathSize(dirPath)
	if err != nil {
		t.Fatalf("GetPathSize error: %v", err)
	}
	require.Equal(t, "7\t../testdata/testdirectory", result)
}
