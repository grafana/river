package scanner_test

import (
	"testing"

	"github.com/grafana/river/scanner"
	"github.com/stretchr/testify/require"
)

func TestIsValidIdentifier(t *testing.T) {
	tt := []struct {
		name       string
		identifier string
		expect     bool
	}{
		{"empty", "", false},
		{"start_number", "0identifier_1", false},
		{"start_char", "identifier_1", true},
		{"start_underscore", "_identifier_1", true},
		{"special_chars", "!@#$%^&*()", false},
		{"special_char", "identifier_1!", false},
		{"spaces", "identifier _ 1", false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expect, scanner.IsValidIdentifier(tc.identifier))
		})
	}
}

func TestSanitizeIdentifierOptions(t *testing.T) {
	tt := []struct {
		name             string
		identifier       string
		expectIdentifier string
		expectErr        string
		opts             *scanner.SanitizeIdentifierOptions
	}{
		{"empty", "", "empty", "", nil},
		{"start_number", "0identifier_1", "id_0identifier_1", "", nil},
		{"start_char", "identifier_1", "identifier_1", "", nil},
		{"start_underscore", "_identifier_1", "_identifier_1", "", nil},
		{"special_chars", "!@#$%^&*()", "id___________", "", nil},
		{"special_char", "identifier_1!", "identifier_1_", "", nil},
		{"spaces", "identifier _ 1", "identifier___1", "", nil},
		{"bad emptyValue", "", "", "emptyValue `\"\"` is not a valid river identifier", &scanner.SanitizeIdentifierOptions{EmptyValue: "", Prefix: "", Replacement: ""}},
		{"bad prefix", "", "", "prefix `\"123\"` is not a valid river identifier", &scanner.SanitizeIdentifierOptions{EmptyValue: "empty2", Prefix: "123", Replacement: ""}},
		{"bad replacement", "", "", "replacement `\"!\"` must be either a valid river identifier or empty", &scanner.SanitizeIdentifierOptions{EmptyValue: "empty2", Prefix: "prefix2_", Replacement: "!"}},
		{"different empty", "", "empty2", "", &scanner.SanitizeIdentifierOptions{EmptyValue: "empty2", Prefix: "prefix2_", Replacement: ""}},
		{"different prefix", "0identifier_1", "prefix2_0identifier_1", "", &scanner.SanitizeIdentifierOptions{EmptyValue: "empty2", Prefix: "prefix2_", Replacement: ""}},
		{"different replacement", "identifier _ 1", "identifiera_a1", "", &scanner.SanitizeIdentifierOptions{EmptyValue: "empty2", Prefix: "prefix2_", Replacement: "a"}},
		{"empty replacement", "identifier _ 1", "identifier_1", "", &scanner.SanitizeIdentifierOptions{EmptyValue: "empty2", Prefix: "prefix2_", Replacement: ""}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newIdentifier, err := scanner.SanitizeIdentifier(tc.identifier, tc.opts)
			if tc.expectErr != "" {
				require.EqualError(t, err, tc.expectErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectIdentifier, newIdentifier)
		})
	}
}
