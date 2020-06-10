package thecatapi

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

//Fetching cat images
func (c *Client) GetImageSearch(options ImageSearchOptions) (*ImageResults, error) {

	//Default settings for GET parameters
	if options == (ImageSearchOptions{}) {
		options = ImageSearchOptions{"med", "jpg,gif,png", "RANDOM", 1, 1}
	}

	//GET params
	params, _ := query.Values(options)

	req, err := c.makeRequest("/images/search?"+params.Encode(), http.MethodGet)
	if err != nil {
		return nil, err
	}

	v := &ImageResults{}

	err = c.runRequest(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

//Fetching categories
func (c *Client) GetCategories() (*CategoryResults, error) {

	req, err := c.makeRequest("/categories", http.MethodGet)
	if err != nil {
		return nil, err
	}

	v := &CategoryResults{}

	err = c.runRequest(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
