package go5e

type Condition struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Name string `json:"name"`
  Desc []string `json:"desc"`
  Url string `json:"url"`
}

func GetCondition(index string) (Condition, error) {
  var ret Condition
  err := doRequestAndUnmarshal("conditions/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}
