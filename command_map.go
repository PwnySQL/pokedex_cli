package main

import (
	"fmt"
	"errors"
)

func commandMap(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.GetLocationList(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	cfg.nextLocationsUrl = locationResp.Next
	cfg.prevLocationsUrl = locationResp.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsUrl == nil {
		return errors.New("you're on the first page")
	}
	locationResp, err := cfg.pokeapiClient.GetLocationList(cfg.prevLocationsUrl)
	if err != nil {
		return err
	}
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	cfg.nextLocationsUrl = locationResp.Next
	cfg.prevLocationsUrl = locationResp.Previous

	return nil
}
