package utils

import (
	"exo-planet-app/models"
	"net/http"
)

func HandleHttpMethodErr(w http.ResponseWriter) error {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return nil
}

func ValidateExoPlanetDetails(w http.ResponseWriter, details models.ExoPlanet) error {

	if details.Distance < 10 || details.Distance > 1000 {
		http.Error(w, "distance must be between 10 to 1000 light years", http.StatusBadRequest)
		return nil
	}

	if details.Radius < 0.1 || details.Radius > 10 {
		http.Error(w, "radius must be between 0.1 and 10 Earth-radius units", http.StatusBadRequest)
		return nil
	}

	if details.Type == models.Terrestrial && (details.Mass < 0.1 || details.Mass > 10) {
		http.Error(w, "mass must be between 0.1 and 10 Earth-mass units", http.StatusBadRequest)
		return nil
	}

	return nil
}
