package utils

import (
    "strings"
)

func SingleJoiningSlash(a, b string) string {
    aslash := strings.HasSuffix(a, "/")
    bslash := strings.HasPrefix(b, "/")
    switch {
    case aslash && bslash:
        return a + b[1:]
    case !aslash && !bslash:
        return a + "/" + b
    }
    return a + b
}

func AddSlashes(s string) string {
    preSlash := strings.HasPrefix(s, "/")
    postSlash := strings.HasSuffix(s, "/")
    switch {
    case preSlash && postSlash:
        return s
    case !preSlash && postSlash:
        return "/" + s
    case preSlash && !postSlash:
        return s + "/"
    }
    return "/" + s + "/"
}
