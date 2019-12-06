package util

func IsEmptyMap(v map[string][]string) bool {
	return len(v) == 0
}

func IsEmptyStringMap(v map[string]string) bool {
	return len(v) == 0
}
