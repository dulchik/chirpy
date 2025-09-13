package main

import (
	"net/http"
	"log"
)

func main() {	
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", healthHandler)

	

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	} 

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}


}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.WriteHeader(200)

	w.Write([]byte("OK"))
}
