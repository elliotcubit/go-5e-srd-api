package go5e

import (
	"errors"
)

const baseURL = "https://www.dnd5eapi.co/api/"

// Get all endpoints as a string slice
func Endpoints() []string {
	return []string{
		"ability-scores",
		"classes",
		"conditions",
		"damage-types",
		"equipment-categories",
		"equipment",
		"features",
		"languages",
		"magic-schools",
		"monsters",
		"proficiencies",
		"races",
		"skills",
		"spellcasting",
		"spells",
		"starting-equipment",
		"subclasses",
		"subraces",
		"traits",
		"weapon-properties",
	}
}

// Get the full list for an endpoint
func GetResourceList(endpoint string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	if endpoint == "starting-equipment" {
		// ಠ_ಠ
		// The indices for this field are type int instead of string.
		return ret, errors.New("Not implemented due to an API issue with indexing.")
	} else if endpoint == "spellcasting" {
		return ret, errors.New("Not implemented")
	} else {
		err := doRequestAndUnmarshal(endpoint, &ret)
		return ret, err
	}
}
