package middleware

import (
	"net/http"
	"slices"
)

type Middleware func(http.Handler) http.Handler

type Chain struct {
	middlewares []Middleware
}

func New(mw ...Middleware) Chain {
	return Chain{
		middlewares: slices.Clone(mw),
	}
}

func (c Chain) Append(mw ...Middleware) Chain {
	return Chain{
		middlewares: append(slices.Clone(c.middlewares), mw...),
	}
}

func (c Chain) Then(h http.Handler) http.Handler {
	for _, mw := range slices.Backward(c.middlewares) {
		h = mw(h)
	}
	return h
}

func (c Chain) ThenFunc(h http.HandlerFunc) http.Handler {
	return c.Then(h)
}
