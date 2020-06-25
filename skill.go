package go5e

type Skill struct {
  Id string `json:"_id"`
  Index string `json:"index"`
  Name string `json:"name"`
  Desc []string `json:"desc"`
  AbilityScore NamedAPIResource `json:"ability_score"`
  Url string `json:"url"`
}

func GetSkill(index string) (Skill, error) {
  var ret Skill
  err := doRequestAndUnmarshal("skills/"+index, &ret)
  if err != nil {
    return ret, err
  }
  return ret, nil
}
