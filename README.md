# Driver Match

## Overview

Driver Match is a Go-based application that provides two primary services: **Location Service** and **Match Service**.
The Location Service manages driver geolocation data, while the Match Service finds the nearest drivers based on given
coordinates.

## Features

- **Location Service**:
    - Import driver locations from CSV files.
      Features
- Authentication. It only requests from The Match API.
- An endpoint for creating a driver location. It would support
  batch operations to handle the bulk update.
- An endpoint for searching with GeoJSON point and radius
  fields. The matched result must have the distance field from
  the given coordinate.

- **Match Service**:
 - The Matching API matches a suitable driver with the rider
  using Driver Location API. It allows the query to find the nearest
  driver around a given GeoJSON point.
  Features
- Authentication. It only accepts authenticated user requests. If a JWT payload contains authenticated: true field, it can be admitted that the user is authenticated.

location: {
type: "Point",
coordinates: [40.848447, -73.856077] }

- The endpoint that allows searching with a GeoJSON point to find a driver if it matches the given criteria. Otherwise, the service should respond with a 404 - Not Found

## Getting Started

### Prerequisites

- Go (1.17 or higher)
- MongoDB
- Docker (optional)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/burakkuru5534/driver-match.git
   cd driver-match
   
2. Dockerize the services

- cd driver-location
- docker build -t location-service .
- docker run -p 8081:8081 location-service
- cd ../match
- docker build -t match-service .
- docker run -p 8082:8082 match-service
- otherwise both services can be run like this:
- go run main.go

3. location service endpoints

- 	/auth (generate token)
 -   /location (to create new driver location)
  -   /import (import csv files)
  -   /driver/nearest (need token to call this endpoint)

4. match service endpoints

- 	/match/nearest

5. Testing

- go test ./...

## Conclusion 

- Two services implemented. One of them for location service for drivers.
Other one for users (match service).
- User should send auth request first and then can send request for find nearest driver

- circuit breaker design pattern Implemented

- The haversine formula determines the great-circle distance between is used



