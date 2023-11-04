package update

type Geolocation struct {
	Type        string      `json:"type,omitempty"`
	Coordinates Coordinates `json:"coordinates,omitempty"`
	Place       Place       `json:"place,omitempty"`
	ShowMap     string      `json:"showmap,omitempty"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type Place struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Latitude  int    `json:"latitude,omitempty"`
	Longitude int    `json:"longitude,omitempty"`
	Created   int    `json:"created,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Country   string `json:"country,omitempty"`
	City      string `json:"city,omitempty"`
}
