package simple

import (
	"context"
	"html/template"
	"time"
)

type CancelFunc = context.CancelFunc

type Context interface {
	context.Context
	RequestReader
	ResponseWriter

	Next()
	Abort()
	AbortWithStatus(status int)
	AbortWithError(status int, err error)
	AbortWithText(status int, contentType string, text string)
	AbortWithHtml(status int, html string)
	AbortWithTemplate(status int, t *template.Template, obj interface{})
	AbortWithJson(status int, obj interface{})
	AbortWithXml(status int, obj interface{})
	AbortWithYaml(status int, obj interface{})
}

func WithValue(ctx Context, key, val interface{}) Context {
	panic("not implemented")
}

func WithCancel(ctx Context) (Context, CancelFunc) {
	panic("not implemented")
}

func WithTimeout(ctx Context, d time.Duration) (Context, CancelFunc) {
	panic("not implemented")
}

func WithDeadline(ctx Context, t time.Time) (Context, CancelFunc) {
	panic("not implemented")
}
