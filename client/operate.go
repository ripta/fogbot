package client

func (c *Client) Assign(caseID, assignToPersonID int) error {
	_, err := c.Request("assign", Request{
		"ixBug":              caseID,
		"ixPersonAssignedTo": assignToPersonID,
	})
	return err
}

func (c *Client) Close(caseID int) error {
	_, err := c.Request("close", Request{
		"ixBug": caseID,
	})
	return err
}

func (c *Client) Reactivate(caseID int) error {
	_, err := c.Request("reactivate", Request{
		"ixBug": caseID,
	})
	return err
}

func (c *Client) Reopen(caseID int) error {
	_, err := c.Request("reopen", Request{
		"ixBug": caseID,
	})
	return err
}

func (c *Client) Resolve(caseID int) error {
	_, err := c.Request("resolve", Request{
		"ixBug": caseID,
	})
	return err
}
