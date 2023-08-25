package sentinel

import "net/http"

const HeaderXContentTypeOptions = "X-Content-Type-Options"

const DirectiveNoSniff = "nosniff"

type xContentTypeOptions string

func (cto xContentTypeOptions) String() string {
	return string(cto)
}

func (cto xContentTypeOptions) Apply(w http.ResponseWriter, r *http.Request) {
	if cto.String() != "" {
		w.Header().Set(HeaderXContentTypeOptions, cto.String())
	}
}
