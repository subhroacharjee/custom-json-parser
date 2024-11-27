package parser

import "fmt"

type NullParser struct{}

func (NullParser) IsValid(target string) bool {
	return target == "null"
}

func (NullParser) Parse(target string) (any, error) {
	if target == "null" {
		return nil, nil
	}
	return nil, fmt.Errorf("%s is not null parseable", target)
}
