package commons

import (
	"os"
	"testing"
)

func TestPathPrepare(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Replace $HOME",
			path:     "$HOME/test",
			expected: os.Getenv("HOME") + "/test",
		},
		{
			name:     "Replace $TMP",
			path:     "$TMP/test",
			expected: os.Getenv("TMP") + "/test",
		},
		{
			name:     "Replace $CWD",
			path:     "$CWD/test",
			expected: os.Getenv("PWD") + "/test",
		},
		{
			name:     "No replacements",
			path:     "/test/path",
			expected: "/test/path",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PathPrepare(test.path)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			} else {
				t.Logf("Passed Test %s, Input %s, Expected %s, Result %s", test.name, test.path, test.expected, result)
			}
		})
	}
}
