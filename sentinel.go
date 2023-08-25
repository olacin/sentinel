package sentinel

import (
	"fmt"
	"net/http"
	"time"
)

const durationOneYear = time.Hour * 24 * 365

type HeaderMiddleware interface {
	fmt.Stringer
	Apply(w http.ResponseWriter, r *http.Request)
}

type Sentinel struct {
	middlewares []HeaderMiddleware
}

// New creates a new Sentinel handler with the provided options.
func New(options ...Option) *Sentinel {
	s := &Sentinel{}
	for _, opt := range options {
		opt(s)
	}
	return s
}

// Handler creates a new Sentinel handler with passed options.
func Handler(options ...Option) func(next http.Handler) http.Handler {
	s := New(options...)
	return s.Handler
}

// DefaultHandler creates a new Sentinel handler with the default options.
func DefaultHandler() func(next http.Handler) http.Handler {
	s := New(
		WithCrossOriginOpenerPolicy(DirectiveOpenerSameOrigin),
		WithCrossOriginResourcePolicy(DirectiveResourceSameOrigin),
		WithReferrerPolicy(DirectiveNoReferrer),
		WithStrictTransportSecurity(durationOneYear, true, false),
		WithXContentTypeOptions(),
		WithXDNSPrefetchControl(DirectiveXDNSPrefetchControlOff),
		WithXFrameOptions(XFrameOptionsDirectiveSameOrigin),
	)
	return s.Handler
}

// Handler applies the security configuration on the request, and add relevant headers as necessary.
func (s *Sentinel) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, mw := range s.middlewares {
			mw.Apply(w, r)
		}
	})
}
