package models

type ExoPlanetType string

const (
	GasGiant    ExoPlanetType = "GasGiant"
	Terrestrial ExoPlanetType = "Terrestrial"
)

// Details of an ExoPlanet
type ExoPlanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`       // distance of exoplanet from earth in light years
	Radius      float64       `json:"radius"`         // radius of the planet
	Mass        float64       `json:"mass,omitempty"` // mass of the planet, only in case of Terrestrial type of planet
	Type        ExoPlanetType `json:"type"`
}

// ExoPlanets is an in-memory map to persist the application data
var ExoPlanets = make(map[string]ExoPlanet)
