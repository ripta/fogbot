package client

import (
	"reflect"
	"strings"
)

type Case struct {
	ID       int   `json:"ixBug,omitempty" fogbot:"newOptional"`
	ParentID int   `json:"ixBugParent,omitempty"`
	ChildIDs []int `json:"ixBugChildren,omitempty"`

	BugEventID       int `json:"ixBugEvent,omitempty" fogbot:"notNew"`
	LatestBugEventID int `json:"ixBugEventLatestText,omitempty"`

	Open     bool     `json:"fOpen"`
	TagNames []string `json:"tags,omitempty"`

	Title         string `json:"sTitle"`
	OriginalTitle string `json:"sOriginalTitle,omitempty"`
	LatestSummary string `json:"sLatestTextSummary,omitempty"`

	ProjectID   int `json:"ixProject,omitempty"`
	AreaID      int `json:"ixArea,omitempty"`
	GroupID     int `json:"ixGroup,omitempty"`
	MilestoneID int `json:"ixFixFor,omitempty"`
	CategoryID  int `json:"ixCategory,omitempty"`
	PriorityID  int `json:"ixPriority,omitempty"`

	AssignedToPersonID   int `json:"ixPersonAssignedToID,omitempty"`
	OpenedByPersonID     int `json:"ixPersonOpenedBy,omitempty"`
	ClosedByPersonID     int `json:"ixPersonClosedBy,omitempty"`
	ResolvedByPersonID   int `json:"ixPersonResolvedBy,omitempty"`
	LastEditedByPersonID int `json:"ixPersonLastEditedBy,omitempty"`

	StatusCode    StatusCode `json:"ixStatus,omitempty"`
	DuplicateIDs  int        `json:"ixBugDuplicates,omitempty"`
	DuplicateOfID int        `json:"ixBugOriginal,omitempty"`

	DueOn            *Time `json:"dtDue,omitempty"`
	OriginalEstimate int   `json:"hrsOrigEst,omitempty"`
	CurrentEstimate  int   `json:"hrsCurrEst,omitempty"`
	ExtraEstimate    int   `json:"hrsElapsedExtra,omitempty"`
	ElapsedEstimate  int   `json:"hrsElapsed,omitempty"`
	StoryPoints      int   `json:"dblStoryPts,omitempty"`

	Version       string `json:"sVersion,omitempty"`  // custom field #1
	Computer      string `json:"sComputer,omitempty"` // custom field #2
	CustomerEmail string `json:"sCustomerEmail,omitempty"`
	MailboxID     int    `json:"ixMailbox,omitempty"`

	KanbanColumnID int `json:"ixKanbanColumn,omitempty"`

	ScoutDescription     string `json:"sScoutDescription,omitempty" fogbot:"newOnly"`
	ScoutMessage         string `json:"sScoutMessage,omitempty"`
	ScoutStopReporting   bool   `json:"fScoutStopReporting,omitempty"`
	ScoutOccurrenceCount int    `json:"c,omitempty"` // add 1 to number to get actual count

	EventText     string `json:"sEvent,omitempty"`
	RichEventText bool   `json:"fRichEvent,omitempty"`

	TicketURL string `json:"sTicket,omitempty"`

	Events []*BugEvent `json:"events,omitempty"`

	// not implemented
	UploadFiles     string `json:"-"`
	UploadFileCount int    `json:"-"`

	Operations []string `json:"operations,omitempty"`

	Client *Client `json:"-"`
}

func (c *Case) Assign(assignToPersonID int) error {
	_, err := c.Client.Request("assign", Request{
		"ixBug":              c.ID,
		"ixPersonAssignedTo": assignToPersonID,
	})
	return err
}

func (c *Case) Close() error {
	_, err := c.Client.Request("close", Request{
		"ixBug": c.ID,
	})
	return err
}

func (c *Case) LastEvent() *BugEvent {
	if len(c.Events) == 0 {
		return nil
	}
	return c.Events[len(c.Events)-1]
}

func (c *Case) LatestEvents(limit int) []*BugEvent {
	i := len(c.Events)
	if i <= limit {
		return c.Events
	}
	return c.Events[i-limit : i]
}

func (c *Case) ListCheckins() (CheckinList, error) {
	return c.Client.ListCheckins(c.ID)
}

// MarkRead marks the case as being viewed up to the latest bugEventID.
func (c *Case) MarkRead() error {
	return c.Client.MarkRead(c.ID)
}

// MarkReadUntil marks the case as viewed up to a bugEventID.
func (c *Case) MarkReadUntil(bugEventID int) error {
	return c.Client.MarkReadUntil(c.ID, bugEventID)
}

// MarkUnread marks the case as unread.
func (c *Case) MarkUnread() error {
	return c.Client.MarkUnread(c.ID)
}

func (c *Case) Reactivate() error {
	_, err := c.Client.Request("reactivate", Request{
		"ixBug": c.ID,
	})
	return err
}

func (c *Case) Reopen() error {
	_, err := c.Client.Request("reopen", Request{
		"ixBug": c.ID,
	})
	return err
}

func (c *Case) Resolve() error {
	_, err := c.Client.Request("resolve", Request{
		"ixBug": c.ID,
	})
	return err
}

func (c *Case) Status() *StatusDescription {
	if st, ok := c.Client.LookupStatus(c.StatusCode); ok {
		return st
	}
	return nil
}

func (c *Case) StatusName() Status {
	if st, ok := c.Client.LookupStatus(c.StatusCode); ok {
		return st.Name
	}
	return ""
}

func (c *Case) Tags() TagList {
	tags := make(TagList, 0)
	for _, name := range c.TagNames {
		if tag, ok := c.Client.LookupTag(name); ok {
			tags = append(tags, tag)
		}
	}
	return tags
}

func AllColumnNames() []string {
	t := reflect.TypeOf(Case{})
	cols := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		if jt, ok := f.Tag.Lookup("json"); ok {
			if jt == "" {
				continue // unhandled: when field name matches JSON name
			}
			if idx := strings.Index(jt, ","); idx != -1 {
				cols = append(cols, jt[:idx])
			} else {
				cols = append(cols, jt)
			}
		}
	}
	return cols
}
