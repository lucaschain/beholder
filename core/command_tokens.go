package core

import (
	"strings"

	"github.com/fsnotify/fsnotify"
)

func replaceTokens(text string, event *fsnotify.Event) string {
	return strings.ReplaceAll(text, "{file}", event.Name)
}

func Replace(command []string, event *fsnotify.Event) []string {
	var replacedCommand []string

	for _, token := range command {
		replacedCommand = append(replacedCommand, replaceTokens(token, event))
	}

	return replacedCommand
}
