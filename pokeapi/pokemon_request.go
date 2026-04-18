package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) CatchResponse(name string) (Pokemon, error) {
	endpoint := baseUrl + "/pokemon/" + name + "/"

	res, err := c.httpClient.Get(endpoint)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error with the GET request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error reading the response: %w", err)
	}

	var response Pokemon
	if err := json.Unmarshal(data, &response); err != nil {
		return Pokemon{}, fmt.Errorf("Error while unmarshalling: %w", err)
	}

	return response, nil
}
