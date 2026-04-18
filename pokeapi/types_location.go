package pokeapi

type ResultApi struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ResponseApi struct {
	Count    int         `json:"count"`
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []ResultApi `json:"results"`
}
