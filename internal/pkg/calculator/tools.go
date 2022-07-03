package calculator

func stringSliceContain(item string, list []string) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
}
