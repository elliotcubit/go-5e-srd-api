package go5e

type EquipmentCategory struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Name string `json:"name"`
  Equipment []NamedAPIResource `json:"equipment"`
  Url string `json:"url"`
}

func GetEquipmentCategory(index string) (EquipmentCategory, error) {
  var ret EquipmentCategory
  err := doRequestAndUnmarshal("equipment-categories/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}

func GetEquipmentCategoryByUrl(url string) (EquipmentCategory, error) {
	var ret EquipmentCategory
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchEquipmentCategoryByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("equipment-categories/?name="+query, &ret)
	return ret, err
}
