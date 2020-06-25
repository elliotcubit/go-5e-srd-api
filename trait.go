package go5e

type Trait struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Races []NamedAPIResource `json:"races"`
  Subraces []NamedAPIResource `json:"subraces"`
  Name string `json:"name"`
  Desc []string `json:"desc"`
  Url string `json:"url"`
}

func GetTrait(index string) (Trait, error) {
  var ret Trait
  err := doRequestAndUnmarshal("traits/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}
