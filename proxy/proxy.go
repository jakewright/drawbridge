package proxy

import (
    "drawbridge/domain"
    "drawbridge/utils"
    "net/http"
    "net/http/httputil"
)

type Proxy struct {
    api *domain.Api
    reverseProxy *httputil.ReverseProxy
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if p.api.AllowCrossOrigin {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH")
        w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
    }

    p.reverseProxy.ServeHTTP(w, r)
}

func New(api *domain.Api) *Proxy {
    target := api.UpstreamUrl
    reverseProxy := httputil.ReverseProxy{
        Director: func(req *http.Request) {
            req.URL.Scheme = target.Scheme
            req.URL.Host = target.Host
            req.URL.Path = utils.SingleJoiningSlash(target.Path, req.URL.Path)
            req.Host = target.Host
        },
    }

    return &Proxy{api: api, reverseProxy: &reverseProxy}
}
