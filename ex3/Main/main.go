package main

import (
	"fmt"
	"net/http"
	"os"
)

type Episode struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func readBook(path string) []Episode {

	recordFile, err := os.Open(path)
	defer recordFile.Close()
	if err != nil {
		fmt.Println("An error encounteres ::", err)
	}

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	http.ListenAndServe(":8080", mux)
}
