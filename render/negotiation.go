package render

import (
	`net/http`

	"github.com/dimiro1/toolkit-defaults/internal/contenttype"
)

type ContentNegotiationRenderer struct {
	JSONRenderer JSON
	XMLRenderer  XML
	TextRenderer Text
}

func (c ContentNegotiationRenderer) Render(w http.ResponseWriter, r *http.Request, status int, i interface{}) error {
	switch contenttype.Detect(r) {
	case "xml":
		return c.XMLRenderer.Render(w, r, status, i)
	case "json":
		return c.JSONRenderer.Render(w, r, status, i)
	default:
		return c.TextRenderer.Render(w, r, status, i)
	}
}

func NewContentNegotiationRenderer() ContentNegotiationRenderer {
	return ContentNegotiationRenderer{}
}
