package scanner

import (
	"fmt"

	"github.com/grafana/river/token"
)

// IsValidIdentifier returns true if the given string is a valid river
// identifier.
func IsValidIdentifier(in string) bool {
	s := New(nil, []byte(in), nil, 0)
	_, tok, lit := s.Scan()
	return tok == token.IDENT && lit == in
}

type SanitizeIdentifierOptions struct {
	// Prefix is what will be prepended to the identifier if it does not start
	// with a letter or underscore. This must be a valid river identifier.
	Prefix string

	// Replacement is what will be used to replace invalid characters. This
	// must be a valid river identifier or empty.
	Replacement string
}

func sanitizeIdentifierOptionsDefault() *SanitizeIdentifierOptions {
	return &SanitizeIdentifierOptions{
		Prefix:      "_",
		Replacement: "_",
	}
}

// validate will return an error if the options are invalid.
func (opts *SanitizeIdentifierOptions) validate() error {
	if !IsValidIdentifier(opts.Prefix) {
		return fmt.Errorf("prefix `%q` is not a valid river identifier", opts.Prefix)
	}

	if !(IsValidIdentifier(opts.Replacement) || opts.Replacement == "") {
		return fmt.Errorf("replacement `%q` must be either a valid river identifier or empty", opts.Replacement)
	}

	return nil
}

// SanitizeIdentifier will return the given string mutated into a valid river
// identifier. If the given string is already a valid identifier, it will be
// returned unchanged.
//
// This should be used with caution since the different inputs can result in
// identical outputs.
func SanitizeIdentifier(in string, opts *SanitizeIdentifierOptions) (string, error) {
	if IsValidIdentifier(in) {
		return in, nil
	}

	if opts == nil {
		opts = sanitizeIdentifierOptionsDefault()
	}

	if err := opts.validate(); err != nil {
		return "", err
	}

	return generateNewIdentifier(in, opts.Prefix, opts.Replacement)
}

// generateNewIdentifier expects a valid river prefix and replacement
// string and returns a new identifier based on the given input.
func generateNewIdentifier(in string, prefix string, replacement string) (string, error) {
	if in == "" {
		return "", fmt.Errorf("cannot generate a new identifier for an empty string")
	}

	newValue := ""
	for i, c := range in {
		if i == 0 {
			if !isLetter(c) {
				newValue = prefix
			}
		}

		if !(isLetter(c) || isDigit(c)) {
			newValue += replacement
			continue
		}

		newValue += string(c)
	}

	return newValue, nil
}
