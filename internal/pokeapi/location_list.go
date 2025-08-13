package pokeapi

import (
	"encoding/json"
	"net/http"
)


func (c *Client) doPokeapiRequest(pageUrl *string) (*http.Response, error) {
	url := baseURL + "location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func (c *Client) GetLocationList(url *string) (RespLocation, error) {
	res, err := c.doPokeapiRequest(url)
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

