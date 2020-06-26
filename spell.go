package go5e

type Spell struct {
	Id            string             `json:"_id"`
	Index         string             `json:"index"`
	Name          string             `json:"name"`
	Desc          []string           `json:"desc"`
	HigherLevel   []string           `json:"higher_level"`
	Range         string             `json:"range"`
	Components    []string           `json:"components"`
	Material      string             `json:"material"`
	Ritual        bool               `json:"ritual"`
	Duration      string             `json:"duration"`
	Concentration bool               `json:"concentration"`
	CastingTime   string             `json:"casting_time"`
	Level         int                `json:"level"`
	School        NamedAPIResource   `json:"school"`
	Classes       []NamedAPIResource `json:"classes"`
	Subclasses    []NamedAPIResource `json:"subclasses"`
	Url           string             `json:"url"`
}

func GetSpell(index string) (Spell, error) {
	var ret Spell
	err := doRequestAndUnmarshal("spells/"+index, &ret)
	return ret, err
}

func GetSpellByUrl(url string) (Spell, error) {
	var ret Spell
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchSpellByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("spells/?name="+query, &ret)
	return ret, err
}
