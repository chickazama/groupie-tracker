package ui

import "matthewhope/groupie-tracker/models"

type ArtistViewModel struct {
	Data           models.Artist
	DatesLocations map[string][]string
}

func (a ArtistViewModel) Prev() int {
	res := (a.Data.ID - 1) % 52
	if res == 0 {
		res = 52
	}
	return res
}

func (a ArtistViewModel) Next() int {
	res := (a.Data.ID + 1) % 52
	if res == 0 {
		res = 52
	}
	return res
}
