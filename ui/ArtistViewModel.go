package ui

import "matthewhope/groupie-tracker/models"

type ArtistViewModel struct {
	Data           models.Artist
	DatesLocations map[string][]string
}
