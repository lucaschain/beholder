package core

import (
	"strings"
)

func replaceTokens(text string, event *ChangeEvent) string {
	text = strings.ReplaceAll(text, "{type}", event.Type.String())
	return strings.ReplaceAll(text, "{file}", event.FileName)
}

func CommandTokens(command []string, event *ChangeEvent) []string {
	var replacedCommand []string

	for _, token := range command {
		replacedCommand = append(replacedCommand, replaceTokens(token, event))
	}

	return replacedCommand
}
