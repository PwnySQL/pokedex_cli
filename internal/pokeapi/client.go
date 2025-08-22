package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PwnySQL/pokedex_cli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Minute),
	}
}

func (c *Client) doPokeapiRequest(url *string) (*http.Response, error) {
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func (c *Client) doCachedPokeapiRequest(url *string) ([]byte, error) {
	fmt.Printf("Request '%s' from cache\n", *url)
	cachedRes, isCached := c.cache.Get(*url)
	if !isCached {
		fmt.Println("Populate cache")
		res, err := c.doPokeapiRequest(url)
		if err != nil {
			return cachedRes, err
		}
		defer res.Body.Close()
		cachedRes, err = io.ReadAll(res.Body)
		if err != nil {
			return cachedRes, err
		}
		c.cache.Add(*url, cachedRes)
	}
	return cachedRes, nil
}
