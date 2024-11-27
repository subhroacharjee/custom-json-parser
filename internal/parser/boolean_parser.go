package parser

import "fmt"

type BooleanParser struct{}

func (BooleanParser) IsValid(target string) bool {
	for _, expectedVal := range []string{"true", "false"} {
		if expectedVal == target {
			return true
		}
	}

	return false
}

func (BooleanParser) Parse(target string) (bool, error) {
	if target == "true" {
		return true, nil
	} else if target == "false" {
		return false, nil
	}
	return false, fmt.Errorf("%s is not boolean parsable", target)
}
