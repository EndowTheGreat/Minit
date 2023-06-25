package http

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/go-chi/chi"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	h.DB.View(func(txn *badger.Txn) error {
		key, err := txn.Get([]byte(chi.URLParam(r, "uri")))
		if err != nil {
			return err
		}
		uri, err := key.ValueCopy(nil)
		if err != nil {
			return err
		}
		http.Redirect(w, r, string(uri), 301)
		return nil
	})
}

func (h *Handler) ShortenURI(w http.ResponseWriter, r *http.Request) {
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := url.ParseRequestURI(request.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.DB.Update(func(txn *badger.Txn) error {
		key := newKey(6)
		txn.SetEntry(badger.NewEntry(key, []byte(request.Body)).WithTTL(time.Hour * 24))
		w.Write(key)
		return nil
	})
}

func newKey(size int) []byte {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	key := make([]byte, size)
	for i := 0; i < size; i++ {
		key[i] = chars[rand.Intn(len(chars))]
	}
	return key
}
