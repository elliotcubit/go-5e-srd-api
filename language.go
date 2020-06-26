package go5e

type Language struct {
	Id              string   `json:"_id"`
	Index           string   `json:"index"`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	TypicalSpeakers []string `json:"typical_speakers"`
	Script          string   `json:"script"`
	Url             string   `json:"url"`
}

func GetLanguage(index string) (Language, error) {
  var ret Language
  err := doRequestAndUnmarshal("languages/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}

func GetLanguageByUrl(url string) (Language, error) {
	var ret Language
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchLanguageByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("languages/?name="+query, &ret)
	return ret, err
}
