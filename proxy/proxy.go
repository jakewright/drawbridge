package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/jakewright/drawbridge/utils"
)

// Proxy is a handler that takes an incoming request and sends it
// to the upstream API, proxying the response back to the client.
type Proxy struct {
	allowCrossOrigin bool
	reverseProxy     *httputil.ReverseProxy
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p.allowCrossOrigin {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	}

	p.reverseProxy.ServeHTTP(w, r)
}

// New returns a Proxy to the given target
func New(target *url.URL, allowCrossOrigin bool) *Proxy {
	reverseProxy := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = utils.SingleJoiningSlash(target.Path, req.URL.Path)
			req.Host = target.Host
		},
	}

	return &Proxy{allowCrossOrigin: allowCrossOrigin, reverseProxy: &reverseProxy}
}
