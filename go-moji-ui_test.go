package gomojiui

import (
	"fmt"
	"testing"
)

func TestContainsEmoji(t *testing.T) {
	// Test with emoji
	if !ContainsEmoji("Hello, 😊") {
		t.Error("Expected ContainsEmoji to return true for 'Hello, 😊'")
	}

	// Test without emoji
	if ContainsEmoji("Hello, World!") {
		t.Error("Expected ContainsEmoji to return false for 'Hello, World!'")
	}

	// Test with multiple emojis
	if !ContainsEmoji("Hello, 😊👍") {
		t.Error("Expected containsEmoji to return true for 'Hello, 😊👍'")
	}
}

func TestRemoveEmoji(t *testing.T) {
	// Test with emoji
	if RemoveEmoji("Hello, 😊") != "Hello, " {
		t.Error("Expected RemoveEmoji to return 'Hello, ' for 'Hello, 😊'")
	}

	// Test without emoji
	if RemoveEmoji("Hello, World!") != "Hello, World!" {
		t.Error("Expected RemoveEmoji to return 'Hello, World!' for 'Hello, World!'")
	}

	// Test with multiple emojis
	if RemoveEmoji("Hello, 😊👍") != "Hello, " {
		t.Error("Expected RemoveEmoji to return 'Hello, ' for 'Hello, 😊👍'")
	}
}

func TestCountEmojiRunes(t *testing.T) {
	// Test with basic emoji.
	if CountEmojiRunes("Hello, 😊")["😊"] != 1 {
		t.Error("Expected CountEmojiRunes to return 1 for '😊'")
	}

	// Test with grapheme cluster emoji.
	if CountEmojiRunes("Hello, 👨‍👩‍👧‍👦")["👨‍👩‍👧‍👦"] != 7 {
		fmt.Println(CountEmojiRunes("Hello, 👨‍👩‍👧‍👦"))
		t.Error("Expected CountEmojiRunes to return 7 for '👨‍👩‍👧‍👦'")
	}

	// Test with multiple emojis.
	if CountEmojiRunes("Hello, 😊👍")["😊"] != 1 || CountEmojiRunes("Hello, 😊👍")["👍"] != 1 {
		t.Error("Expected CountEmojiRunes to return 1 for '😊' and '👍'")
	}

	// Test with 2 rune emoji.
	if CountEmojiRunes("Hello, ❤️")["❤️"] != 2 {
		t.Error("Expected CountEmojiRunes to return 2 for '❤️'")
	}
}

func TestRemoveEmojiSize(t *testing.T) {
	// Emoji of rune size 1 should be removed, when 1 is provided.
	if FilterEmojisBySize("Hello, 😊", 1) != "Hello, " {
		t.Error("Expected RemoveEmojiSize to return 'Hello, ' for 'Hello, 😊'")
	}

	// Emoji of rune size 1 should be persisted, when 2 is provided.
	if FilterEmojisBySize("Hello, 😊", 2) != "Hello, 😊" {
		t.Error("Expected RemoveEmojiSize to return 'Hello, 😊' for 'Hello, 😊'")
	}

	// Functions with multiple emojis.
	if FilterEmojisBySize("Hello, 😊👍", 1) != "Hello, " {
		t.Error("Expected RemoveEmojiSize to return 'Hello, ' for 'Hello, 😊👍'")
	}

	// Functions with multiple emojis, where some emojis are persisted, some aren't
	if FilterEmojisBySize("Hello, 😊❤️", 2) != "Hello, 😊" {
		t.Error("Expected RemoveEmojiSize to return 'Hello, 😊' for 'Hello, 😊❤️'")
	}

	// Functions with 2 rune emoji.
	if FilterEmojisBySize("Hello, ❤️", 1) != "Hello, " {
		t.Error("Expected RemoveEmojiSize to return 'Hello, ' for 'Hello, ❤️'")
	}
}
