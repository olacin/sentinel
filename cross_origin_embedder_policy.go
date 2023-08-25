package sentinel

import "net/http"

const HeaderCrossOriginEmbedderPolicy = "Cross-Origin-Embedder-Policy"

const (
	DirectiveEmbedderUnsafeNone     CrossOriginEmbedderPolicyDirective = "unsafe-none"
	DirectiveEmbedderRequireCorp    CrossOriginEmbedderPolicyDirective = "require-corp"
	DirectiveEmbedderCredentialless CrossOriginEmbedderPolicyDirective = "credentialless"
)

type (
	CrossOriginEmbedderPolicyDirective string

	crossOriginEmbedderPolicy string
)

func (cp crossOriginEmbedderPolicy) String() string {
	return string(cp)
}

// Apply applies the Cross-Origin-Embedder-Policy configuration to the response.
func (cp crossOriginEmbedderPolicy) Apply(w http.ResponseWriter, r *http.Request) {
	if cp.String() != "" {
		w.Header().Set(HeaderCrossOriginEmbedderPolicy, cp.String())
	}
}
