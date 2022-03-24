package check

// Reads a raw string, and emits header key and value. set success to true if able to do, false otherwise.
func GetHeaderKeyValue(headerkeyValue string) (success bool, key string, value string) {
	success = false
	headerkeyValueLen := len(headerkeyValue)
	if 0 == headerkeyValueLen {
		return
	}
	for pos, char := range headerkeyValue {
		if char == ':' && ((headerkeyValueLen > (pos + 1)) && headerkeyValue[pos+1] == ' ') {
			till := pos + 3
			if headerkeyValueLen > (till) {
				key = headerkeyValue[:pos]
				value = headerkeyValue[till-1:]
				break
			}
		}
	}
	if key != "" && value != "" {
		success = true
	}
	return
}
