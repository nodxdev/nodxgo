package nodx

import "strings"

// EscapeHTML escapes the input string to prevent XSS attacks.
func EscapeHTML(input string) string {
	replacements := map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"\"": "&quot;",
		"'":  "&#39;",
	}

	var builder strings.Builder

	for _, char := range input {
		strChar := string(char)
		if replacement, exists := replacements[strChar]; exists {
			builder.WriteString(replacement)
		} else {
			builder.WriteString(strChar)
		}
	}

	return builder.String()
}
