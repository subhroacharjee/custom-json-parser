package parser

type Parser[T any] interface {
	IsValid(string) bool
	Parse(string) (T, error)
}
