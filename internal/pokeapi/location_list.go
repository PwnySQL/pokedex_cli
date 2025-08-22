package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocationList(url *string) (RespLocation, error) {
	locationUrl := baseURL + "/" + "location-area"
	if url != nil {
		locationUrl = *url
	}
	var decodedRes RespLocation
	cachedRes, err := c.doCachedPokeapiRequest(&locationUrl)
	if err != nil {
		return decodedRes, err
	}

	if err := json.Unmarshal(cachedRes, &decodedRes); err != nil {
		return decodedRes, err
	}
	return decodedRes, nil
}
