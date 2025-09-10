package http

type ResponseArgs[B any] struct {
	Method string
	Path   string
	Body   B
}

type ResponseWant[R any] struct {
	StatusCode  int
	ContentType string
	Location    string
	Response    R
}
