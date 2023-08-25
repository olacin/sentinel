package sentinel

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleCrossOriginOpenerPolicy(t *testing.T) {
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
			"DirectiveUnsafeNone",
			"unsafe-none",
			[]Option{
				WithCrossOriginOpenerPolicy(DirectiveOpenerUnsafeNone),
			},
		},
		{
			"DirectiveSameOriginAllowPopups",
			"same-origin-allow-popups",
			[]Option{
				WithCrossOriginOpenerPolicy(DirectiveOpenerSameOriginAllowPopups),
			},
		},
		{
			"DirectiveSameOrigin",
			"same-origin",
			[]Option{
				WithCrossOriginOpenerPolicy(DirectiveOpenerSameOrigin),
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

				assert.Equal(t, tc.expected, res.Header().Get(HeaderCrossOriginOpenerPolicy))
			})
		})
	}
}
