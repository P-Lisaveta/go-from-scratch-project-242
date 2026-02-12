package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathSizeFile(t *testing.T) {

	opts := Options{}
	filePath := filepath.Join("testdata", "test1")
	result, err := GetPathSize(filePath, opts)
	if err != nil {
		t.Fatalf("GetPathSize error: %v", err)
	}
	require.Equal(t, int64(1), result)
}

func TestPathSizeDir(t *testing.T) {
	opts := Options{}
	dirPath := filepath.Join("testdata", "testdirectory")
	result, err := GetPathSize(dirPath, opts)
	if err != nil {
		t.Fatalf("GetPathSize error: %v", err)
	}
	require.Equal(t, int64(5), result)
}

func TestHiddenFilesFilteredByDefault(t *testing.T) {
	dirPath := filepath.Join("testdata", "testdirectory")
	opts := Options{}

	sizeNoHidden, err := GetPathSize(dirPath, opts)
	require.NoError(t, err)
	// считаем только видимые файлы (test3 и test4)
	require.Equal(t, int64(5), sizeNoHidden)

	sizeWithHidden, err := GetPathSize(dirPath, Options{All: true})
	require.NoError(t, err)
	// считаем и видимые, и скрытые (добавляется .test3)
	require.Equal(t, int64(8), sizeWithHidden)
}

func TestFormatSizeRaw(t *testing.T) {
	filePath := filepath.Join("testdata", "test1")
	opts := Options{}
	sizeFile, err := GetPathSize(filePath, opts)
	require.NoError(t, err)
	require.Equal(t, "1B", FormatSize(sizeFile, Options{Human: false}))
	require.Equal(t, "1B", FormatSize(sizeFile, Options{Human: true}))

	dirPath := filepath.Join("testdata", "testdirectory")
	sizeDir, err := GetPathSize(dirPath, Options{All: true})
	require.NoError(t, err)
	require.Equal(t, "8B", FormatSize(sizeDir, Options{}))
	require.Equal(t, "1.2MB", FormatSize(1234567, Options{Human: true}))
}
