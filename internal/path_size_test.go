package size

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathSizeFile(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "test1")
	result, err := GetPathSize(filePath, false, false)
	if err != nil {
		t.Fatalf("GetPathSize error: %v", err)
	}
	require.Equal(t, int64(1), result)
}

func TestPathSizeDir(t *testing.T) {
	dirPath := filepath.Join("..", "testdata")
	result, err := GetPathSize(dirPath, false, false)
	if err != nil {
		t.Fatalf("GetPathSize error: %v", err)
	}
	require.Equal(t, int64(3), result)
}

func TestHiddenFilesFilteredByDefault(t *testing.T) {
	dirPath := filepath.Join("..", "testdata")

	sizeWithHidden, err := GetPathSize(dirPath, true, false)
	require.NoError(t, err)
	require.Equal(t, int64(6154), sizeWithHidden)
}

func TestFormatSizeRaw(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "test1")
	sizeFile, err := GetPathSize(filePath, true, false)
	require.NoError(t, err)
	require.Equal(t, "1B", FormatSize(sizeFile, false))
	require.Equal(t, "1B", FormatSize(sizeFile, true))

	dirPath := filepath.Join("..", "testdata")
	sizeDir, err := GetPathSize(dirPath, true, false)
	require.NoError(t, err)
	require.Equal(t, "6154B", FormatSize(sizeDir, false))

	require.Equal(t, "1.2MB", FormatSize(1234567, true))
}

func TestHiddenFilesRecursive(t *testing.T) {
	dirPath := filepath.Join("..", "testdata")

	sizeWithRecursive, err := GetPathSize(dirPath, false, true)
	require.NoError(t, err)
	require.Equal(t, int64(6), sizeWithRecursive)
}
