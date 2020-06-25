package go5e

type Equipment struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Name string `json:"name"`
  EquipmentCategory NamedAPIResource `json:"equipment_category"`
  GearCategory string `json:"gear_category"`
  Cost Cost `json:"cost"`
  Weight int `json:"weight"`
  Url string `json:"url"`
}

func GetEquipment(index string) (Equipment, error) {
  var ret Equipment
  err := doRequestAndUnmarshal("equipment/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}
