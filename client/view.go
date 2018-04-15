package client

// MarkRead marks the case as being viewed up to the latest bugEventID.
func (c *Client) MarkRead(caseID int) error {
	_, err := c.Request("search", Request{
		"ixBug": caseID,
	})
	return err
}

// MarkReadUntil marks the case as viewed up to a bugEventID.
func (c *Client) MarkReadUntil(caseID, bugEventID int) error {
	_, err := c.Request("search", Request{
		"ixBug":      caseID,
		"ixBugEvent": bugEventID,
	})
	return err
}

// MarkUnread marks the case as unread.
func (c *Client) MarkUnread(caseID int) error {
	return c.MarkReadUntil(caseID, 1)
}
