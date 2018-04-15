package client

import "encoding/json"

// Checkin represents metadata relating to a checkin.
type Checkin struct {
	BugID      int `json:"ixBug"`
	Repository int `json:"ixRepository,omitempty"`

	File         string `json:"sFile,omitempty"`
	PrevRevision string `json:"sPrev,omitempty"`
	NewRevision  string `json:"sNew,omitempty"`
}

type CheckinList []*Checkin

// ListCheckinsResponse represents a tag list.
type ListCheckinsResponse struct {
	Checkins CheckinList `json:"checkins"`
}

// ListCheckins lists tags.
func (c *Client) ListCheckins(caseID int) (CheckinList, error) {
	rsp, err := c.Request("listCheckins", Request{
		"ixBug": caseID,
	})
	if err != nil {
		return nil, err
	}

	// Decode response
	var lcr ListCheckinsResponse
	if err := json.Unmarshal(rsp, &lcr); err != nil {
		return nil, err
	}
	return lcr.Checkins, nil
}

func (c *Client) NewCheckin(ci *Checkin) error {
	req, err := BuildRequest(ci)
	if err != nil {
		return err
	}
	if _, err := c.Request("newCheckin", req); err != nil {
		return err
	}
	return nil
}

func (cl CheckinList) ByRevision() map[string]CheckinList {
	byRev := make(map[string]CheckinList)
	for _, ci := range cl {
		byRev[ci.NewRevision] = append(byRev[ci.NewRevision], ci)
	}
	return byRev
}
