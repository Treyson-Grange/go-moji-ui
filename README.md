# go-moji-ui

[![Report Card](https://goreportcard.com/badge/github.com/Treyson-Grange/go-moji-ui?style=flat-square&label=Go%20Report)](https://goreportcard.com/report/github.com/treyson-grange/JukeTUI)

A Golang library that provides functions to detect, remove, and count emojis in a string, including those composed of multiple runes.

Emojis can enhance terminal output, but they can also introduce rendering issues. Some terminals render emojis with incorrect spacing, causing misalignment and disrupting text layout. This library is designed to address these issues by providing utilities to detect, remove, or process emojis in text effectively.

This was created for [JukeTUI](https://github.com/Treyson-Grange/JukeTUI), but can be useful elsewhere.

## Library Features

-   Emoji Detection: Identify the presence of emojis in a string via regex
-   Emoji Removal: Strip emojis from a given string via regex
-   Emoji Rune Count: Issues arise when emojis are 2+ runes. If we can account for the runes in each emoji, we can account for these spacing issues.

## Installation

Install the library on your terminal:

```bash
go get github.com/Treyson-Grange/go-moji-ui
```

## Usage

```go
import (
    moji "github.com/Treyson-Grange/go-moji-ui"
)

func main() {
    text := "Hello, World! 😊"
    fmt.Println(moji.ContainsEmoji(text)) // Prints: true

    text = "No Mojis! 🤓🤓🤓🤓🤓🤓🤓🤓🤓"
    fmt.Println(moji.RemoveEmoji(text)) // Prints: No Mojis!

    text = "grapheme cluster moji! 👨‍👩‍👧‍👦"
    fmt.Println(moji.CountEmojiRunes(text)) // Prints: map[👨‍👩‍👧‍👦:7]
}
```
