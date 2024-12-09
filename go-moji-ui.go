package gomojiui

import (
	"regexp"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

// Regex pattern that identifies emojis, including those made of multiple runes.
var emojiRegex = regexp.MustCompile(`(?:\p{So}|\p{Sk}|\x{1F600}-\x{1F64F}|\x{1F300}-\x{1F5FF}|\x{1F680}-\x{1F6FF}|\x{2600}-\x{26FF}|\x{2700}-\x{27BF})`)

// Checks if a string contains any emoji.
func ContainsEmoji(text string) bool {
	return emojiRegex.MatchString(text)
}

// Removes all emojis from a string.
func RemoveEmoji(text string) string {
	return emojiRegex.ReplaceAllString(text, "")
}

// Removes emojis that are larger or equal to the number of runes provided.
func FilterEmojisBySize(text string, runes int) string {
	count := CountEmojiRunes(text)

	var result string
	graphemes := uniseg.NewGraphemes(text)

	for graphemes.Next() {
		cluster := graphemes.Str()

		if emojiRegex.MatchString(cluster) {
			if size, exists := count[cluster]; exists && size >= runes {
				continue
			}
		}

		result += cluster
	}

	return result
}

// CountEmojiRunes counts the number of runes for each emoji in a string. Returns a map of emoji to rune count.
func CountEmojiRunes(text string) map[string]int {
	results := make(map[string]int)
	graphemes := uniseg.NewGraphemes(text)

	for graphemes.Next() {
		cluster := graphemes.Str()
		if emojiRegex.MatchString(cluster) {
			results[cluster] = utf8.RuneCountInString(cluster)
		}
	}
	return results
}
