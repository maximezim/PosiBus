package data

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
)

var (
	vehiclePos = "https://proxy.transport.data.gouv.fr/resource/lemet-metz-gtfs-rt-vehicle-position"
	tripUpdate = "https://proxy.transport.data.gouv.fr/resource/lemet-metz-gtfs-rt-trip-update"
)

func GetVehiclePos() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", vehiclePos, nil)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	HandleError(err)
	body, err := io.ReadAll(resp.Body)
	HandleError(err)
	feed := gtfs.FeedMessage{}
	err = proto.Unmarshal(body, &feed)
	HandleError(err)
	for _, entity := range feed.Entity {
		vehicle, _ := NewVehicleFromFeed(entity)
		fmt.Println(vehicle, "\n")
	}
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
