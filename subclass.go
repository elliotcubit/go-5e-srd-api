package go5e

type Subclass struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Class NamedAPIResource `json:"class"`
  Name string `json:"name"`
  SubclassFlavour string `json:"subclass_flavor"`
  Desc []string `json:"desc"`
  Features []NamedAPIResource `json:"features"`
  Url string `json:"url"`
}

func GetSubclass(index string) (Subclass, error) {
  var ret Subclass
  err := doRequestAndUnmarshal("subclasses/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}
