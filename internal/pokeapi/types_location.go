package pokeapi

type RespLocation struct {
	Count    int     `json:count`
	Next     *string `json:next`
	Previous *string `json:previous`
	Results  []struct {
		Name string `json:name`
		Url  string `json:url`
	} `json:results`
}

type RespLocation struct {
	Count    int        `json:count`
	Next     *string    `json:next`
	Previous *string    `json:previous`
	Results  []Location `json:results`
}
