package main

import "ride-sharing/shared/types" // Replace with the correct module path

type PreviewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
