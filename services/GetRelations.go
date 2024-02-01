package services

import (
	"encoding/json"
	"fmt"
	"matthewhope/groupie-tracker/models"
)

func GetRelations() ([]models.Relation, error) {
	var ret models.RelationJSON
	url := fmt.Sprintf("%s%s", baseUrl, relationsEndpoint)
	res, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return ret.Relations, nil
}
