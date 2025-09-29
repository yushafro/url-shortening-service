package test

type Test[A, W any] struct {
	Name  string
	Args  A
	Want  W
	Error error
}
