package utils

import "strings"

func SplitStringSlice(s string) []string {
	parts := strings.Split(s, ",")

	var result []string

	for _, part := range parts {
		trimmedPart := strings.TrimSpace(strings.Trim(part, " "))
		if trimmedPart != "" {
			result = append(result, trimmedPart)
		}
	}
	return result
}
