package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathSizeFile(t *testing.T) {
	t.Helper()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "file")
	err := os.WriteFile(filePath, []byte("x"), 0o644)
	require.NoError(t, err)

	result, err := GetPathSize(filePath, false, false)
	require.NoError(t, err)
	require.Equal(t, int64(1), result)
}

func TestPathSizeDir(t *testing.T) {
	t.Helper()

	tmpDir := t.TempDir()
	// Create 5 bytes total in directory
	err := os.WriteFile(filepath.Join(tmpDir, "a"), []byte("xx"), 0o644)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tmpDir, "b"), []byte("xxx"), 0o644)
	require.NoError(t, err)

	result, err := GetPathSize(tmpDir, false, false)
	require.NoError(t, err)
	require.Equal(t, int64(5), result)
}

func TestHiddenFilesFilteredByDefault(t *testing.T) {
	t.Helper()

	tmpDir := t.TempDir()
	// Visible file: 5 bytes
	err := os.WriteFile(filepath.Join(tmpDir, "visible"), []byte("xxxxx"), 0o644)
	require.NoError(t, err)
	// Hidden file: 3 bytes
	err = os.WriteFile(filepath.Join(tmpDir, ".hidden"), []byte("yyy"), 0o644)
	require.NoError(t, err)

	sizeNoHidden, err := GetPathSize(tmpDir, false, false)
	require.NoError(t, err)
	require.Equal(t, int64(5), sizeNoHidden)

	sizeWithHidden, err := GetPathSize(tmpDir, false, true)
	require.NoError(t, err)
	require.Equal(t, int64(8), sizeWithHidden)
}

func TestFormatSize(t *testing.T) {
	t.Helper()

	// Raw mode
	require.Equal(t, "1B", FormatSize(1, false))

	// Human-readable exact bytes < 1024
	require.Equal(t, "1B", FormatSize(1, true))

	// Human-readable for larger number
	require.Equal(t, "1.2MB", FormatSize(1234567, true))
}

