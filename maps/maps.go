package maps

// Keys returns the keys of the given map
func Keys(input map[string]interface{}) []string {
	keys := make([]string, 0, len(input))
	for k := range input {
			keys = append(keys, k)
	}
	return keys
}