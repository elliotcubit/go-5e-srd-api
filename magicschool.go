package go5e

type MagicSchool struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Name string `json:"name"`
  Desc string `json:"desc"`
  Url string `json:"url"`
}

func GetMagicSchool(index string) (MagicSchool, error) {
  var ret MagicSchool
  err := doRequestAndUnmarshal("magic-schools/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}

func GetMagicSchoolByUrl(url string) (MagicSchool, error) {
	var ret MagicSchool
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchMagicSchoolByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("magic-schools/?name="+query, &ret)
	return ret, err
}
