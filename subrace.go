package go5e

type Subrace struct {
	Id                         string             `json:"_id"`
	Index                      string             `json:"index"`
	Name                       string             `json:"name"`
	Race                       NamedAPIResource   `json:"race"`
	Desc                       string             `json:"desc"`
	AbilityBonuses             []NamedAPIResource `json:"ability_bonuses"`
	StartingProficiencies      []NamedAPIResource `json:"starting_proficiencies"`
	StartingProficiencyOptions NamedChoice        `json:"starting_proficiency_options"`
	Languages                  []NamedAPIResource `json:"languages"`
	LanguageDesc               string             `json:"language_desc"`
	LanguageOptions            NamedChoice        `json:"language_options"`
	RacialTraits               []NamedAPIResource `json:"racial_traits"`
	RacialTraitOptions         NamedChoice        `json:"racial_trait_options"`
	Url                        string             `json:"url"`
}

func GetSubrace(index string) (Subrace, error) {
	var ret Subrace
	err := doRequestAndUnmarshal("subraces/"+index, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetSubraceByUrl(url string) (Subrace, error) {
	var ret Subrace
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchSubraceByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("subraces/?name="+query, &ret)
	return ret, err
}
