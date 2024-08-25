package handler

import (
	"net/url"
	"strconv"
)

func ParseString(item string, m url.Values, defaultValue string) string {
	if v := m.Get(item); len(v) != 0 {
		return v
	}
	return defaultValue
}

func ParseInt(item string, m url.Values, defaultValue int) int {
	if v := m.Get(item); len(v) != 0 {
		if i, err := strconv.Atoi(v); err != nil {
			return i
		}
	}

	return defaultValue
}
