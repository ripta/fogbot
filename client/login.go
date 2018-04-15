package client

import (
	"encoding/json"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func (c *Client) Login(email, password string) error {
	rsp, err := c.Request("logon", Request{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return err
	}

	var login LoginResponse
	if err := json.Unmarshal(rsp, &login); err != nil {
		return err
	}
	return nil
}
