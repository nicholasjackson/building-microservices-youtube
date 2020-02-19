package handlers

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

// GZipResponseMiddleware detects if the client can handle
// zipped content and if so returns the response in GZipped format
func GZipResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// try and determine the content type
		fmt.Println("File")

		// if client cant handle gzip send plain
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			//f.log.Debug("Unable to handle gzipped", "file", fp)
			fmt.Println("Not gzip")
			next.ServeHTTP(rw, r)
			return
		}

		// client can handle gziped content send gzipped to speed up transfer
		// set the content encoding header for gzip
		rw.Header().Add("Content-Encoding", "gzip")

		// file server sets the content stream
		// nice
		//rw.Header().Add("Content-Type", "application/octet-stream")

		wr := NewWrappedResponseWriter(rw)
		defer wr.Flush()

		// write the file
		next.ServeHTTP(wr, r)
	})
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {

	// wrap the default writer in a gzip writer
	gw := gzip.NewWriter(rw)

	return &WrappedResponseWriter{rw, gw}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *WrappedResponseWriter) WriteHeader(statusCode int) {
	wr.rw.WriteHeader(statusCode)
}

func (wr *WrappedResponseWriter) Flush() {
	// flush and close the writer
	wr.gw.Flush()
	wr.gw.Close()
}
