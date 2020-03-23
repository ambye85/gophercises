package cyoa

import (
	"html/template"
	"log"
	"net/http"
)

func init() {
	tpl = template.Must(template.New("default").Parse(defaultTemplate))
}

var tpl *template.Template

var defaultTemplate = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>Choose Your Own Adventure</title>
</head>
<body>
	<section class="page">
		<h1>{{ .Title }}</h1>
		{{ range .Paragraphs }}
		<p>{{ . }}</p>
		{{ end }}
		<ul>
			{{ range .Options }}
			<li><a href="/{{ .Chapter }}">{{ .Text }}</a></li>
			{{ end }}
		</ul>
	</section>
	<style>
		body {
			font-family: helvetica, arial;
		}
		h1 {
			text-align:center;
			position:relative;
		}
		.page {
			width: 80%;
			max-width: 500px;
			margin: auto;
			margin-top: 40px;
			margin-bottom: 40px;
			padding: 80px;
			background: #FFFCF6;
			border: 1px solid #eee;
			box-shadow: 0 10px 6px -6px #777;
		}
		ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
		}
		li {
			padding-top: 10px;
		}
		a, a:visited {
			text-decoration: none;
			color: #6295b5;
		}
		a:active, a:hover {
			color: #7792a2;
		}
		p {
			text-indent: 1em;
		}
	</style>
</body>
</html>`

type HandlerOption func(h *storyHandler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *storyHandler) {
		h.template = t
	}
}

func WithChapterNameParser(fn func(r *http.Request) string) HandlerOption {
	return func(h *storyHandler) {
		h.chapterNameParser = fn
	}
}

func defaultChapterNameParser(r *http.Request) string {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

func CreateHandler(s Story, opts ...HandlerOption) http.Handler {
	handler := storyHandler{story: s, template: tpl, chapterNameParser: defaultChapterNameParser}
	for _, opt := range opts {
		opt(&handler)
	}
	return handler
}

type storyHandler struct {
	story             Story
	template          *template.Template
	chapterNameParser func(r *http.Request) string
}

func (h storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chapterName := h.chapterNameParser(r)
	if chapter, ok := h.story[chapterName]; ok {
		err := h.template.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}
