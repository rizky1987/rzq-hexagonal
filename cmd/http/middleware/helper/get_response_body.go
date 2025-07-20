package middleware_helper

import (
	"bytes"
	"net/http"
)

type ResponseCaptureWriter struct {
	http.ResponseWriter
	Body *bytes.Buffer
}

func (w *ResponseCaptureWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}
