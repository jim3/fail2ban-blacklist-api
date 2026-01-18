// Fail2ban Blacklist API

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type Response struct {
	Blacklist []string `json:"blacklist"`
}

// Returns the lastest blacklisted fail2ban IP addresses in JSON format
func getBlacklist(w http.ResponseWriter, r *http.Request) {
	output := exec.Command("sudo", "fail2ban-client", "get", "sshd", "banip")
	out, err := output.Output()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	trimmed := strings.TrimSpace(string(out))
	slc := strings.Fields(trimmed)

	res := Response{
		Blacklist: slc,
	}

	blacklist, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(blacklist)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/blacklist", getBlacklist)

	s := http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
