package thecatapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const base = "https://api.thecatapi.com/v1"

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	c := &Client{}
	c.httpClient = &http.Client{}
	c.apiKey = apiKey

	return c
}

//create request
func (c *Client) makeRequest(endpoint string, method string) (*http.Request, error) {

	//Debug endpoint
	//fmt.Println(endpoint)

	url := base + endpoint
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	return req, err
}

//running the request
func (c *Client) runRequest(req *http.Request, results interface{}) error {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	//Debug json
	//body, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))

	err = json.NewDecoder(response.Body).Decode(results)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
