package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	clirunner "github.com/subhroacharjee/custom-json-parser/internal/cli_runner"
)

func setup(fileName string, content string) string {
	filePath, err := CreateJsonFile(fileName, content)
	if err != nil {
		panic(err)
	}
	return filePath
}

func teardown(fileName string) {
	if err := RemoveJsonFile(fileName); err != nil {
		panic(err)
	}
}

func TestJSONParserValidation(t *testing.T) {
	createTmpDirIfNotExists()

	tests := []struct {
		testName    string
		fileName    string
		content     string
		shouldPanic bool
	}{
		{
			testName:    "Should not panic when the json is valid empty",
			fileName:    "step-1-valid.json",
			content:     "{}",
			shouldPanic: false,
		},
		{
			testName:    "Should panic when the json is not empty",
			fileName:    "step-1-invalid.json",
			content:     "",
			shouldPanic: true,
		},
		{
			testName:    "should panic when the json object has a comma after key and value",
			fileName:    "step-2-invalid-1.json",
			content:     "{\"key\": \"value\",}",
			shouldPanic: true,
		},
		{
			testName:    "should panic when the json object has a key not in quotes",
			fileName:    "step-2-invalid-2.json",
			content:     "{\n \"key\": \"value\",\n key2: \"value\"\n}",
			shouldPanic: true,
		},
		{
			testName:    "should not panic when the json object is valid with one key and value",
			fileName:    "step-2-valid-1.json",
			content:     "{\"key\": \"value\"}",
			shouldPanic: false,
		},
		{
			testName:    "should not panic when the json object is formatted",
			fileName:    "step-2-valid-2.json",
			content:     "{\n \"key\": \"value\",\n \"key2\": \"value\" }",
			shouldPanic: false,
		},
		{
			testName: "should panic when the json object has a invalid value",
			fileName: "step-3-invalid-1.json",
			content: `{
  "key1": true,
  "key2": False,
  "key3": null,
  "key4": "value",
  "key5": 101
}`,
			shouldPanic: true,
		},
		{
			testName: "should not panic when the json object has all valid values",
			fileName: "step-3-valid-1.json",
			content: `{
  "key1": true,
  "key2": false,
  "key3": null,
  "key4": "value",
  "key5": 101
}`,
			shouldPanic: false,
		},
		{
			testName: "should not panic when the json object has all valid json datastructures",
			fileName: "step-4-valid-1.json",
			content: `{
  "key": "value",
  "key-n": 101,
  "key-o": {},
  "key-l": []
}`,
			shouldPanic: false,
		},
		{
			testName: "should not panic when the json object has all valid nested datastructures",
			fileName: "step-4-valid-2.json",
			content: `{
  "key": "value",
  "key-n": 101,
  "key-o": {
    "inner key": "inner value"
  },
  "key-l": ["list value"]
}`,
			shouldPanic: false,
		},
		{
			testName: "should panic when the json object has all invalid nested datastructures",
			fileName: "step-4-invalid-1.json",
			content: `{
  "key": "value",
  "key-n": 101,
  "key-o": {
    "inner key": "inner value"
  },
  "key-l": ['list value']
}`,
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()
			filepath := setup(tt.fileName, tt.content)
			defer teardown(tt.fileName)

			validateCmd := clirunner.ValidateCmd{
				Path: filepath,
			}

			err := validateCmd.Run()

			if tt.shouldPanic {
				assert.NotNil(t, err, "should throw an error when invalid")
			} else {
				assert.Nil(t, err, "should not throw an error when not nil")
			}
		})
	}
}
