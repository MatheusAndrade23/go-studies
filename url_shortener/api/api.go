package api

import (
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"
	"os"
	"url_shortener/omdb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp) 

	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(
			w, 
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)

	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
	}
}


func NewHandler(db map[string]string) http.Handler {
	apiKey := os.Getenv("OMDB_KEY")

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shortener", handlePost(db))
	r.Get("/{code}", handleGet(db))

	r.Get("/movie/search", handleSearch(apiKey))

	return r
}

type PostBody struct {
	URL string
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data any `json:"data,omitempty"`
}

func handlePost(db map[string]string) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request){
		var body PostBody 

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return 
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(w, Response{Error: "invalid url"}, http.StatusBadGateway)
		}

		code := genCode()
		db[code] = body.URL
		sendJSON(w, Response{Data: code}, http.StatusCreated)
	}
}

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func genCode() string {
	const n = 8
	bytes := make([]byte, 8)

	for i := range n {
		bytes[i] = characters[rand.IntN(len(characters))]
	}

	return string(bytes)
}

func handleGet (db map[string]string) http.HandlerFunc { 
	return func(w http.ResponseWriter, r *http.Request){
		code := chi.URLParam(r, "code")
		data, ok := db[code]

		if !ok {
			http.Error(w, "url nao encontrada", http.StatusNotFound)
		}

		http.Redirect(w, r, data, http.StatusPermanentRedirect)
	}
}

func handleSearch (apiKey string) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("s")
		res, err := omdb.Search(apiKey, search)

		if err != nil {
			sendJSON(
				w,
				Response{Error: "something went wrong with omdb"},
				http.StatusBadGateway,
			)
			return
		}

		sendJSON(w, Response{Data: res}, http.StatusOK)
	}
}
