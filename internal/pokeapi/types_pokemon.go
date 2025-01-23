package pokeapi

type Stat struct {
	Name string `json:"name"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type Type struct {
	Name string `json:"name"`
}

type Types struct {
	Type Type `json:"type"`
}

type RespPokemon struct {
	Name   string  `json:"name"`
	BaseXP int     `json:"base_experience"`
	Height int     `json:"height"`
	Weight int     `json:"weight"`
	Stats  []Stats `json:"stats"`
	Types  []Types `json:"types"`
}
