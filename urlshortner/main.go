package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

// gdgdhhd (

// )

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println(hasher)
	data := hasher.Sum(nil)
	fmt.Println(data)
	hash := hex.EncodeToString(data)
	fmt.Println(hash[:8])

	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("failed")
	}
	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	fmt.Println("welcome")
	OriginalURL := "https://github.com/yashkumar1998/batch"
	generateShortURL(OriginalURL)

	http.HandleFunc("/", handler)

	// start server part
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("error there", err)
	}
}
