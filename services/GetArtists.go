package services

import (
	"encoding/json"
	"fmt"
	"matthewhope/groupie-tracker/models"
)

func GetArtists() ([]models.Artist, error) {
	var ret []models.Artist
	url := fmt.Sprintf("%s%s", baseUrl, artistsEndpoint)
	res, err := Fetch(url)
	if err != nil {
		return ret, err
	}
	err = json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
