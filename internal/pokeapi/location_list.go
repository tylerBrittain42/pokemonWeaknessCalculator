package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationSet, error) {
	url := baseURL + "/location-area?limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	fmt.Println("url")
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationSet{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationSet{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationSet{}, err
	}

	locationSetResp := LocationSet{}
	err = json.Unmarshal(dat, &locationSetResp)

	if err != nil {
		return LocationSet{}, err
	}
	return locationSetResp, nil
}
