package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handleSimulationResponse)

	port := "8080"
	fmt.Printf("General api simulator running on http://localhost:%s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}

func handleSimulationResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	endpoint := filepath.Base(r.URL.Path)
	filename := fmt.Sprintf("%s.json", endpoint)

	data, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	fmt.Printf("Served %s for endpoint %s\n", filename, r.URL.Path)
}
