package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/bolinwu/test/internal/calc"
)

type sumResponse struct {
	Total int `json:"total"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthzHandler)
	mux.HandleFunc("/sum", sumHandler)
	mux.HandleFunc("/multiply", multiplyHandler)

	addr := ":8080"
	log.Printf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func healthzHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	raw := r.URL.Query().Get("numbers")
	if strings.TrimSpace(raw) == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "numbers query parameter is required"})
		return
	}
	parts := strings.Split(raw, ",")
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		if strings.TrimSpace(p) == "" {
			continue
		}
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: "numbers must be integers"})
			return
		}
		nums = append(nums, n)
	}

	total, err := calc.Sum(nums...)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, sumResponse{Total: total})
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	xStr := query.Get("x")
	yStr := query.Get("y")

	x, errX := strconv.Atoi(strings.TrimSpace(xStr))
	y, errY := strconv.Atoi(strings.TrimSpace(yStr))
	if errX != nil || errY != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "x and y must be integers"})
		return
	}

	writeJSON(w, http.StatusOK, sumResponse{Total: calc.Multi(x, y)})
}

func writeJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("encode response error: %v", err)
	}
}
