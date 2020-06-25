package go5e

type AbilityScore struct {
	Desc     []string           `json:"desc"`
	FullName string             `json:"full_name"`
	Id       string             `json:"_id"`
	Index    string             `json:"index"`
	Name     string             `json:"name"`
	Skills   []NamedAPIResource `json:"skills"`
	Url      string             `json:"url"`
}

func GetAbilityScore(index string) (AbilityScore, error) {
	var ret AbilityScore
	err := doRequestAndUnmarshal("ability-scores/"+index, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
