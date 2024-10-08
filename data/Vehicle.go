package data

import (
	"time"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
)

func NewVehicleFromFeed(entity *gtfs.FeedEntity) (*Vehicle, error) {
	vehicle := entity.GetVehicle()
	position := vehicle.GetPosition()
	trip := vehicle.GetTrip()

	timestamp := int64(vehicle.GetTimestamp())

	return &Vehicle{
		ID:     vehicle.GetVehicle().GetId(),
		Status: vehicle.GetCurrentStatus().String(),
		Position: Position{
			Bearing:   position.GetBearing(),
			Latitude:  position.GetLatitude(),
			Longitude: position.GetLongitude(),
		},
		StopId:    vehicle.GetStopId(),
		Timestamp: time.Unix(timestamp, 0),
		Trip: Trip{
			RouteId: trip.GetRouteId(),
			TripId:  trip.GetTripId(),
		},
	}, nil
}
