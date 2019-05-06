package proxy

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestProxy_ServeHTTP(t *testing.T) {
	// Create a server that returns a valid response
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.URL.String(), "/foo")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("this is the body"))
	})
	s := httptest.NewServer(h)
	defer s.Close()

	// Create a recorder that implements the http.ResponseWriter interface
	rr := httptest.NewRecorder()

	// Create a request
	r := newTestRequest(t, "GET", "/foo", "")

	// Create a new proxy
	p := newTestProxy(t, s.URL)

	p.ServeHTTP(rr, r)

	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, rr.Body.String(), "this is the body")
}

func newTestRequest(t *testing.T, method, url, body string) *http.Request {
	b := strings.NewReader(body)
	r, err := http.NewRequest(method, url, b)
	assert.NilError(t, err)
	return r
}

func newTestProxy(t *testing.T, rawURL string) *Proxy {
	// Set the API's upstream URL to that of the server
	u, err := url.Parse(rawURL)
	assert.NilError(t, err)

	return New(u, true)
}
