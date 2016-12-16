package del

import (
	"io"
	"net/http/httptest"

	"github.com/unrolled/render"
)

type Templates struct {
	renderer *render.Render
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	res := httptest.NewRecorder()
	t.renderer.HTML(res, 200, name, data)
	_, err := w.Write(res.Body.Bytes())
	return err
}

var DefaultOptions = render.Options{
	Directory:       "templates",                                      // Specify what path to load the templates from.
	Layout:          "application",                                    // Specify a layout template. Layouts can call {{ yield }} to render the current template.
	Extensions:      []string{".tmpl", ".html"},                       // Specify extensions to load for templates.
	Charset:         "UTF-8",                                          // Sets encoding for json and html content-types. Default is "UTF-8".
	IndentJSON:      true,                                             // Output human readable JSON.
	IndentXML:       true,                                             // Output human readable XML.
	PrefixXML:       []byte("<?xml version='1.0' encoding='UTF-8'?>"), // Prefixes XML responses with the given bytes.
	HTMLContentType: "application/xhtml+xml",                          // Output XHTML content type instead of default "text/html".
	IsDevelopment:   true,                                             // Render will now recompile the templates on every HTML response.
	UnEscapeHTML:    true,                                             // Replace ensure '&<>' are output correctly (JSON only).
	StreamingJSON:   true,                                             // Streams the JSON response via json.Encoder.
}

// Render(w io.Writer, name string, data interface{}) error
func New(opts render.Options) *Templates {
	r := render.New(opts)
	return &Templates{r}
}
