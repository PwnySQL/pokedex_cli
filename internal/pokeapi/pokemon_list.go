package pokeapi

import (
	"encoding/json"
	"errors"
)

func (c *Client) GetPokemonList(locationName *string) (RespSingleLocation, error) {
	var decodedRes RespSingleLocation
	if locationName == nil {
		return decodedRes, errors.New("No location name given. Cannot get pokemon list")
	}

	singleLocationUrl := baseURL + "/" + "location-area" + "/" + *locationName

	cachedRes, err := c.doCachedPokeapiRequest(&singleLocationUrl)
	if err != nil {
		return decodedRes, err
	}

	if err := json.Unmarshal(cachedRes, &decodedRes); err != nil {
		return decodedRes, err
	}
	return decodedRes, nil
}
