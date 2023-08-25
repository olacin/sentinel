package sentinel

import (
	"fmt"
	"net/http"
	"strings"
)

const HeaderXXSSProtection = "X-XSS-Protection"

type (
	xXSSProtection struct {
		reportURI string
		enabled   bool
		block     bool
	}
)

func newXXSSProtection(enabled, block bool, reportURI string) *xXSSProtection {
	return &xXSSProtection{
		reportURI, enabled, block,
	}
}

func (xp xXSSProtection) String() string {
	directives := []string{}

	if xp.enabled {
		directives = append(directives, "1")
	} else {
		directives = append(directives, "0")
	}

	if xp.block && xp.reportURI == "" {
		directives = append(directives, "mode=block")
	} else if xp.reportURI != "" {
		directives = append(directives, fmt.Sprintf("report=%s", xp.reportURI))
	}

	return strings.Join(directives, "; ")
}

func (xp xXSSProtection) Apply(w http.ResponseWriter, r *http.Request) {
	if xp.String() != "" {
		w.Header().Set(HeaderXXSSProtection, xp.String())
	}
}
