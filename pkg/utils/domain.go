package utils

import (
	"regexp"
)

// IsValidDomainName returns whether the input `name` is a valid domain name or not
func IsValidDomainName(name string) bool {
	v, _ := regexp.MatchString(
		"(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]", name)
	return bool(v)
}
