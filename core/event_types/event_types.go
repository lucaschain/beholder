package event_types

type EventType = string

var (
	Create EventType = "CREATE"
	Write            = "WRITE"
	Remove           = "REMOVE"
	Rename           = "RENAME"
)
