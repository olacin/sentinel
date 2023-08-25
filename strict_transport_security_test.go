package sentinel

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandleStrictTransportSecurity(t *testing.T) {
	cases := []struct {
		name     string
		expected string
		options  []Option
	}{
		{
			"NoConfig",
			"",
			[]Option{},
		},
		{
			"BasicConfig",
			"max-age=3600",
			[]Option{
				WithStrictTransportSecurity(time.Hour, false, false),
			},
		},
		{
			"IncludeSubdomains",
			"max-age=3600; includeSubdomains",
			[]Option{
				WithStrictTransportSecurity(time.Hour, true, false),
			},
		},
		{
			"ZeroMaxAge",
			"",
			[]Option{
				WithStrictTransportSecurity(time.Duration(0), false, false),
			},
		},
		{
			"Preload",
			"max-age=31536000; includeSubdomains; preload",
			[]Option{
				WithStrictTransportSecurity(time.Hour*24*365, true, true),
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

				assert.Equal(t, tc.expected, res.Header().Get(HeaderStrictTransportSecurity))
			})
		})
	}
}
