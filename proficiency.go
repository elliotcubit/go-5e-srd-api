package go5e

type Proficiency struct {
	Id      string             `json:"_id"`
	Index   string             `json:"index"`
	Type    string             `json:"type"`
	Name    string             `json:"name"`
	Classes []NamedAPIResource `json:"classes"`
	Races   []NamedAPIResource `json:"races"`
	Url     string             `json:"url"`
}

func GetProficiency(index string) (Proficiency, error) {
  var ret Proficiency
  err := doRequestAndUnmarshal("proficiencies/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}

func GetProficiencyByUrl(url string) (Proficiency, error) {
	var ret Proficiency
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchProficiencyByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("proficiencies/?name="+query, &ret)
	return ret, err
}
