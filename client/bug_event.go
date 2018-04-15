package client

import (
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
)

type BugEventID int

const (
	BugEventOpened            BugEventID = 1
	BugEventEdited                       = 2
	BugEventAssigned                     = 3
	BugEventReactivated                  = 4
	BugEventReopened                     = 5
	BugEventClosed                       = 6
	BugEventMoved                        = 7
	BugEventUnknown                      = 8
	BugEventReplied                      = 9
	BugEventForwarded                    = 10
	BugEventReceived                     = 11
	BugEventSorted                       = 12
	BugEventNotSorted                    = 13
	BugEventResolved                     = 14
	BugEventEmailed                      = 15
	BugEventReleaseNoted                 = 16
	BugEventDeletedAttachment            = 17
)

type BugEvent struct {
	ID        int        `json:"ixBugEvent"`
	EventID   BugEventID `json:"evt"`
	EventTime *Time      `json:"dt"`
	Verb      string     `json:"sVerb"`

	Change      string `json:"sChanges"`
	Description string `json:"evtDescription"`

	EventPersonID      int `json:"ixPerson"`
	AssignedToPersonID int `json:"ixPersonAssignedTo"`

	Text   string `json:"s"`
	HTML   string `json:"sHTML"`
	Format string `json:"sFormat"`

	CreatedFromEmail    bool `json:"fEmail"`
	CreatedFromExternal bool `json:"fExternal"`
	CreatedAsHTML       bool `json:"fHTML"`

	MessageFrom    string `json:"sFrom"`
	MessageTo      string `json:"sTo"`
	MessageCC      string `json:"sCC"`
	MessageBCC     string `json:"sBCC"`
	MessageReplyTo string `json:"sReplyTo"`
	MessageSubject string `json:"sSubject"`
	MessageDate    string `json:"sDate"`
	MessageText    string `json:"sBodyText"`
	MessageHTML    string `json:"sBodyHTML"`
}

func (e *BugEvent) IsHTML() bool {
	return e.Format == "html"
}

func (e *BugEvent) WrappedText() []string {
	lines := make([]string, 0)
	for _, line := range strings.Split(e.Text, "\n") {
		seg := wordwrap.WrapString(line, 80)
		for _, segline := range strings.Split(seg, "\n") {
			lines = append(lines, segline)
		}
	}
	return lines
}
