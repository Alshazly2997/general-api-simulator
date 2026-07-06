package main

import (
	"log"
	"net/http"
	"os"
)

func reqOtp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := GetConfigValue("prepaid-query", "")
	data, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	log.Printf("served %s for endpoint %s\n", filename, r.URL.Path)
}

func topup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := GetConfigValue("recharge-result", "")
	data, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	log.Printf("served %s for endpoint %s\n", filename, r.URL.Path)
}

/*func checkBlackList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := GetConfigValue("blacklist-result", "")
	data, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	log.Printf("served %s for endpoint %s\n", filename, r.URL.Path)
}*/
