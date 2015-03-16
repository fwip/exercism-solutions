package house

import (
	"fmt"
	"strings"
)

type Thingy struct {
	noun, verb string
}

var things = [...]Thingy{
	{"house that Jack built.", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

func Embed(n, m string) string {
	return fmt.Sprintf("%s %s", n, m)
}

func Verse(start string, items []string, end string) string {
	verse := start
	for _, embed := range items {
		verse += Embed("", embed)
	}
	verse += " " + end
	return verse
}

func Song() string {

	var verses []string
	for i := range things {
		verse := "This is"
		for j := i; j >= 0; j-- {
			thing := things[j]
			verse += " the " + thing.noun
			if thing.verb != "" {
				verse += "\nthat " + thing.verb
			}
		}
		verses = append(verses, verse)
	}

	return strings.Join(verses, "\n\n")
}
