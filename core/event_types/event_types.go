package event_types

type EventType string

var (
	Create EventType = "CREATE"
	Write  EventType = "WRITE"
	Remove EventType = "REMOVE"
	Rename EventType = "RENAME"
	Chmod  EventType = "CHMOD"
)

var EventTypes = []EventType{Create, Write, Remove, Rename, Chmod}

func (e EventType) String() string {
	return string(e)
}

func FromString(str string) EventType {
	for _, event := range EventTypes {
		if event.String() == str {
			return event
		}
	}
	return ""
}
