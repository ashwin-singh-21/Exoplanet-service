package apihandler

import (
	"exo-planet-app/operations"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.Use(loggingMiddleware) // Using the logging middleware to track the api requests.

	// Api routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "welcome to exo-planet-app") })
	router.HandleFunc("/exoplanets", operations.AddExoPlanet).Methods("POST")
	router.HandleFunc("/exoplanets", operations.ListExoPlanets).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", operations.GetExoPlanet).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", operations.UpdateExoPlanet).Methods("PUT")
	router.HandleFunc("/exoplanets/{id}", operations.DeleteExoPlanet).Methods("DELETE")
	router.HandleFunc("/exoplanets/{id}/fuel", operations.FuelEstimation).Methods("GET")

	log.Println("Server started and listening at http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s, %s, %s\n", r.Method, r.RemoteAddr, r.URL)
		next.ServeHTTP(w, r)
	})
}
