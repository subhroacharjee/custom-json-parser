package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subhroacharjee/custom-json-parser/internal/parser"
)

func TestBooleanParser(t *testing.T) {
	tests := []struct {
		name          string
		target        string
		expectedValue bool
		expectErr     bool
		validate      bool
	}{
		{
			name:          "Should return true when 'true' as a string is passed to isValid method",
			target:        "true",
			expectedValue: true,
			expectErr:     false,
			validate:      true,
		},
		{
			name:          "Should return true when 'false' as a string is passed to isValid method",
			target:        "true",
			expectedValue: true,
			expectErr:     false,
			validate:      true,
		},
		{
			name:          "Should return false when random string is passed to isValid method",
			target:        "not a valid boolean string",
			expectedValue: false,
			expectErr:     false,
			validate:      true,
		},
		{
			name:          "Should return true and error as nil when 'true' as a string is passed to parse method",
			target:        "true",
			expectedValue: true,
			expectErr:     false,
			validate:      false,
		},
		{
			name:          "Should return false and error as nil when 'false' as a string is passed to parse method",
			target:        "false",
			expectedValue: false,
			expectErr:     false,
			validate:      false,
		},
		{
			name:          "Should return false and error when random string is passed to isValid method",
			target:        "not a valid boolean string",
			expectedValue: false,
			expectErr:     true,
			validate:      false,
		},
	}

	boolParser := parser.BooleanParser{}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			if tt.validate {
				assert.Equal(t, tt.expectedValue, boolParser.IsValid(tt.target), "should match expected value of validate")
			} else {
				res, err := boolParser.Parse(tt.target)
				assert.Equal(t, tt.expectedValue, res, "parser value should match the expectedValue")
				if tt.expectErr {
					assert.NotNil(t, err, "parser is expected to give some error")
				}
			}
		})
	}
}
