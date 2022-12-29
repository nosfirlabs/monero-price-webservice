package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Coin struct {
	Name  string  `json:"name"`
	Price float64 `json:"price_usd"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.coinmarketcap.com/v1/ticker/monero/")
		if err != nil {
			http.Error(w, "Error retrieving Monero price", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var coin Coin
		if err := json.NewDecoder(resp.Body).Decode(&coin); err != nil {
			http.Error(w, "Error parsing Monero price JSON", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, coin); err != nil {
			http.Error(w, "Error rendering HTML template", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
