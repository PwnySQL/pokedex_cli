package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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
	var decodedRes RespLocation
	fmt.Printf("Request '%s' from cache\n", locationUrl)
	cachedRes, isCached := c.cache.Get(locationUrl)
	if !isCached {
		fmt.Println("Populate cache")
		res, err := c.doPokeapiRequest(&locationUrl)
		if err != nil {
			return decodedRes, err
		}
		defer res.Body.Close()
		cachedRes, err = io.ReadAll(res.Body)
		if err != nil {
			return decodedRes, err
		}
		c.cache.Add(locationUrl, cachedRes)
	}
	if err := json.Unmarshal(cachedRes, &decodedRes); err != nil {
		return decodedRes, err
	}
	return decodedRes, nil
}
