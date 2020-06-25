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
