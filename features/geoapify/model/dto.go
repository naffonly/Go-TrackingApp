package model

type GeoDTO struct {
	Type    string   `json:"type"`
	PickUp  WayPoint `json:"pickUp"`
	DropOff WayPoint `json:"dropOff"`
}

type WayPoint struct {
	Lon string `json:"lon"`
	Lat string `json:"lat"`
}
