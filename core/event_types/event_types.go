package event_types

type EventType string

var (
	Create EventType = "CREATE"
	Write  EventType = "WRITE"
	Remove EventType = "REMOVE"
	Rename EventType = "RENAME"
)

func (e EventType) IsValid() bool {
	switch e {
	case Create, Write, Remove, Rename:
		return true
	}
	return false
}

func (e EventType) String() string {
	return string(e)
}
