package client

import "encoding/json"

type Status string
type StatusCode int

type StatusDescription struct {
	Name  Status     `json:"sStatus"`
	Code  StatusCode `json:"ixStatus"`
	Order int        `json:"iOrder"`

	Category    int  `json:"ixCategory"`
	WorkDone    bool `json:"fWorkDone"`
	Resolved    bool `json:"fResolved"`
	Duplicate   bool `json:"fDuplicate"`
	Deleted     bool `json:"fDeleted"`
	Reactivated bool `json:"fReactivate"`
}

type ListStatusesResponse struct {
	Statuses []*StatusDescription `json:"statuses"`
}

func (c *Client) ListStatuses() ([]*StatusDescription, error) {
	if len(c.statusCache) == 0 {
		rsp, err := c.Request("listStatuses", Request{})
		if err != nil {
			return nil, err
		}

		var list ListStatusesResponse
		if err := json.Unmarshal(rsp, &list); err != nil {
			return nil, err
		}
		c.statusCache = list.Statuses
	}
	return c.statusCache, nil
}
