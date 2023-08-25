package sentinel

import (
	"fmt"
	"net/http"
	"strings"
)

const HeaderStrictTransportSecurity = "Strict-Transport-Security"

type (
	StrictTransportSecurityDirective string

	strictTransportSecurity struct {
		maxAge            int
		includeSubdomains bool
		preload           bool
	}
)

const (
	DirectiveIncludeSubdomains StrictTransportSecurityDirective = "includeSubdomains"
	DirectivePreload           StrictTransportSecurityDirective = "preload"
)

func newStrictTransportSecurity(maxAge int, includeSubdomains bool, preload bool) *strictTransportSecurity {
	return &strictTransportSecurity{
		maxAge,
		includeSubdomains,
		preload,
	}
}

func (sts strictTransportSecurity) String() string {
	maxAgeDirective := fmt.Sprintf("max-age=%d", sts.maxAge)
	directives := []string{maxAgeDirective}

	if sts.includeSubdomains {
		directives = append(directives, string(DirectiveIncludeSubdomains))
	}
	if sts.preload {
		directives = append(directives, string(DirectivePreload))
	}
	return strings.Join(directives, "; ")
}

func (sts strictTransportSecurity) Apply(w http.ResponseWriter, r *http.Request) {
	if sts.maxAge > 0 {
		w.Header().Set(HeaderStrictTransportSecurity, sts.String())
	}
}
