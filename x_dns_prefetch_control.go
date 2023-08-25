package sentinel

import "net/http"

const HeaderXDNSPrefetchControl = "X-DNS-Prefetch-Control"

const (
	DirectiveXDNSPrefetchControlOn  XDNSPrefetchControlDirective = "on"
	DirectiveXDNSPrefetchControlOff XDNSPrefetchControlDirective = "off"
)

type (
	XDNSPrefetchControlDirective string

	xDNSPrefetchControl string
)

func (xdpc xDNSPrefetchControl) String() string {
	return string(xdpc)
}

// Apply applies the X-DNS-Prefetch-Control configuration to the response.
func (xdpc xDNSPrefetchControl) Apply(w http.ResponseWriter, r *http.Request) {
	if xdpc.String() != "" {
		w.Header().Set(HeaderXDNSPrefetchControl, xdpc.String())
	}
}
