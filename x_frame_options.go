package sentinel

import "net/http"

const HeaderXFrameOptions = "X-Frame-Options"

const (
	XFrameOptionsDirectiveDeny       XFrameOptionsDirective = "DENY"
	XFrameOptionsDirectiveSameOrigin XFrameOptionsDirective = "SAMEORIGIN"
)

type (
	XFrameOptionsDirective string

	xFrameOptions string
)

func (xfo xFrameOptions) String() string {
	return string(xfo)
}

func (xfo xFrameOptions) Apply(w http.ResponseWriter, r *http.Request) {
	if xfo.String() != "" {
		w.Header().Set(HeaderXFrameOptions, xfo.String())
	}
}
