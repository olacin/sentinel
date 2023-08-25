package sentinel

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleReferrerPolicy(t *testing.T) {
	cases := []struct {
		name     string
		expected string
		options  []Option
	}{
		{
			"NoPolicy",
			"",
			[]Option{},
		},
		{
			"BasicPolicy",
			"no-referrer",
			[]Option{
				WithReferrerPolicy(DirectiveNoReferrer),
			},
		},
		{
			"FallbackPolicy",
			"no-referrer, strict-origin-when-cross-origin",
			[]Option{
				WithReferrerPolicy(DirectiveNoReferrer, DirectiveStrictOriginWhenCrossOrigin),
			},
		},
	}

	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			s := Handler(tc.options...)

			req, _ := http.NewRequest(http.MethodGet, "https://example.com/foo", nil)

			t.Run("Handler", func(t *testing.T) {
				res := httptest.NewRecorder()
				s(testHandler).ServeHTTP(res, req)

				assert.Equal(t, tc.expected, res.Header().Get(HeaderReferrerPolicy))
			})
		})
	}
}
