package go5e

type DamageType struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Name string `json:"name"`
  Desc []string `json:"desc"`
  Url string `json:"url"`
}

func GetDamageType(index string) (DamageType, error) {
  var ret DamageType
  err := doRequestAndUnmarshal("damage-types/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}

func GetDamageTypeByUrl(url string) (DamageType, error) {
	var ret DamageType
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchDamageTypeByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("damage-types/?name="+query, &ret)
	return ret, err
}
