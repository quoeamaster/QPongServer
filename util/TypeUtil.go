package util

import "strings"

/**
 *  helper method to check if the given string is empty or not
 */
func IsStringEmpty(val string) bool {
	return strings.Compare("", strings.TrimSpace(val))==0
}
