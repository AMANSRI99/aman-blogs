// pkg/utils/content/content.go
package content

import (
	"strings"
	"unicode/utf8"

	"github.com/gosimple/slug"
)

func generateSlug(title string) string {
	return slug.Make(title)
}

func calculateReadTime(content string) int {
	// Average reading speed: 200 words per minute
	wordCount := len(strings.Fields(content))
	readTime := wordCount / 200
	if readTime < 1 {
		return 1
	}
	return readTime
}

func GenerateDescription(content string, maxLength int) string {
	if utf8.RuneCountInString(content) <= maxLength {
		return content
	}

	runes := []rune(content)
	truncated := string(runes[:maxLength])
	lastSpace := strings.LastIndex(truncated, " ")

	if lastSpace != -1 {
		truncated = truncated[:lastSpace]
	}

	return truncated + "..."
}
