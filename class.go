package go5e

type Class struct {
	ClassLevels        ClassAPIResource   `json:"class_levels"`
	HitDie             int                `json:"hit_die"`
	Id                 string             `json:"_id"`
	Index              string             `json:"index"`
	Name               string             `json:"name"`
	Proficiencies      []NamedAPIResource `json:"proficiencies"`
	ProficiencyChoices []NamedChoice      `json:"proficiency_choices"`
	SavingThrows       []NamedAPIResource `json:"saving_throws"`
	Spellcasting       ClassAPIResource   `json:"spellcasting"`
	StartingEquipment  ClassAPIResource   `json:"starting_equipment"`
	Subclasses         []ClassAPIResource `json:"subclasses"`
	Url                string             `json:"url"`
}

func GetClass(index string) (Class, error) {
	var ret Class
	err := doRequestAndUnmarshal("classes/"+index, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetClassByUrl(url string) (Class, error) {
	var ret Class
	err := doRequestRawAndUnmarshal(url, ret)
	return ret, err
}

func SearchClassByName(query string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("classes/?name="+query, &ret)
	return ret, err
}

// There should be more endpoints like this
func GetClassSubclasses(index string) (NamedAPIResourceList, error) {
	var ret NamedAPIResourceList
	err := doRequestAndUnmarshal("classes/"+index+"/subclasses", &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
