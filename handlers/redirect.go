package handlers

import (
	"net/http"
	"strings"

	"github.com/AnuragDahiwade/url-shortner/database"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {

	code := strings.TrimPrefix(r.URL.Path, "/")

	if code == "" || code == "shorten" {
		http.NotFound(w, r)
		return
	}

	var longURL string

	query := `SELECT long_url FROM urls WHERE short_code=$1`

	err := database.DB.QueryRow(query, code).Scan(&longURL)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
