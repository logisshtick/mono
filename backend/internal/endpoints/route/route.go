package route

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"googlemaps.github.io/maps"
)

var (
	client *maps.Client

	endPoint string
	mlog     *log.Logger
	wlog     *log.Logger
	elog     *log.Logger
)

func Start(ep string, m *log.Logger, w *log.Logger, e *log.Logger) error {
	endPoint = ep
	mlog = m
	wlog = w
	elog = e

	var err error
	client, err = maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if err != nil {
		elog.Printf("Failed to create Google Maps client: %v\n", err)
		return err
	}

	mlog.Printf("%s route module inited\n", endPoint)

	return nil
}

func Stop() error {
	mlog.Printf("%s route module stoped\n", endPoint)
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	mlog.Printf("%s connected %s\n", r.URL.Path, r.RemoteAddr)

	coordinates := []maps.LatLng{
		{Lat: 37.7749, Lng: -122.4194}, // San Francisco
		{Lat: 34.0522, Lng: -118.2437}, // Los Angeles
		{Lat: 47.6062, Lng: -122.3321}, // Seattle
	}

	waypoints := make([]string, 0, 128)
	for _, coord := range coordinates {
		waypoints = append(waypoints, coord.String())
	}

	req := &maps.DirectionsRequest{
		Origin:      waypoints[0],
		Destination: waypoints[len(coordinates)-1],
		Waypoints:   waypoints[1 : len(coordinates)-1],
		Optimize:    true,
	}

	route, _, err := client.Directions(context.Background(), req)
	if err != nil {
		elog.Printf("Failed to request directions: %s", err)
	}

	for i, waypoint := range route[0].Legs {
		fmt.Fprintf(w, "<span>%d. %s</span>\n", i+1, waypoint.StartAddress)
	}
	fmt.Fprintf(w, "<span>%d. %s</span>\n", len(route[0].Legs)+1, route[0].Legs[len(route[0].Legs)-1].EndAddress)
}
