package ui

import (
	"log"
	"matthewhope/groupie-tracker/services"
)

var (
	IndexVM IndexViewModel
)

func init() {
	artists, err := services.GetArtists()
	if err != nil {
		log.Fatal(err.Error())
	}
	// fmt.Println(artists)
	relations, err := services.GetRelations()
	if err != nil {
		log.Fatal(err.Error())
	}
	size := len(artists)
	if len(relations) != size {
		log.Fatal("failed to retrieve all data")
	}
	artistVMs := make([]ArtistViewModel, size)
	for i := 0; i < size; i++ {
		artistVMs[i] = ArtistViewModel{
			Data:           artists[i],
			DatesLocations: relations[i].DatesLocations,
		}
	}
	IndexVM = IndexViewModel{Artists: artistVMs}
}
