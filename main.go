package main

import (
    "drawbridge/config"
    "drawbridge/proxy"
    "drawbridge/utils"
    "drawbridge/middleware"
    "github.com/gorilla/mux"
    "github.com/urfave/negroni"
    "log"
    "net/http"
    "os"
)


// The function that will be called when the program is run
func main() {
    // Create instances of mux and negroni
    router := mux.NewRouter()
    negroni := negroni.New()

    // Create middleware
    logger := middleware.Log(log.New(os.Stdout, "", log.Lshortfile))

    // Add middleware to the negroni stack
    negroni.UseFunc(logger)

    configuration := config.LoadConfig()

    // Loop over all APIs
    for _, apiDefinition := range configuration.Apis {
        log.Printf("%v", apiDefinition)

        // Add surrounding slashes to the prefix
        prefix := utils.AddSlashes(apiDefinition.Prefix)

        // Create a new proxy
        proxy := proxy.New(apiDefinition.UpstreamUrl)

        // Strip the prefix before passing to the proxy. Without this, the proxy will make a
        // request to upstreamUrl/prefix/path instead of upstreamUrl/path.
        handler := http.StripPrefix(prefix, proxy)

        // Add a handler to the router for the prefix
        router.PathPrefix(prefix).Handler(handler)
    }

    // The router must be the last middleware in the chain
    negroni.UseHandler(router)

    // Start the web server
    log.Fatal(http.ListenAndServe(":80", negroni))
}
