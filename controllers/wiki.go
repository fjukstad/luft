package controllers

import (
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func wikiTemplateHandler(w http.ResponseWriter, r *http.Request, reqPath string) {
	wikiTemplate, err := template.ParseFiles("views/base.html",
		"views/header.html", "views/navbar.html", "views/navbar2.html",
		"views/wiki.html", "views/footer.html", "views/scripts.html",
		filepath.Join(".", reqPath))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = wikiTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func WikiHandler(w http.ResponseWriter, r *http.Request) {
	reqPath := r.URL.Path
	reqExt := strings.ToUpper(path.Ext(reqPath))
	if reqExt == ".HTML" || reqExt == ".HTM" {
		wikiTemplateHandler(w, r, reqPath)
		return
	}

	http.ServeFile(w, r, filepath.Join(".", reqPath))
}
