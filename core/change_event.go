package core

type ChangeEvent struct {
	Type     string
	FileName string
}

type ChangeCallback func(*ChangeEvent, *error)
