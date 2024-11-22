package main

import (
	"encoding/json"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/")))

	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		username := q.Get("username")
		password := q.Get("password")
		success, token := Login(username, password)

		w.Header().Set("Content-Type", "application/json")
		if !success {
			w.WriteHeader(http.StatusUnauthorized)
		}

		bytes, err := json.Marshal(map[string]interface{}{
			"success": success,
			"error":   token,
		})
		check(err)

		w.Write(bytes)
	})

	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		username := q.Get("username")
		password := q.Get("password")
		success, token := Register(username, password)

		w.Header().Set("Content-Type", "application/json")
		if !success {
			w.WriteHeader(http.StatusUnauthorized)
		}

		bytes, err := json.Marshal(map[string]interface{}{
			"success": success,
			"error":   token,
		})
		check(err)

		w.Write(bytes)
	})

	http.HandleFunc("/api/zoltan/list", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.Marshal(&zoltanRegistry)
		check(err)

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.ListenAndServe(":3005", nil)
}
