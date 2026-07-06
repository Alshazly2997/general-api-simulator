package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(0)

	mux := http.NewServeMux()
	mux.HandleFunc("/bbs-simulate/querybbs", reqOtp)
	mux.HandleFunc("/bbs-simulate/topup", topup)
	// mux.HandleFunc("/bbs-simulate/checkblacklist", checkBlackList)

	port := GetConfigValue("port", "8080")
	fmt.Println("BBS Simulate web service, running on")
	fmt.Println("http://localhost:" + port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
