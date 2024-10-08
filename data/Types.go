package data

import "time"

type Vehicle struct {
	ID        string    `json:"id"`
	Status    string    `json:"current_status"`
	Position  Position  `json:"position"`
	StopId    string    `json:"stop_id"`
	Timestamp time.Time `json:"timestamp"`
	Trip      Trip      `json:"trip"`
}

type Position struct {
	Bearing   float32 `json:"bearing"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Trip struct {
	RouteId string `json:"route_id"`
	TripId  string `json:"trip_id"`
}
