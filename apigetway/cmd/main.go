package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/api/auth").HandlerFunc(authServiceHandler)
	r.PathPrefix("/api/chat").HandlerFunc(chatServiceHandler)
	r.PathPrefix("/api/profile").HandlerFunc(profileServiceHandler)

	corsObj := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	port := ":8080"
	log.Printf("API Gateway запущен на %s\n", port)
	if err := http.ListenAndServe(port, handlers.CORS(corsHeaders, corsMethods, corsObj)(r)); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func authServiceHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8081"+r.URL.Path, http.StatusTemporaryRedirect)
}

func chatServiceHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8082"+r.URL.Path, http.StatusTemporaryRedirect)
}

func profileServiceHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8083"+r.URL.Path, http.StatusTemporaryRedirect)
}
