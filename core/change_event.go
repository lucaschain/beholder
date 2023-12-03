package core

import "github.com/lucaschain/beholder/core/event_types"

type ChangeEvent struct {
	Type     event_types.EventType
	FileName string
}

type ChangeCallback func(*ChangeEvent, *error) *error
