package size

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathSizeFile(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "test1")
	result, err := GetPathSize(filePath, false)
	if err != nil {
		t.Fatalf("GetPathSize error: %v", err)
	}
	require.Equal(t, int64(1), result)
}

func TestPathSizeDir(t *testing.T) {
	dirPath := filepath.Join("..", "testdata", "testdirectory")
	result, err := GetPathSize(dirPath, false)
	if err != nil {
		t.Fatalf("GetPathSize error: %v", err)
	}
	require.Equal(t, int64(5), result)
}

func TestHiddenFilesFilteredByDefault(t *testing.T) {
	dirPath := filepath.Join("..", "testdata", "testdirectory")

	sizeNoHidden, err := GetPathSize(dirPath, false)
	require.NoError(t, err)
	// считаем только видимые файлы (test3 и test4)
	require.Equal(t, int64(5), sizeNoHidden)

	sizeWithHidden, err := GetPathSize(dirPath, true)
	require.NoError(t, err)
	// считаем и видимые, и скрытые (добавляется .test3)
	require.Equal(t, int64(8), sizeWithHidden)
}

func TestFormatSizeRaw(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "test1")
	sizeFile, err := GetPathSize(filePath, true)
	require.NoError(t, err)
	require.Equal(t, "1B", FormatSize(sizeFile, false))
	require.Equal(t, "1B", FormatSize(sizeFile, true))

	dirPath := filepath.Join("..", "testdata", "testdirectory")
	sizeDir, err := GetPathSize(dirPath, true)
	require.NoError(t, err)
	require.Equal(t, "8B", FormatSize(sizeDir, false))

	// сохраняем пример из ТЗ для human-формата
	require.Equal(t, "1.2MB", FormatSize(1234567, true))
}
