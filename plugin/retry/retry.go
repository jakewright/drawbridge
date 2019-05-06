package retry

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/jakewright/drawbridge/log"
	"github.com/jakewright/drawbridge/plugin"
	"github.com/jakewright/muxinator"
)

func init() {
	plugin.RegisterPlugin("retry", &Retry{})
}

type Retry struct{}

// New returns middleware that retries the next handler according to the given config
func (r *Retry) Middleware(cfg map[string]interface{}) (muxinator.Middleware, error) {
	var opts Options
	if err := plugin.DecodeConfig(cfg, &opts); err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		// Save a copy of the body so we can reset before each retry
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		}
		if err = r.Body.Close(); err != nil {
			log.Printf("Failed to close request body: %v", err)
		}

		// Wrap the ResponseWriter so we can
		// see the status code that is written
		ww := &responseWriter{w: w}
		defer func() {
			if _, err := ww.Flush(); err != nil {
				log.Printf("Failed to write response: %v", err)
			}
		}()

		i := 1
		for {
			// Replace the body before passing onto the next handler
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			// Perform the request and break if everything looks good
			next(ww, r)
			if !shouldRetry(&opts, r, ww.s) {
				break
			}

			if i == opts.Attempts {
				log.Printf("Request failed with status %d; giving up after %d attempts", ww.s, i)
				break
			}

			// Sleep for the backoff period, unless the context is cancelled.
			period := opts.backoffStrategy.Calculate(i)
			log.Printf("Request failed with status %d [%d]; backing off for %s", ww.s, i, period.String())
			select {
			case <-time.After(period):
			case <-r.Context().Done():
				return
			}

			// Reset ww to clear any buffered response
			ww.Reset()
			i++
		}

	}, nil
}

func shouldRetry(opts *Options, r *http.Request, statusCode int) bool {
	if !shouldRetryMethod(r.Method, opts.RetryMethods) {
		return false
	}

	if !shouldRetryStatus(statusCode, opts.expression) {
		return false
	}

	return true
}

func shouldRetryMethod(method string, retryMethods []string) bool {
	for _, m := range retryMethods {
		if method == m {
			return true
		}
	}

	return false
}

func shouldRetryStatus(statusCode int, expression *govaluate.EvaluableExpression) bool {
	result, err := expression.Evaluate(map[string]interface{}{
		"statusCode": statusCode,
	})
	if err != nil {
		log.Printf("Failed to evaluate expression: %v", expression)
		return false
	}

	return result.(bool)
}
