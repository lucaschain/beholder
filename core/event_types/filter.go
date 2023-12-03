package event_types

func Filter(event_type EventType, allowed_types []EventType) bool {
	for _, allowed_type := range allowed_types {
		if event_type == allowed_type {
			return true
		}
	}
	return false
}
