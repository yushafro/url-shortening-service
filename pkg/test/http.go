package test

type Response[R any] struct {
	StatusCode  int
	Response    R
	ContentType string
	WantError   bool
}

type Request[B any] struct {
	Method      string
	Path        string
	ContentType string
	Body        B
}
