package sentinel

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("bar"))
})

func TestDefaultHandler(t *testing.T) {
	s := DefaultHandler()
	req, _ := http.NewRequest(http.MethodGet, "https://example.com/foo", nil)

	t.Run("Handler", func(t *testing.T) {
		res := httptest.NewRecorder()
		s(testHandler).ServeHTTP(res, req)

		assert.Equal(t, "same-origin", res.Header().Get(HeaderCrossOriginOpenerPolicy))
		assert.Equal(t, "same-origin", res.Header().Get(HeaderCrossOriginResourcePolicy))
		assert.Equal(t, "no-referrer", res.Header().Get(HeaderReferrerPolicy))
		assert.Equal(t, "max-age=31536000; includeSubdomains", res.Header().Get(HeaderStrictTransportSecurity))
		assert.Equal(t, "nosniff", res.Header().Get(HeaderXContentTypeOptions))
		assert.Equal(t, "off", res.Header().Get(HeaderXDNSPrefetchControl))
		assert.Equal(t, "SAMEORIGIN", res.Header().Get(HeaderXFrameOptions))
	})
}
