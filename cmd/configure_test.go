// cmd/configure_test.go
package cmd

import "testing"

// TestRedactKey validates the logic for redacting API keys for display[cite: 4].
func TestRedactKey(t *testing.T) {
	// We use a table-driven test to check multiple cases efficiently.
	testCases := []struct {
		name     string
		inputKey string
		expected string
	}{
		{
			name:     "Key is set",
			inputKey: "sk-12345abcdefg",
			expected: "<set>",
		},
		{
			name:     "Key is not set (empty string)",
			inputKey: "",
			expected: "<not set>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := redactKey(tc.inputKey)
			if result != tc.expected {
				t.Errorf("redactKey(%q) = %q; want %q", tc.inputKey, result, tc.expected)
			}
		})
	}
}