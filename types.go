package go5e

type NamedAPIResource struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

type NamedAPIResourceList struct {
	Count   int                `json:"count"`
	Results []NamedAPIResource `json:"results"`
}

type ClassAPIResource struct {
	Index string `json:"index"`
	Class string `json:"class"`
	Url   string `json:"url"`
}

type NamedChoice struct {
	Choose int
	Type   string
	From   []NamedAPIResource
}

type ClassChoice struct {
	Choose int
	Type   string
	From   []ClassAPIResource
}

type Level struct {
	Id                  string             `json:"_id"`
	Level               int                `json:"level"`
	AbilityScoreBonuses int                `json:"ability_score_bonuses"`
	ProfBonus           int                `json:"prof_bonus"`
	FeatureChoices      []NamedAPIResource `json:"feature_choices"`
	Features            []NamedAPIResource `json:"features"`
	Spellcasting        LevelSpellcasting  `json:"spellcasting"`
	ClassSpecific       string             `json:"class_specific"` // TODO generic interface ?
	Index               int                `json:"index"`
	Class               NamedAPIResource   `json:"class"`
	Url                 string             `json:"url"`
}

// ???
type LevelSpellcasting struct {
	CantripsKnown int `json:"cantrips_known"`
	SpellsKnown   int `json:"spells_known"`
	SpellSlots1   int `json:"spell_slots_level_1"`
	SpellSlots2   int `json:"spell_slots_level_2"`
	SpellSlots3   int `json:"spell_slots_level_3"`
	SpellSlots4   int `json:"spell_slots_level_4"`
	SpellSlots5   int `json:"spell_slots_level_5"`
	SpellSlots6   int `json:"spell_slots_level_6"`
	SpellSlots7   int `json:"spell_slots_level_7"`
	SpellSlots8   int `json:"spell_slots_level_8"`
	SpellSlots9   int `json:"spell_slots_level_9"`
}

type AbilityScore struct {
	Id       string             `json:"_id"`
	Index    string             `json:"index"`
	Name     string             `json:"name"`
	FullName string             `json:"full_name"`
	Desc     []string           `json:"desc"`
	Skills   []NamedAPIResource `json:"skills"`
	Url      string             `json:"url"`
}

type Proficiency struct {
	Id      string             `json:"_id"`
	Index   string             `json:"index"`
	Type    string             `json:"type"`
	Name    string             `json:"name"`
	Classes []NamedAPIResource `json:"classes"`
	Races   []NamedAPIResource `json:"races"`
	Url     string             `json:"url"`
}

type Language struct {
	Id              string   `json:"_id"`
	Index           string   `json:"index"`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	TypicalSpeakers []string `json:"typical_speakers"`
	Script          string   `json:"script"`
	Url             string   `json:"url"`
}

type Class struct {
	Id                 string             `json:"_id"`
	Index              string             `json:"index"`
	Name               string             `json:"name"`
	HitDie             int                `json:"hit_die"`
	ProficiencyChoices []NamedChoice      `json:"proficiency_choices"`
	Proficiencies      []NamedAPIResource `json:"proficiencies"`
	SavingThrows       []NamedAPIResource `json:"saving_throws"`
	StartingEquipment  ClassAPIResource   `json:"starting_equipment"`
	ClassLevels        ClassAPIResource   `json:"class_levels"`
	Subclasses         []ClassAPIResource `json:"subclasses"`
	Spellcasting       ClassAPIResource   `json:"spellcasting"`
	Url                string             `json:"url"`
}

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

type Cost struct {
	Quantity int
	Unit     string
}
