package sentinel

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleXSSProtection(t *testing.T) {
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
			"DirectiveDisable",
			"0",
			[]Option{
				WithXSSProtection(false, false, ""),
			},
		},
		{
			"DirectiveEnable",
			"1",
			[]Option{
				WithXSSProtection(true, false, ""),
			},
		},
		{
			"DirectiveBlock",
			"1; mode=block",
			[]Option{
				WithXSSProtection(true, true, ""),
			},
		},
		{
			"DirectiveReport",
			"1; report=https://report.example.com/",
			[]Option{
				WithXSSProtection(true, true, "https://report.example.com/"),
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

				assert.Equal(t, tc.expected, res.Header().Get(HeaderXXSSProtection))
			})
		})
	}
}
