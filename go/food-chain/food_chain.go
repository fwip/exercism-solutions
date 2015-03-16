package foodchain

import (
	"fmt"
	"strings"
)

const (
	TestVersion = 1
)

type Animal struct {
	name      string // The name of the animal
	rhyme     string // The "funny rhyme" we say when we first introduce it
	catchDesc string // If specified, a different catch name to use when catching
}

var animals = [...]Animal{
	{"fly", "", ""},
	{"spider", "It wriggled and jiggled and tickled inside her.", " that wriggled and jiggled and tickled inside her"},
	{"bird", "How absurd to swallow a bird!", ""},
	{"cat", "Imagine that, to swallow a cat!", ""},
	{"dog", "What a hog, to swallow a dog!", ""},
	{"goat", "Just opened her throat and swallowed a goat!", ""},
	{"cow", "I don't know how she swallowed a cow!", ""},
	{"horse", "She's dead, of course!", ""},
}
var intro_format = "I know an old lady who swallowed a %s."
var reason_format = "She swallowed the %s to catch the %s%s."
var last_line = "I don't know why she swallowed the fly. Perhaps she'll die."

func Verse(n int) string {
	if n < 1 || n > len(animals) {
		return ""
	}

	var lines []string

	index := n - 1
	lines = append(lines, fmt.Sprintf(intro_format, animals[index].name))
	if animals[index].rhyme != "" {
		lines = append(lines, animals[index].rhyme)
	}

	if n < len(animals) {
		for index = n - 1; index > 0; index-- {
			lines = append(lines, fmt.Sprintf(reason_format, animals[index].name, animals[index-1].name, animals[index-1].catchDesc))
		}
		lines = append(lines, last_line)
	}

	verse := strings.Join(lines, "\n")
	return verse
}

func Verses(start, end int) string {
	var verses []string
	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}

func Song() (song string) {
	song = Verses(1, len(animals))
	return song
}
