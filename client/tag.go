package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Tag represents a tag.
type Tag struct {
	ID       int    `json:"ixTag"`
	Name     string `json:"sTag"`
	UseCount int    `json:"cTagUses"`
}

type TagList []*Tag

// ListTagsResponse represents a tag list.
type ListTagsResponse struct {
	Tags []*Tag `json:"tags"`
}

func (c *Client) LookupStatus(code StatusCode) (*StatusDescription, bool) {
	statuses, err := c.ListStatuses()
	if err != nil {
		return nil, false
	}
	for _, status := range statuses {
		if status.Code == code {
			return status, true
		}
	}
	return nil, false
}

func (c *Client) LookupTag(name string) (*Tag, bool) {
	tags, err := c.ListTags()
	if err != nil {
		return nil, false
	}
	for _, tag := range tags {
		if tag.Name == name {
			return tag, true
		}
	}
	return nil, false
}

// ListTags lists tags.
func (c *Client) ListTags() (TagList, error) {
	if c.tagCache == nil {
		rsp, err := c.Request("listTags", Request{})
		if err != nil {
			return nil, err
		}

		// Decode response
		var ltr ListTagsResponse
		if err := json.Unmarshal(rsp, &ltr); err != nil {
			return nil, err
		}
		c.tagCache = ltr.Tags
	}
	return c.tagCache, nil
}

func (t *Tag) String() string {
	return fmt.Sprintf("%s (%d)", t.Name, t.ID)
}

func (tl TagList) String() string {
	n := []string{}
	for _, t := range tl {
		n = append(n, t.String())
	}
	return strings.Join(n, ", ")
}
