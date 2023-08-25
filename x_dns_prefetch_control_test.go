package sentinel

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleXDNSPrefetchControl(t *testing.T) {
	cases := []struct {
		name     string
		expected string
		options  []Option
	}{
		{
			"NoDirective",
			"",
			[]Option{},
		},
		{
			"DirectiveOn",
			"on",
			[]Option{
				WithXDNSPrefetchControl(DirectiveXDNSPrefetchControlOn),
			},
		},
		{
			"DirectiveOff",
			"off",
			[]Option{
				WithXDNSPrefetchControl(DirectiveXDNSPrefetchControlOff),
			},
		},
	}

	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			s := New(tc.options...)

			req, _ := http.NewRequest(http.MethodGet, "https://example.com/foo", nil)

			t.Run("Handler", func(t *testing.T) {
				res := httptest.NewRecorder()
				s.Handler(testHandler).ServeHTTP(res, req)

				assert.Equal(t, tc.expected, res.Header().Get(HeaderXDNSPrefetchControl))
			})
		})
	}
}
