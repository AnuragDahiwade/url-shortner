package handlers

import (
	"net/http"
	"strings"

	"github.com/AnuragDahiwade/url-shortner/database"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/shorten" {
		http.NotFound(w, r)
		return
	}

	code := strings.TrimPrefix(r.URL.Path, "/")

	if code == "" {
		http.ServeFile(w, r, "./static/index.html")
		return
	}

	var longURL string

	err := database.DB.QueryRow(
		"SELECT long_url FROM urls WHERE short_code=$1",
		code,
	).Scan(&longURL)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
