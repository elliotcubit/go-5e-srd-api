package go5e

type WeaponProperty struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Name string `json:"name"`
  Desc []string `json:"desc"`
  Url string `json:"url"`
}

func GetWeaponProperty(index string) (WeaponProperty, error) {
  var ret WeaponProperty
  err := doRequestAndUnmarshal("weapon-properties/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}
