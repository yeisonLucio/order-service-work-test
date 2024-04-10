package helpers

// StringContains permite validar si un slice contiene un string
func StringContains(options []string, value string) bool {
	found := false
	for _, v := range options {
		if v == value {
			found = true
			break
		}
	}

	return found
}
