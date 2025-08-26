package pokeapi

import (
	"encoding/json"
	"errors"
)

func (c *Client) GetPokemon(pokemonName *string) (RespPokemon, error) {
	var decodedRes RespPokemon
	if pokemonName == nil {
		return decodedRes, errors.New("No pokemon name given. Cannot get pokemon")
	}

	pokemonUrl := baseURL + "/" + "pokemon" + "/" + *pokemonName

	cachedRes, err := c.doCachedPokeapiRequest(&pokemonUrl)
	if err != nil {
		return decodedRes, err
	}

	if err := json.Unmarshal(cachedRes, &decodedRes); err != nil {
		return decodedRes, err
	}
	return decodedRes, nil
}
