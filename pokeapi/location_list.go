package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) RequestFunction(url *string) (ResponseApi, error) {
	endpoint := baseUrl + "/location-area"
	if url != nil {
		endpoint = *url
	}

	res, err := c.httpClient.Get(endpoint)
	if err != nil {
		return ResponseApi{}, fmt.Errorf("Error with the GET request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseApi{}, fmt.Errorf("Error reading the response: %w", err)
	}

	var response ResponseApi
	if err := json.Unmarshal(data, &response); err != nil {
		return ResponseApi{}, fmt.Errorf("Error while unmarshalling: %w", err)
	}
	
	return response, nil
}