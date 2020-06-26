package go5e

type Feature struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Name string `json:"name"`
  Level int `json:"level"`
  Class NamedAPIResource `json:"class"`
  Subclass NamedAPIResource `json:"subclass"`
  Desc []string `json:"desc"`
  Url string `json:"url"`
}

func GetFeature(index string) (Feature, error) {
  var ret Feature
  err := doRequestAndUnmarshal("features/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}

func GetFeatureByUrl(url string) (Feature, error) {
	var ret Feature
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchFeatureByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("features/?name="+query, &ret)
	return ret, err
}
