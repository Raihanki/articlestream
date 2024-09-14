package handlers

import "net/http"

func GetHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Welcome to ArticleStream"))
}
