package go5e

type Race struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Name string `json:"name"`
  Speed int `json:"speed"`
  AbilityBonuses []NamedAPIResource `json:"ability_bonuses"`
  Alignment string `json:"aligment"`
  Age string `json:"age"`
  Size string `json:"size"`
  SizeDescription string `json:"size_description"`
  StartingProficiencies []NamedAPIResource `json:"starting_proficiencies"`
  StartingProficiencyOptions []NamedChoice `json:"starting_proficiency_options"`
  Languages []NamedAPIResource `json:"languages"`
  LanguageDesc string `json:"language_desc"`
  Traits []NamedAPIResource `json:"traits"`
  Subraces []NamedAPIResource `json:"subraces"`
  Url string `json:"url"`
}

func GetRace(index string) (Race, error) {
  var ret Race
  err := doRequestAndUnmarshal("races/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}

func GetRaceByUrl(url string) (Race, error) {
	var ret Race
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchRaceByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("races/?name="+query, &ret)
	return ret, err
}
