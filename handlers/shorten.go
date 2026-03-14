package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AnuragDahiwade/url-shortner/database"
	"github.com/AnuragDahiwade/url-shortner/utils"
)

type request struct {
	URL string `json:"url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var id int

	query := `INSERT INTO urls (long_url) VALUES ($1) RETURNING id`

	err = database.DB.QueryRow(query, req.URL).Scan(&id)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	shortCode := utils.EncodeBase62(id)

	updateQuery := `UPDATE urls SET short_code=$1 WHERE id=$2`
	_, err = database.DB.Exec(updateQuery, shortCode, id)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"short_url": "http://localhost:8080/" + shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
