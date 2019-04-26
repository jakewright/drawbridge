package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jakewright/drawbridge/config"
	"github.com/jakewright/drawbridge/middleware"
	"github.com/jakewright/drawbridge/proxy"
	"github.com/jakewright/drawbridge/utils"
	"github.com/jakewright/muxinator"
)

// The function that will be called when the program is run
func main() {
	router := muxinator.NewRouter()

	// Create middleware
	logger := middleware.Log(log.New(os.Stdout, "", log.Lshortfile))
	router.AddMiddleware(logger)

	configuration := config.LoadConfig()

	// Loop over all APIs
	for _, apiDefinition := range configuration.Apis {
		log.Printf("%v", apiDefinition)

		// Add surrounding slashes to the prefix
		prefix := utils.AddSlashes(apiDefinition.Prefix)

		// Create a new proxy
		proxy := proxy.New(&apiDefinition)

		// Strip the prefix before passing to the proxy. Without this, the proxy will make a
		// request to upstreamUrl/prefix/path instead of upstreamUrl/path.
		handler := http.StripPrefix(prefix, proxy)

		// Handle /prefix/* with the proxy. The empty string in the first argument means handle all methods.
		router.Handle("", prefix+"*", handler)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Fatal(router.ListenAndServe(":" + port))
}
