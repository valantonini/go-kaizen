package runes

import (
	"github.com/matryer/is"
	"testing"
)

// https://golangdocs.com/rune-in-golang Stings are always defined using characters or bytes. In GoLang, strings are
// always made of bytes. Go uses UTF-8 encoding, so any valid character can be represented in Unicode code points. A
// character is defined using “code points” in Unicode. Go language introduced a new term for this code point called
// rune.
//
// A character is defined using “code points” in Unicode. Go language introduced a new term for this code point called rune.
//
// Go rune is also an alias of type int32 because Go uses UTF-8 encoding. Some interesting points about rune and strings.
//
// Strings are made of bytes and they can contain valid characters that can be represented using runes.
// We can use the rune() function to convert string to an array of runes.
// For ASCII characters, the rune value will be the same as the byte value.

func Test_Runes(t *testing.T) {
	Is := is.New(t)

	t.Run("len of ascii will return character count", func(t *testing.T) {
		Is.Equal(len("foo"), 3)
	})

	t.Run("len of ascii []rune will match len of ascii", func(t *testing.T) {
		asciiRunes := []rune("foo")
		Is.Equal(len(asciiRunes), 3)
	})

	t.Run("len of utf8 will not match character count", func(t *testing.T) {
		Is.Equal(len("boo👻"), 7)
	})

	t.Run("len of utf8 will not match character count", func(t *testing.T) {
		utf8Runes := []rune("boo👻")
		Is.Equal(len(utf8Runes), 4)
	})
}
