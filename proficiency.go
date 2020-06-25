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
