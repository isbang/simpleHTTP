package simple

import (
	"net/http"
)

type RequestReader interface {
	Request() *http.Request

	Method() string
	UserAgent() string
	ContentType() string
	BasicAuth() (string, string, bool)
	ClientIP() string

	BindUrl(interface{}) error
	BindHeader(interface{}) error
	BindQuery(interface{}) error
	BindBody(interface{}) error
}
