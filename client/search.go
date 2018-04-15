package client

import (
	"encoding/json"
	"strconv"
	"strings"
)

// SearchResponse represents a response of a query.
type SearchResponse struct {
	Count     int     `json:"count"`
	TotalHits int     `json:"totalHits"`
	Cases     []*Case `json:"cases"`

	Client *Client `json:"-"`
	Query  string  `json:"-"`
}

// Search searches FogBugz for cases matching a specific query.
func (c *Client) Search(query string, limit int) (*SearchResponse, error) {
	rsp, err := c.Request("search", Request{
		"q":    query,
		"max":  limit,
		"cols": AllColumnNames(),
	})
	if err != nil {
		return nil, err
	}

	// Decode response
	var sr SearchResponse
	if err := json.Unmarshal(rsp, &sr); err != nil {
		return nil, err
	}

	// Decorate result
	sr.Client = c
	sr.Query = query
	for _, cs := range sr.Cases {
		cs.Client = c
	}
	return &sr, nil
}

func (c *Client) SearchByID(id int) (*Case, error) {
	sr, err := c.Search(strconv.Itoa(id), 1)
	if err != nil {
		return nil, err
	}
	if len(sr.Cases) >= 1 {
		return sr.Cases[0], nil
	}
	return nil, nil
}

func (c *Client) SearchByIDs(ids []int) ([]*Case, error) {
	idStrings := make([]string, len(ids))
	for _, id := range ids {
		idStrings = append(idStrings, strconv.Itoa(id))
	}
	sr, err := c.Search(strings.Join(idStrings, ","), len(ids))
	if err != nil {
		return nil, err
	}
	if len(sr.Cases) >= 1 {
		return sr.Cases, nil
	}
	return nil, nil
}
