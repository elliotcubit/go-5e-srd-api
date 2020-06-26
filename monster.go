package go5e

// Of course! This is the one that is
// Completely fucking undocumented!!!
// LOL!!! haha .... xd!!!
type Monster struct {
	Id                    string             `json:"_id"`
	Index                 string             `json:"index"`
	Name                  string             `json:"name"`
	Size                  string             `json:"size"`
	Type                  string             `json:"type"`
	Subtype               string             `json:"subtype"`
	Alignment             string             `json:"alignment"`
	ArmorClass            int                `json:"armor_class"`
	HitPoints             int                `json:"hit_points"`
	HitDice               string             `json:"hit_dice"`
	Speed                 Speed              `json:"speed"`
	Strength              int                `json:"strength"`
	Dexterity             int                `json:"dexterity"`
	Constitution          int                `json:"constitution"`
	Intelligence          int                `json:"intelligence"`
	Wisdom                int                `json:"wisdom"`
	Charisma              int                `json:"charisma"`
	Proficiencies         []NamedAPIResource `json:"proficiencies"`
	DamageVulnerabilities []string           `json:"damage_vulnerabilities"`
	DamageResistances     []string           `json:"damage_resistances"`
	DamageImmunities      []string           `json:"damage_immunities"`
	ConditionImmunities   []NamedAPIResource `json:"condition_immunities"`
	Senses                Senses             `json:"senses"`
	// Yes, this is really a string, and not NamedAPIResource
	Languages        string    `json:"languages"`
	ChallengeRating  int       `json:"challenge_rating"`
	SpecialAbilities []Ability `json:"special_abilities"`
	LegendaryActions []Ability `json:"legendary_actions"`
	Url              string    `json:"url"`
}

func GetMonster(index string) (Monster, error) {
	var ret Monster
	err := doRequestAndUnmarshal("monsters/"+index, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetMonsterByUrl(url string) (Monster, error) {
	var ret Monster
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchMonsterByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("monsters/?name="+query, &ret)
	return ret, err
}
