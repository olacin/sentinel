package sentinel

import "net/http"

const HeaderCrossOriginResourcePolicy = "Cross-Origin-Resource-Policy"

const (
	DirectiveResourceSameSite    CrossOriginResourcePolicyDirective = "same-site"
	DirectiveResourceSameOrigin  CrossOriginResourcePolicyDirective = "same-origin"
	DirectiveResourceCrossOrigin CrossOriginResourcePolicyDirective = "cross-origin"
)

type (
	CrossOriginResourcePolicyDirective string

	crossOriginResourcePolicy string
)

func (cp crossOriginResourcePolicy) String() string {
	return string(cp)
}

// Apply applies the Cross-Origin-Embedder-Policy configuration to the response.
func (cp crossOriginResourcePolicy) Apply(w http.ResponseWriter, r *http.Request) {
	if cp.String() != "" {
		w.Header().Set(HeaderCrossOriginResourcePolicy, cp.String())
	}
}
