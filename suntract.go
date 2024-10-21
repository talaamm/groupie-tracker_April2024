package main

type AllLocations struct {
	Index []struct {
		Locations []string `json:"locations"`
	} `json:"index"`
}
type fnan struct {
	Id         int      `json:"id"`
	Image      string   `json:"image"`
	Name       string   `json:"name"`
	Members    []string `json:"members"`
	Creation   int      `json:"creationDate"`
	FirstAlbum string   `json:"firstAlbum"`
}
type locaco struct {
	Locations []string `json:"locations"`
}
type taree5 struct {
	Dates []string `json:"dates"`
}
type realtionship struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}
type Details struct {
	Artists           fnan
	Locations         locaco
	Concertdates      taree5
	DatesAndLocations realtionship
}
