package sentinel

import "net/http"

const HeaderCrossOriginOpenerPolicy = "Cross-Origin-Opener-Policy"

const (
	DirectiveOpenerUnsafeNone            CrossOriginOpenerPolicyDirective = "unsafe-none"
	DirectiveOpenerSameOriginAllowPopups CrossOriginOpenerPolicyDirective = "same-origin-allow-popups"
	DirectiveOpenerSameOrigin            CrossOriginOpenerPolicyDirective = "same-origin"
)

type (
	CrossOriginOpenerPolicyDirective string

	crossOriginOpenerPolicy string
)

func (cp crossOriginOpenerPolicy) String() string {
	return string(cp)
}

// Apply applies the Cross-Origin-Opener-Policy configuration to the response.
func (cp crossOriginOpenerPolicy) Apply(w http.ResponseWriter, r *http.Request) {
	if cp.String() != "" {
		w.Header().Set(HeaderCrossOriginOpenerPolicy, cp.String())
	}
}
