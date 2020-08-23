package simple

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"

	"github.com/isbang/simple/pkg/pattern"
)

type RequestReader interface {
	io.ReadCloser

	Request() *http.Request
	Method() string
	UserAgent() string
	ContentType() string
	BasicAuth() (string, string, bool)
	ClientIP() string

	BindUrl(interface{}) error
	BindHeader(interface{}) error
	BindQuery(interface{}) error
}

type requestReader struct {
	req     *http.Request
	pattern pattern.Pattern
}

func (r requestReader) Read(p []byte) (n int, err error) {
	return r.req.Body.Read(p)
}

func (r requestReader) Close() error {
	return r.req.Body.Close()
}

func (r requestReader) Request() *http.Request {
	return r.req
}

func (r requestReader) Method() string {
	return r.req.Method
}

func (r requestReader) UserAgent() string {
	return r.req.UserAgent()
}

func (r requestReader) ContentType() string {
	return r.req.Header.Get(ContentType)
}

func (r requestReader) BasicAuth() (string, string, bool) {
	return r.req.BasicAuth()
}

func (r requestReader) ClientIP() string {
	if v := r.req.Header.Get(XForwardedFor); v != "" {
		return v
	}
	return strings.Split(r.req.RemoteAddr, ":")[0]
}

func (r requestReader) BindUrl(i interface{}) error {
	d := schema.NewDecoder()
	d.SetAliasTag("url")
	d.IgnoreUnknownKeys(true)
	return d.Decode(i, r.pattern.GetMatched(r.req.URL.Path))
}

func (r requestReader) BindHeader(i interface{}) error {
	d := schema.NewDecoder()
	d.SetAliasTag("header")
	d.IgnoreUnknownKeys(true)
	return d.Decode(i, r.req.Header)
}

func (r requestReader) BindQuery(i interface{}) error {
	d := schema.NewDecoder()
	d.SetAliasTag("query")
	d.IgnoreUnknownKeys(true)

	v, err := url.ParseQuery(r.req.URL.RawQuery)
	if err != nil {
		return err
	}
	return d.Decode(i, v)
}
