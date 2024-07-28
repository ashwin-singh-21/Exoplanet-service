# Exoplanet Application

## Overview

This microservice supports space voyagers in managing and studying different exoplanets (planets outside of our solar system). Currently, the service supports two types of exoplanets:
1. **Gas Giant**: Composed of only gaseous compounds.
2. **Terrestrial**: Earth-like planets, more rocky and larger than Earth.

## Features

The microservice provides the following functionalities:
1. **Add an Exoplanet**: Add a new exoplanet with the following properties:
   - `name`
   - `description`
   - `distance` (from Earth in light years)
   - `radius` (in Earth-radius units)
   - `mass` (only for Terrestrial type, in Earth-mass units)
   - `type`: `GasGiant` or `Terrestrial`
   
2. **List Exoplanets**: Retrieve a list of all available exoplanets.
3. **Get Exoplanet by ID**: Retrieve information about a specific exoplanet by its unique ID.
4. **Update Exoplanet**: Update the details of an existing exoplanet.
5. **Delete Exoplanet**: Remove an exoplanet from the catalog.
6. **Fuel Estimation**: Retrieve an overall fuel cost estimation for a trip to any particular exoplanet for a given crew capacity.

### Prerequisites

- Go 1.21 or later
- Docker

### Building and Running

1. **Build the application:**
    ```sh
    go build -o exoplanet-app main.go
    ```

2. **Run the application:**
    ```sh
    ./exoplanet-app
    ```

### Using Docker

1. **Build the Docker image:**
    ```sh
    docker build -t exoplanet-app .
    ```

2. **Run the Docker container:**
    ```sh
    docker run -n exoplanet -d -p 8080:8080 exoplanet-app 
    ```
