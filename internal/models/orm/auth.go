package orm

type Auth struct {
	AccToken string `json:"accToken"`
	RefToken string `json:"refToken"`
}
