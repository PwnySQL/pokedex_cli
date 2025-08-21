package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) doPokeapiRequest(url *string) (*http.Response, error) {
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func (c *Client) GetLocationList(url *string) (RespLocation, error) {
	locationUrl := baseURL + "/" + "location-area"
	if url != nil {
		locationUrl = *url
	}
	res, err := c.doPokeapiRequest(&locationUrl)
	if err != nil {
		return RespLocation{}, err
	}
	defer res.Body.Close()
	var decodedRes RespLocation
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedRes); err != nil {
		return RespLocation{}, err
	}
	return decodedRes, nil
}
