package dbutil

// Config store configuration for this util
type Config map[string]interface{}

// GetString return string value by given key
func (c Config) GetString(key string, fallback ...string) string {
	var ret string
	if len(fallback) > 0 {
		ret = fallback[0]
	}

	if value, exist := c[key]; exist {
		if v, ok := value.(string); ok {
			return v
		}
	}

	return ret
}

// GetInt return int value by given key
func (c Config) GetInt(key string, fallback ...int) int {
	var ret int
	if len(fallback) > 0 {
		ret = fallback[0]
	}

	if value, exist := c[key]; exist {
		if v, ok := value.(int); ok {
			return v
		}
	}

	return ret
}

// TODO: add support of other types
