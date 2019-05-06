package retry

import "net/http"

// responseWriter implements the http.ResponseWriter interface but saves the status code
// that is written for the purpose of deciding whether or not to retry the request.
type responseWriter struct {
	w http.ResponseWriter
	s int
	b []byte
}

// Header proxies to the underlying ResponseWriter Header() function
func (w *responseWriter) Header() http.Header {
	return w.w.Header()
}

// Write proxies to the underlying ResponseWriter Write() function
func (w *responseWriter) Write(b []byte) (int, error) {
	w.b = b
	return len(b), nil
}

// WriteHeader saves the status code and then proxies
// to the underlying ResponseWriter WriteHeader() function
func (w *responseWriter) WriteHeader(statusCode int) {
	w.s = statusCode
}

// Reset clears the headers that have been queued up for the
// client and should be called before each retry is invoked
func (w *responseWriter) Reset() {
	for k := range w.w.Header() {
		w.w.Header().Del(k)
	}
}

// Flush writes the buffered content to the client
func (w *responseWriter) Flush() (int, error) {
	w.w.WriteHeader(w.s)
	return w.w.Write(w.b)
}
