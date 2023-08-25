package sentinel

import (
	"net/http"
	"strings"
)

const HeaderReferrerPolicy = "Referrer-Policy"

const (
	DirectiveNoReferrer                  ReferrerPolicyDirective = "no-referrer"
	DirectiveNoReferrerWhenDowngrade     ReferrerPolicyDirective = "no-referrer-when-downgrade"
	DirectiveOrigin                      ReferrerPolicyDirective = "origin"
	DirectiveOriginWhenCrossOrigin       ReferrerPolicyDirective = "origin-when-cross-origin"
	DirectiveSameOrigin                  ReferrerPolicyDirective = "same-origin"
	DirectiveStrictOrigin                ReferrerPolicyDirective = "strict-origin"
	DirectiveStrictOriginWhenCrossOrigin ReferrerPolicyDirective = "strict-origin-when-cross-origin"
	DirectiveUnsafeURL                   ReferrerPolicyDirective = "unsafe-url"
)

type (
	ReferrerPolicyDirective string

	referrerPolicy struct {
		directives []ReferrerPolicyDirective
	}
)

func newReferrerPolicy(directives ...ReferrerPolicyDirective) *referrerPolicy {
	return &referrerPolicy{
		directives: directives,
	}
}

func (r referrerPolicy) String() string {
	directives := []string{}
	for _, directive := range r.directives {
		directives = append(directives, string(directive))
	}
	return strings.Join(directives, ", ")
}

// Apply applies the Referrer-Policy configuration to the response.
func (rp referrerPolicy) Apply(w http.ResponseWriter, r *http.Request) {
	if len(rp.directives) > 0 {
		w.Header().Set(HeaderReferrerPolicy, rp.String())
	}
}
