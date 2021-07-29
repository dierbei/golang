package xlt

type Header map[string][]string

// Add key value
// value is slice append value
func (h Header) Add(key, value string) {
	h[key] = append(h[key], value)
}

// Set key value
// value is slice set value
func (h Header) Set(key, value string) {
	h[key] = []string{value}
}

// Get value
// according to key
func (h Header) Get(key string) string {
	if value, ok := h[key]; ok && len(value) > 0 {
		return value[0]
	}
	return ""
}
