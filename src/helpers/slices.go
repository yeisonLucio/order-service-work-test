package helpers

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
