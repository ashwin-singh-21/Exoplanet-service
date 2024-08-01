package operations

import (
	"encoding/json"
	"exo-planet-app/models"
	"exo-planet-app/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func generateId() string {
	return uuid.New().String()
}

// AddExoPlanet adds an exo-planet
func AddExoPlanet(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received a request to add an Exoplanet!")

	var exoplanet models.ExoPlanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate the exoplanet details such as distance, radius, mass, etc.
	utils.ValidateExoPlanetDetails(w, exoplanet)

	exoplanet.ID = generateId() // set an unique id for an exoplanet

	// create an exoplanet with given details
	models.ExoPlanets[exoplanet.ID] = exoplanet

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)

	log.Println("Successfully added an exoplanet")
}

// ListExoPlanets gets a list of all available exoplanets.
func ListExoPlanets(w http.ResponseWriter, r *http.Request) {

	log.Println("Received a request to list all the exoplanets!")

	exoplanetList := make([]models.ExoPlanet, 0, len(models.ExoPlanets))

	for _, v := range models.ExoPlanets {
		exoplanetList = append(exoplanetList, v)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exoplanetList)

	log.Println("Successfully fetched exoplanets list")
}

// GetExoPlanet returns an exoplanet against the requested id
func GetExoPlanet(w http.ResponseWriter, r *http.Request) {

	log.Println("Received a request to get details of an exoplanet")

	params := mux.Vars(r)
	id := params["id"]

	log.Println("received a request to get details of exoplanet with id:", id)

	exoplanet, exists := models.ExoPlanets[id]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		log.Println("error: Exoplanet not found with id:", id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exoplanet)

	log.Println("Successfully fetched the exoplanet with id:", id)
}

// UpdateExoPlanet updates the details of an exoplanet against requested id
func UpdateExoPlanet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	log.Println("Received a request to update details of exoplanet with id:", id)

	exoplanet, exist := models.ExoPlanets[id]
	if !exist {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		log.Println("error: Exoplanet not found with id:", id)
		return
	}

	var updateExoplanet models.ExoPlanet

	err := json.NewDecoder(r.Body).Decode(&updateExoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("error while decoding update details of an exoplanet")
		return
	}

	// validate the exoplanet details such as distance, radius, mass, etc.
	utils.ValidateExoPlanetDetails(w, updateExoplanet)

	// Update the details of the exoplanet
	exoplanet.Name = updateExoplanet.Name
	exoplanet.Description = updateExoplanet.Description
	exoplanet.Distance = updateExoplanet.Distance
	exoplanet.Radius = updateExoplanet.Radius
	exoplanet.Mass = updateExoplanet.Mass
	exoplanet.Type = updateExoplanet.Type

	models.ExoPlanets[id] = exoplanet

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exoplanet)

	log.Println("Successfully updated the details of exoplanet with id:", id)
}

// DeleteExoPlanet deletes the requested exoplanet from the database
func DeleteExoPlanet(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	id := param["id"]

	log.Println("Received a request to delete an exoplanet with id:", id)

	_, exists := models.ExoPlanets[id]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		log.Println("error: Exoplanet not found with id:", id)
		return
	}

	delete(models.ExoPlanets, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	log.Println("Successfully deleted the exoplanet with id:", id)
}

// FuelEstimation calculate the amount of fuel required to reach an exoplanet
func FuelEstimation(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	log.Println("Received a request to find fuel estimation to reach an exoplanet")

	exoplanet, exists := models.ExoPlanets[id]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		log.Println("error: Exoplanet not found with id:", id)
		return
	}

	// Get the crew capacity form request
	crewCapacityStr := r.URL.Query().Get("crew-capacity")
	if crewCapacityStr == "" {
		http.Error(w, "Missing crew capacity", http.StatusBadRequest)
		log.Println("Missing crew capacity")
		return
	}

	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		log.Println("Found invalid crew capacity")
		return
	}

	// Calculate the gravity based on exoplanet type
	var gravity float64

	switch exoplanet.Type {

	case models.Terrestrial:
		gravity = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)

	case models.GasGiant:
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)

	default:
		http.Error(w, "Invalid exoplanet type", http.StatusBadRequest)
		log.Println("Invalid exoplanet type")
		return
	}

	// Now calculate the fuel estimation based on crew capacity and exoplanet type
	fuel := float64(exoplanet.Distance) / (gravity * gravity) * float64(crewCapacity)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{"estimated_fuel": fuel})
	log.Println("Successfully calculated the fuel estimation to reach the exoplanet")
}
