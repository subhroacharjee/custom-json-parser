package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subhroacharjee/custom-json-parser/internal/parser"
)

func TestNullParser(t *testing.T) {
	tests := []struct {
		name          string
		target        string
		expectedValue any
		expectErr     bool
		validate      bool
	}{
		{
			name:          "Should return true when 'null' as a string is passed to isValid method",
			target:        "null",
			expectedValue: true,
			expectErr:     false,
			validate:      true,
		},
		{
			name:          "Should return false when random string is passed to isValid method",
			target:        "random string",
			expectedValue: false,
			expectErr:     false,
			validate:      true,
		},
		{
			name:          "Should return nil and error as nil when 'null' as a string is passed to parse method",
			target:        "null",
			expectedValue: nil,
			expectErr:     false,
			validate:      false,
		},
		{
			name:          "Should return nil and error when random string is passed to parse method",
			target:        "random string",
			expectedValue: nil,
			expectErr:     true,
			validate:      false,
		},
	}

	nullParser := parser.NullParser{}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			if tt.validate {
				assert.Equal(t, tt.expectedValue, nullParser.IsValid(tt.target), "should match expected value of validate")
			} else {
				res, err := nullParser.Parse(tt.target)
				assert.Equal(t, tt.expectedValue, res, "parser value should match the expectedValue")
				if tt.expectErr {
					assert.NotNil(t, err, "parser is expected to give some error")
				}
			}
		})
	}
}
