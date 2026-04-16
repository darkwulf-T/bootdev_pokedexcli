package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type exploreResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) PokemonResponse(name string) (exploreResponse, error) {
	endpoint := baseUrl + "/location-area/" + name

	if data, ok := c.cache.Get(endpoint); ok {
		var response exploreResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return exploreResponse{}, fmt.Errorf("Error while unmarshalling: %w", err)
		}

		return response, nil
	}

	res, err := c.httpClient.Get(endpoint)
	if err != nil {
		return exploreResponse{}, fmt.Errorf("Error with the GET request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return exploreResponse{}, fmt.Errorf("Error reading the response: %w", err)
	}

	c.cache.Add(endpoint, data)

	var response exploreResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return exploreResponse{}, fmt.Errorf("Error while unmarshalling: %w", err)
	}

	return response, nil
}
