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
	return ret, err
}

func GetAbilityScoreByUrl(url string) (AbilityScore, error) {
	var ret AbilityScore
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchAbilityScoreName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("ability-scores/?name="+query, &ret)
	return ret, err
}
