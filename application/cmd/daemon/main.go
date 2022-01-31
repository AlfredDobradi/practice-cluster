package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	version         string   = "1.1.0"
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
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s", "version": "v%s"}`, err.Error(), version))) // nolint
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
	index := rand.Intn(len(availableColors) + 1)

	// Simulate unhappy path every once in a while
	if index >= len(availableColors) {
		return nil, fmt.Errorf("Index out of range: %d", index)
	}

	content := map[string]interface{}{
		"color":   availableColors[index],
		"version": "v" + version,
	}

	raw, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
