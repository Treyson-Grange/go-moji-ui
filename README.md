# go-moji-ui

[![Report Card](https://goreportcard.com/badge/github.com/Treyson-Grange/go-moji-ui?style=flat-square&label=Go%20Report)](https://goreportcard.com/report/github.com/treyson-grange/go-moji-ui)

A Golang library that provides functions to detect, remove, and count emojis in a string, including those composed of multiple runes.

Emojis can enhance terminal output, but they can also introduce rendering issues. Some terminals render emojis with incorrect spacing, causing misalignment and disrupting text layout. This library is designed to address these issues by providing utilities to detect, remove, or process emojis in text effectively.

This was created for [JukeTUI](https://github.com/Treyson-Grange/JukeTUI), but can be useful elsewhere.

## Library Features

- Emoji Detection: Identify the presence of emojis in a string via regex
- Emoji Removal: Strip emojis from a given string via regex
- Emoji Rune Count: Issues arise when emojis are 2+ runes. If we can account for the runes in each emoji, we can account for these spacing issues.
- Emoji Size Based Removal: Remove all emojis that have n or more runes based on a requested size (n).

## Installation

Install the library on your terminal:

```bash
go get github.com/Treyson-Grange/go-moji-ui
```

## Usage

```go
import (
    "fmt"
    moji "github.com/Treyson-Grange/go-moji-ui"
)

func main() {
    // Basic emoji detection
    text := "Hello, World! 😊"
    fmt.Println(moji.ContainsEmoji(text)) // Prints: true

    // Remove all emojis
    text = "No Mojis! 🤓🤓🤓🤓🤓🤓🤓🤓🤓"
    fmt.Println(moji.RemoveEmoji(text)) // Prints: "No Mojis! "

    // Count runes in complex emojis (grapheme clusters)
    text = "Mixed emojis: 😊 👨‍👩‍👧‍👦 ❤️"
    runeCount := moji.CountEmojiRunes(text)
    fmt.Printf("Emoji rune counts: %v\n", runeCount)
    // Prints: Emoji rune counts: map[😊:1 👨‍👩‍👧‍👦:7 ❤️:2]

    for emoji, count := range runeCount {
        fmt.Printf("'%s' uses %d runes\n", emoji, count)
    }
    // Prints:
    // '😊' uses 1 rune
    // '👨‍👩‍👧‍👦' uses 7 runes (family: man, woman, girl, boy)
    // '❤️' uses 2 runes (heart + variation selector)

    // Filter out emojis with 2 or more runes
    text = "Keep simple: 😊 Remove complex: ❤️ and families: 👨‍👩‍👧‍�"
    fmt.Println(moji.FilterEmojisBySize(text, 2))
    // Prints: "Keep simple: 😊 Remove complex:  and families: "
}
```
