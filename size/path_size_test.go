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
	require.Equal(t, int64(1), result)
}

func TestPathSizeDir(t *testing.T) {
	dirPath := filepath.Join("..", "testdata", "testdirectory")
	result, err := GetPathSize(dirPath)
	if err != nil {
		t.Fatalf("GetPathSize error: %v", err)
	}
	require.Equal(t, int64(7), result)
}

func TestFormatSizeRaw(t *testing.T) {
	tests := []struct {
		name   string
		size   int64
		human  bool
		expect string
	}{
		{name: "raw bytes", size: 123, human: false, expect: "123B"},
		{name: "raw zero", size: 0, human: false, expect: "0B"},
		{name: "human small stays bytes", size: 123, human: true, expect: "123B"},
		{name: "human 1KB", size: 1024, human: true, expect: "1.0KB"},
		{name: "human 1.2MB example", size: 1234567, human: true, expect: "1.2MB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FormatSize(tt.size, tt.human)
			require.Equal(t, tt.expect, actual)
		})
	}
}
