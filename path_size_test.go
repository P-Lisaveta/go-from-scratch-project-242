package code

import (
	cli "github.com/urfave/cli/v3"
	"hexlet/code"
	"path/filepath"
	"testing"
)

func TestPathSizeFile(t *testing.T) {

	filePath := filepath.Join("testdata", "test1")
	result, err := GetPathSize(filePath, false, false, false)
	require.NoError(t, err)
	require.Equal(t, int64(1), result)
}

func TestPathSizeDir(t *testing.T) {
	dirPath := filepath.Join("testdata", "testdirectory")
	result, err := GetPathSize(dirPath, false, false, false)
	require.NoError(t, err)
	require.Equal(t, int64(5), result)
}

func TestHiddenFilesFilteredByDefault(t *testing.T) {
	dirPath := filepath.Join("testdata", "testdirectory")

	sizeNoHidden, err := GetPathSize(dirPath, false, false, false)
	require.NoError(t, err)
	require.Equal(t, int64(5), sizeNoHidden)

	sizeWithHidden, err := GetPathSize(dirPath, false, false, true)
	require.NoError(t, err)
	require.Equal(t, int64(8), sizeWithHidden)
}

func TestFormatSizeRaw(t *testing.T) {
	filePath := filepath.Join("testdata", "test1")
	sizeFile, err := GetPathSize(filePath, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "1B", FormatSize(sizeFile, false))
	require.Equal(t, "1B", FormatSize(sizeFile, true))

	dirPath := filepath.Join("testdata", "testdirectory")
	sizeDir, err := GetPathSize(dirPath, false, false, true)
	require.NoError(t, err)
	require.Equal(t, "8B", FormatSize(sizeDir, false))

	require.Equal(t, "1.2MB", FormatSize(1234567, true))
}
