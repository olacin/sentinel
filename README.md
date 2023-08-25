# sentinel

[![go report](https://goreportcard.com/badge/github.com/olacin/sentinel)](https://goreportcard.com/report/github.com/olacin/sentinel)
[![codecov](https://codecov.io/gh/olacin/sentinel/graph/badge.svg?token=MyCssje2s8)](https://codecov.io/gh/olacin/sentinel)
[![go reference](https://pkg.go.dev/badge/github.com/olacin/sentinel.svg)](https://pkg.go.dev/github.com/olacin/sentinel)

net/http security middleware, inspired from Django security middlewares.

It does not aim to be the perfect security middleware, but can help to setup basic security through setting secure HTTP headers.

**This project won't consider frameworks adapters (gin, echo, fiber...) to stay lightweight and not depend on these libraries.**

## Getting started

```go
package main

import (
	"log"
	"net/http"

	"github.com/olacin/sentinel"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("foo"))
	})

	secure := sentinel.DefaultHandler()
	http.Handle("/", secure(handler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

By default, Sentinel sets these headers with the following configuration:

```go
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
```

## Implemented headers

- [x] Cross-Origin-Embedder-Policy
- [x] Cross-Origin-Opener-Policy
- [x] Cross-Origin-Resource-Policy
- [x] Referrer-Policy
- [x] Strict-Transport-Security
- [x] X-Content-Type-Options
- [x] X-DNS-Prefetch-Control
- [x] X-Frame-Options
- [x] X-XSS-Protection
