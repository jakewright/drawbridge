package proxy

import (
    "drawbridge/domain"
    "drawbridge/utils"
    "net/http"
    "net/http/httputil"
)

type Proxy struct {
    reverseProxy *httputil.ReverseProxy
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    p.reverseProxy.ServeHTTP(w, r)
}

func New(target *domain.Url) *Proxy {
    reverseProxy := httputil.ReverseProxy{
        Director: func(req *http.Request) {
            req.URL.Scheme = target.Scheme
            req.URL.Host = target.Host
            req.URL.Path = utils.SingleJoiningSlash(target.Path, req.URL.Path)
            req.Host = target.Host
        },
    }

    return &Proxy{reverseProxy: &reverseProxy}
}
