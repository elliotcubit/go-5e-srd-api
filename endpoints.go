package go5e

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://www.dnd5eapi.co/api/"

// Get the bytes for an API request response
func doRequest(query string) ([]byte, error) {
	resp, err := http.Get(baseURL + query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

/*
  Requests the endpoint and attempts to convert the returned json to the struct provided
*/
func doRequestAndUnmarshal(query string, v interface{}) error {
	response, err := doRequest(query)
	if err != nil {
		return err
	}
	json.Unmarshal(response, v)
	return nil
}

/*
  Get the full list of an endpoint in the api
  Options for endpoints:
  spells, ability-scores, skills, proficiencies, languages, classes
  subclasses, features, spellcasting, starting-equipment, races
  subraces, traits, equipment, equipment-categories, weapon-properties
  spells, monsters, conditions, damage-types, magic-schools
*/
func GetResourceList(endpoint string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal(endpoint, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// Get the information for an ability score
func GetAbilityScore(index string) (AbilityScore, error) {
	var ret AbilityScore
	err := doRequestAndUnmarshal("ability-scores/"+index, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// Get the information for a proficiency
func GetProficiency(index string) (Proficiency, error) {
	var ret Proficiency
	err := doRequestAndUnmarshal("proficiencies/"+index, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetClass(index string) (Class, error) {
	var ret Class
	err := doRequestAndUnmarshal("classes/"+index, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetClassSubclasses(index string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("classes/"+index+"/subclasses", &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetSpell(index string) (Spell, error) {
	var ret Spell
	err := doRequestAndUnmarshal("spells/"+index, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
