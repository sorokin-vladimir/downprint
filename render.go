package main

import (
	"bytes"
	_ "embed"
	"html/template"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

//go:embed assets/style.css
var styleCSS string

// Style order matters: built-in CSS, then the user stylesheet, then the page
// size rule, so explicit flags win over both stylesheets.
var pageTmpl = template.Must(template.New("page").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<base href="{{.Base}}">
<title>{{.Title}}</title>
<style>{{.CSS}}</style>
{{- if .UserCSS}}
<link rel="stylesheet" href="{{.UserCSS}}">
{{- end}}
{{- if .PageCSS}}
<style>{{.PageCSS}}</style>
{{- end}}
</head>
<body>
<article class="markdown-body">
{{.Body}}
</article>
</body>
</html>
`))

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithStyle("github"),
		),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		// Allow raw HTML from the source document, as GitHub does.
		html.WithUnsafe(),
	),
)

type pageOptions struct {
	Title   string
	BaseDir string // resolves relative links and images
	UserCSS string // absolute path to an extra stylesheet, optional
	PageCSS string // generated @page rule, optional
}

// renderHTML converts markdown to a standalone HTML page.
func renderHTML(src []byte, opts pageOptions) ([]byte, error) {
	var body bytes.Buffer
	if err := md.Convert(src, &body); err != nil {
		return nil, err
	}

	var userCSS template.URL
	if opts.UserCSS != "" {
		userCSS = fileURL(opts.UserCSS)
	}

	var out bytes.Buffer
	err := pageTmpl.Execute(&out, struct {
		Base    template.URL
		Title   string
		CSS     template.CSS
		UserCSS template.URL
		PageCSS template.CSS
		Body    template.HTML
	}{
		Base:    fileURL(opts.BaseDir + "/"),
		Title:   opts.Title,
		CSS:     template.CSS(styleCSS),
		UserCSS: userCSS,
		PageCSS: template.CSS(opts.PageCSS),
		Body:    template.HTML(body.String()),
	})
	return out.Bytes(), err
}

// fileURL builds a file:// URL for a local path. html/template rejects
// file:// URLs unless they are marked as trusted.
func fileURL(path string) template.URL {
	p := filepath.ToSlash(path)
	// Windows paths (C:/...) need a leading slash, or "C:" becomes the host.
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	u := url.URL{Scheme: "file", Path: p}
	return template.URL(u.String())
}
