package simple

import (
	"bytes"
	"html/template"
	"net/http"
)

type ResponseWriter interface {
	http.ResponseWriter

	SetStatus(status int)
	SetHeader(key, val string)
	AddHeader(key, val string)

	Byte(status int, contentType string, b []byte)
	Text(status int, contentType string, text string)
	Html(status int, html string)
	Template(status int, t *template.Template, obj interface{}) error

	Redirect(status int, url string)
}

type responseWriter struct {
	w      http.ResponseWriter
	status int
	body   *bytes.Buffer
	err    error
}

func (w *responseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *responseWriter) Write(p []byte) (int, error) {
	return w.w.Write(p)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.w.WriteHeader(statusCode)
}

func (w *responseWriter) SetStatus(status int) {
	w.status = status
}

func (w *responseWriter) SetHeader(key, val string) {
	w.w.Header().Set(key, val)
}

func (w *responseWriter) AddHeader(key, val string) {
	w.w.Header().Add(key, val)
}

func (w *responseWriter) Byte(status int, contentType string, b []byte) {
	w.SetStatus(status)
	w.SetHeader(ContentType, contentType)

	w.body.Reset()
	w.body.Write(b)
}

func (w *responseWriter) Text(status int, contentType string, text string) {
	w.SetStatus(status)
	w.SetHeader(ContentType, contentType)

	w.body.Reset()
	w.body.WriteString(text)
}

func (w *responseWriter) Html(status int, html string) {
	w.SetStatus(status)
	w.SetHeader(ContentType, Html)

	w.body.Reset()
	w.body.WriteString(html)
}

func (w *responseWriter) Template(status int, t *template.Template, obj interface{}) error {
	w.SetStatus(status)
	w.SetHeader(ContentType, Html)

	w.body.Reset()
	if err := t.Execute(w.body, obj); err != nil {
		return err
	}
	return nil
}

func (w *responseWriter) Redirect(status int, url string) {
	w.SetStatus(status)
	w.SetHeader(Location, url)
}
