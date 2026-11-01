package handler

import (
	"encoding/json"
	"net/http"

	"code-one/internal/shortener"
)

type ShortnerPostResponse struct {
	URL string `json:"url"`
}

type ShortnerPostPayload struct {
	URL string `json:"url"`
}

type ShortnerGetResponse struct {
	URL string `json:"url"`
}

func ShortenerPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var body ShortnerPostPayload

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "unable to decode the payload", http.StatusBadRequest)
		return
	}

	if body.URL == "" {
		http.Error(w, "URL member must have value", http.StatusBadRequest)
		return
	}

	s := shortener.Hash(body.URL)
	resp := ShortnerPostResponse{
		URL: s,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode shortner response to json", http.StatusInternalServerError)
		return
	}
}

func ShortenerGet(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("url")
	ogURL, err := shortener.GetOriginalURL(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := ShortnerGetResponse{
		URL: *ogURL,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode original url", http.StatusInternalServerError)
		return
	}

}
