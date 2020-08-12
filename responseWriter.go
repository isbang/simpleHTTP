package simple

import (
	"html/template"
	"net/http"
)

type ResponseWriter interface {
	http.ResponseWriter

	SetStatus(status int)
	SetHeader(key, val string)
	AddHeader(key, val string)

	Text(status int, contentType string, text string)
	Html(status int, html string)
	Template(status int, t *template.Template, obj interface{})
	Json(status int, obj interface{})
	Xml(status int, obj interface{})
	Yaml(status int, obj interface{})

	Redirect(status int, url string)
}
