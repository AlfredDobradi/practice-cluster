package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	version         string   = "1.0.2"
	availableColors []string = []string{
		"blue",
		"green",
		"yellow",
		"red",
		"magenta",
	}
)

func main() {
	httpAddress := os.Getenv("COLOR_HTTP_ADDRESS")
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		output, err := content()
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s", "version": "v%s"}`, err.Error(), version), http.StatusInternalServerError)
			return
		}
		w.Write(output) // nolint
	})

	s := http.Server{
		Addr:    httpAddress,
		Handler: r,
	}

	log.Printf("Web service listening on %s", httpAddress)
	log.Fatalln(s.ListenAndServe())
}

func content() ([]byte, error) {
	content := map[string]interface{}{
		"color":   availableColors[0],
		"version": "v" + version,
	}

	raw, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
