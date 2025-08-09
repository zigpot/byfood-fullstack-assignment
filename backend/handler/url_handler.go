package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type URLRequest struct {
	URL       string `json:"url"`
	Operation string `json:"operation"`
}

type URLResponse struct {
	ProcessedURL string `json:"processed_url"`
}

func (h *BookHandler) ProcessURL(w http.ResponseWriter, r *http.Request) {
	var req URLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	parsed, err := url.Parse(req.URL)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	switch strings.ToLower(req.Operation) {
	case "canonical":
		parsed.RawQuery = ""
		parsed.Path = strings.TrimSuffix(parsed.Path, "/")
	case "redirection":
		parsed.Host = "www.byfood.com"
		parsed.Scheme = strings.ToLower(parsed.Scheme)
		parsed.Path = strings.ToLower(parsed.Path)
	case "all":
		parsed.RawQuery = ""
		parsed.Path = strings.TrimSuffix(parsed.Path, "/")
		parsed.Host = "www.byfood.com"
		parsed.Scheme = strings.ToLower(parsed.Scheme)
		parsed.Path = strings.ToLower(parsed.Path)
	default:
		http.Error(w, "Invalid operation type", http.StatusBadRequest)
		return
	}

	res := URLResponse{ProcessedURL: parsed.String()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
