package go5e

// Not all fields are guaranteed to be populated;
// see documentation for what to expect to actually be
// in this struct when it is instantiated

// Perhaps these optional fields would be better suited
// as separate types
type NamedAPIResource struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Value int `json:"value"`
	Bonus int `json:"bonus"`
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

type Cost struct {
	Quantity int `json:"quantity"`
	Unit string `json:"unit"`
}

// Might be missing some things
type Speed struct {
	Walk string `json:"walk"`
	Swim string `json:"swim"`
	Fly string `json:"fly"`
}

// Might be missing some things
type Senses struct {
	Darkvision string `json:"darkvision"`
	Blindsight string `json:"blindsight"`
	PassivePerception int `json:"passive_perception"`
}

type Ability struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	DifficultyCheck DifficultyCheck `json:"dc"`
	AttackBonus int `json:"attack_bonus"`
	Damage Damage `json:"damage"`
	Usage Usage `json:"usage"`
}

type Usage struct {
	Type string `json:"type"`
	Times int `json:"times"`
}

type DifficultyCheck struct {
		Type NamedAPIResource `json:"dc_type"`
		Value int `json:"dc_value"`
		SuccessType string `json:"success_type"`
}

type Damage struct {
	Type NamedAPIResource `json:"damage_type"`
	Dice string `json:"damage_dice"`
	Bonus int `json:"damage_bonus"`
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
