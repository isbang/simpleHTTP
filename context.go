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
	IsAborted() bool
	AbortWithStatus(status int)
	AbortWithError(status int, err error)
	AbortWithBytes(status int, contentType string, b []byte)
	AbortWithText(status int, contentType string, text string)
	AbortWithHtml(status int, html string)
	AbortWithTemplate(status int, t *template.Template, obj interface{}) error
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

type simpleContext struct {
	*requestReader
	*responseWriter

	abort    bool
	chain    []HandlerFunc
	chainIdx int
	abortErr error
}

func (c *simpleContext) Deadline() (deadline time.Time, ok bool) {
	return c.requestReader.req.Context().Deadline()
}

func (c *simpleContext) Done() <-chan struct{} {
	return c.requestReader.req.Context().Done()
}

func (c *simpleContext) Err() error {
	return c.requestReader.req.Context().Err()
}

func (c *simpleContext) Value(key interface{}) interface{} {
	return c.requestReader.req.Context().Value(key)
}

func (c *simpleContext) Next() {
	for {
		if c.IsAborted() || c.chainIdx >= len(c.chain) {
			return
		}
		c.chainIdx++
		c.chain[c.chainIdx](c)
	}
}

func (c *simpleContext) Abort() {
	c.abort = true
}

func (c *simpleContext) IsAborted() bool {
	return c.abort
}

func (c *simpleContext) AbortWithStatus(status int) {
	c.Abort()
	c.SetStatus(status)
}

func (c *simpleContext) AbortWithError(status int, err error) {
	c.Abort()
	c.SetStatus(status)
	c.err = err
}

func (c *simpleContext) AbortWithBytes(status int, contentType string, b []byte) {
	c.Abort()
	c.Byte(status, contentType, b)
}

func (c *simpleContext) AbortWithText(status int, contentType string, text string) {
	c.Abort()
	c.Text(status, contentType, text)
}

func (c *simpleContext) AbortWithHtml(status int, html string) {
	c.Abort()
	c.Html(status, html)
}

func (c *simpleContext) AbortWithTemplate(status int, t *template.Template, obj interface{}) error {
	c.Abort()
	return c.Template(status, t, obj)
}
