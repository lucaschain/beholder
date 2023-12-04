package core

func ExtensionFilter(path string, extensions []string) bool {
	if len(extensions) == 0 {
		return true
	}

	for _, extension := range extensions {
		if extension == path[len(path)-len(extension):] {
			return true
		}
	}

	return false
}
